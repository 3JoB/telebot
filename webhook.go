package telebot

import (
	"context"
	"fmt"
	"net"

	realip "github.com/3JoB/atreugo-realip"
	mmdb "github.com/3JoB/maxminddb-golang"
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"
)

type MMDB_ASN struct {
	AutonomousSystemNumber       uint   `maxminddb:"autonomous_system_number"`
	AutonomousSystemOrganization string `maxminddb:"autonomous_system_organization"`
}

// A WebhookTLS specifies the path to a key and a cert so the poller can open
// a TLS listener.
type WebhookTLS struct {
	NoLocal bool   `json:"-"`
	Key     string `json:"key"`
	Cert    string `json:"cert"`
}

// A WebhookEndpoint describes the endpoint to which telegram will send its requests.
// This must be a public URL and can be a loadbalancer or something similar. If the
// endpoint uses TLS and the certificate is self-signed you have to add the certificate
// path of this certificate so telegram will trust it. This field can be ignored if you
// have a trusted certificate (letsencrypt, ...).
type WebhookEndpoint struct {
	PublicURL string `json:"public_url"`
	Cert      string `json:"cert"`
}

// A Webhook configures the poller for webhooks. It opens a port on the given
// listen address. If TLS is filled, the listener will use the key and cert to open
// a secure port. Otherwise it will use plain HTTP.
//
// If you have a loadbalancer ore other infrastructure in front of your service, you
// must fill the Endpoint structure so this poller will send this data to telegram. If
// you leave these values empty, your local address will be sent to telegram which is mostly
// not what you want (at least while developing). If you have a single instance of your
// bot you should consider to use the LongPoller instead of a WebHook.
//
// You can also leave the Listen field empty. In this case it is up to the caller to
// add the Webhook to a http-mux.
type Webhook struct {
	Host           string   `json:"url"`
	Listen         string   `json:"-"`
	SecretToken    string   `json:"secret_token"`
	IP             string   `json:"ip_address"`
	MaxConnections int      `json:"max_connections"`
	AllowedUpdates []string `json:"allowed_updates"`
	DropUpdates    bool     `json:"drop_pending_updates"`

	// (WebhookInfo)
	HasCustomCert bool `json:"has_custom_certificate"`

	PendingUpdates    int    `json:"pending_update_count"`
	ErrorMessage      string `json:"last_error_message"`
	SyncErrorUnixtime int64  `json:"last_synchronization_error_date"`
	ErrorUnixtime     int64  `json:"last_error_date"`

	Verify *WebhookVerify `json:"-"`

	TLS      *WebhookTLS
	Endpoint *WebhookEndpoint

	dest chan<- Update
	bot  *Bot
}

type WebhookVerify struct {
	DB     string // maxmind mmdb path
	reader *mmdb.Reader
}

func (v *WebhookVerify) Verify(ip string) (is bool) {
	var asn MMDB_ASN
	nip := net.ParseIP(ip)
	if v.reader.Lookup(nip, &asn) != nil {
		return
	}
	if asn.AutonomousSystemNumber == 62041 {
		is = true
	}
	return
}

func (h *Webhook) getFiles() map[string]File {
	m := make(map[string]File)

	if h.TLS != nil {
		if !h.TLS.NoLocal {
			m["certificate"] = FromDisk(h.TLS.Cert)
		}
	}
	// check if it is overwritten by an endpoint
	if h.Endpoint != nil {
		if h.Endpoint.Cert == "" {
			// this can be the case if there is a loadbalancer or reverseproxy in
			// front with a public cert. in this case we do not need to upload it
			// to telegram. we delete the certificate from the map, because someone
			// can have an internal TLS listener with a private cert
			delete(m, "certificate")
		} else {
			// someone configured a certificate
			m["certificate"] = FromDisk(h.Endpoint.Cert)
		}
	}
	return m
}

func (h *Webhook) getParams() map[string]any {
	params := make(map[string]any)

	if h.MaxConnections != 0 {
		params["max_connections"] = h.MaxConnections
	}
	if len(h.AllowedUpdates) > 0 {
		data, _ := h.bot.json.Marshal(h.AllowedUpdates)
		params["allowed_updates"] = unsafeConvert.StringPointer(data)
	}
	if h.IP != "" {
		params["ip_address"] = h.IP
	}
	if h.DropUpdates {
		params["drop_pending_updates"] = h.DropUpdates
	}
	if h.SecretToken != "" {
		params["secret_token"] = h.SecretToken
	}

	if h.TLS != nil {
		params["url"] = "https://" + h.Host
	} else {
		// this will not work with telegram, they want TLS
		// but i allow this because telegram will send an error
		// when you register this hook. in their docs they write
		// that port 80/http is allowed ...
		params["url"] = "http://" + h.Host
	}
	if h.Endpoint != nil {
		params["url"] = h.Endpoint.PublicURL
	}
	return params
}

func (h *Webhook) Poll(b *Bot, dest chan Update, stop chan struct{}) {
	if err := b.SetWebhook(h); err != nil {
		b.OnError(err, nil)
		close(stop)
		return
	}

	// store the variables so the HTTP-handler can use 'em
	h.dest = dest
	h.bot = b

	if h.Listen == "" || h.Host == "" {
		h.waitForStop(stop)
		return
	}

	conf := &atreugo.Config{
		Addr: h.Listen,
		Name: "Crare/2",
	}
	if h.TLS != nil {
		conf.CertFile = h.TLS.Cert
		conf.CertKey = h.TLS.Key
	}
	server := atreugo.New(*conf)

	if h.Verify != nil && h.Verify.DB != "" {
		r, err := mmdb.Open(h.Verify.DB)
		if err != nil {
			b.OnError(err, nil)
			close(stop)
			return
		}
		h.Verify.reader = r
		server.UseBefore(h.IPValidation)
	}

	server.UseBefore(h.TokenValidation)
	server.ANY("/", h.Serve)

	go func(stop chan struct{}) {
		h.waitForStop(stop)
		_ = server.ShutdownWithContext(context.Background())
	}(stop)

	if err := server.ListenAndServe(); err != nil {
		b.OnError(err, nil)
	}
}

func (h *Webhook) waitForStop(stop chan struct{}) {
	<-stop
	close(stop)
}

func (h *Webhook) IPValidation(rc *atreugo.RequestCtx) error {
	if h.Verify.Verify(realip.FromRequest(rc)) {
		return rc.Next()
	}
	_ = rc.Conn().Close()
	return nil
}

func (h *Webhook) TokenValidation(rc *atreugo.RequestCtx) error {
	if h.SecretToken != "" && unsafeConvert.StringPointer(rc.Request.Header.Peek("X-Telegram-Bot-Api-Secret-Token")) != h.SecretToken {
		return rc.TextResponse("invalid secret token in request", 401)
	}
	return rc.Next()
}

// The handler simply reads the update from the body of the requests
// and writes them to the update channel.
func (h *Webhook) Serve(rc *atreugo.RequestCtx) error {
	var update Update
	if err := h.bot.json.Unmarshal(rc.Request.Body(), &update); err != nil {
		err = fmt.Errorf("cannot decode update: %v", err)
		h.bot.debug(err)
		return err
	}
	h.dest <- update
	return nil
}

// Webhook returns the current webhook status.
func (b *Bot) Webhook() (*Webhook, error) {
	data, err := b.Raw("getWebhookInfo")
	if err != nil {
		return nil, err
	}
	defer ReleaseBuffer(data)

	var resp Response[Webhook]
	if err := b.json.NewDecoder(data).Decode(&resp); err != nil {
		return nil, wrapError(err)
	}
	return &resp.Result, nil
}

// SetWebhook configures a bot to receive incoming
// updates via an outgoing webhook.
func (b *Bot) SetWebhook(w *Webhook) error {
	d, err := b.sendFiles("setWebhook", w.getFiles(), w.getParams())
	ReleaseBuffer(d)
	return err
}

// RemoveWebhook removes webhook integration.
func (b *Bot) RemoveWebhook(dropPending ...bool) error {
	drop := false
	if len(dropPending) > 0 {
		drop = dropPending[0]
	}
	d, err := b.Raw("deleteWebhook", map[string]bool{
		"drop_pending_updates": drop,
	})
	ReleaseBuffer(d)
	return err
}
