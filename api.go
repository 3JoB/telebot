package crare

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/ulib/pool"
	"github.com/3JoB/unsafeConvert"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/3JoB/telebot/v2/pkg/updates"
)

// Raw lets you call any method of Bot API manually.
// It also handles API errors, so you only need to unwrap
// result field from json data.
//
// It now returns a *bytes.Buffer, which will be automatically returned to the
// pool most of the time, but if you call a method such as Raw alone that returns
// a Buffer pointer, please use the ReleaseBuffer() method to save it back to the pool.
//
// It will be automatically released when err != nil, so there is no need to release it
// again. At the same time, additional judgment is made on the null pointer in the pool.
func (b *Bot) Raw(method string, payload ...any) (*bytes.Buffer, error) {
	url := b.buildUrl(method)

	req, resp := b.client.Acquire()
	defer b.client.Release(req, resp)
	buf := pool.NewBuffer()
	req.SetRequestURI(url)
	req.MethodPOST()
	req.SetWriter(buf)

	if len(payload) > 0 && payload[0] != nil {
		if err := req.WriteJson(payload[0]); err != nil {
			ReleaseBuffer(buf)
			return nil, wrapError(err)
		}
	}

	if err := req.Do(); err != nil {
		ReleaseBuffer(buf)
		return nil, wrapError(err)
	}

	if b.verbose {
		b.verboses(method, payload, buf)
	}

	// returning data as well
	return buf, extractOk(buf)
}

func (b *Bot) buildUrl(method string) string {
	if b.URLCache != "" {
		return litefmt.PSprint(b.URLCache, method)
	}
	b.URLCache = litefmt.PSprint(b.URL, "/bot", b.Token, "/")
	return litefmt.PSprint(b.URLCache, method)
}

func (b *Bot) sendFiles(method string, files map[string]File, params map[string]any) (*bytes.Buffer, error) {
	rawFiles := make(map[string]any)
	for name, f := range files {
		switch {
		case f.InCloud():
			params[name] = f.FileID
		case f.FileURL != "":
			params[name] = f.FileURL
		case f.OnDisk():
			rawFiles[name] = f.FileLocal
		case f.FileReader != nil:
			rawFiles[name] = f.FileReader
		default:
			return nil, fmt.Errorf("telebot: file for field %s doesn't exist", name)
		}
	}

	if len(rawFiles) == 0 {
		return b.Raw(method, params)
	}

	pipeReader, pipeWriter := io.Pipe()
	writer := multipart.NewWriter(pipeWriter)

	go func() {
		defer pipeWriter.Close()

		for field, file := range rawFiles {
			if err := addFileToWriter(writer, files[field].fileName, field, file); err != nil {
				pipeWriter.CloseWithError(err) //nolint:errcheck
				return
			}
		}
		for field, value := range params {
			if err := writer.WriteField(field, value.(string)); err != nil {
				pipeWriter.CloseWithError(err) //nolint:errcheck
				return
			}
		}
		if err := writer.Close(); err != nil {
			pipeWriter.CloseWithError(err) //nolint:errcheck
			return
		}
	}()

	url := b.buildUrl(method)
	req, resp := b.client.Acquire()
	defer b.client.Release(req, resp)
	req.SetRequestURI(url)
	buf := pool.NewBuffer()

	if err := req.WriteFile(writer.FormDataContentType(), pipeReader); err != nil {
		err = wrapError(err)
		pipeReader.CloseWithError(err) //nolint:errcheck
		ReleaseBuffer(buf)
		return nil, err
	}

	if err := req.Do(); err != nil {
		err = wrapError(err)
		pipeReader.CloseWithError(err) //nolint:errcheck
		ReleaseBuffer(buf)
		return nil, err
	}
	defer b.client.ReleaseResponse(resp)

	if resp.IsStatusCode(500) {
		return nil, ErrInternal
	}

	return buf, extractOk(buf)
}

func addFileToWriter(writer *multipart.Writer, filename, field string, file any) error {
	var reader io.Reader
	switch r := file.(type) {
	case io.Reader:
		reader = r
	case string:
		f, err := os.Open(r)
		if err != nil {
			return err
		}
		defer f.Close()
		reader = f
	default:
		return fmt.Errorf("telebot: file for field %v should be io.ReadCloser or string", field)
	}

	part, err := writer.CreateFormFile(field, filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, reader)
	return err
}

func (b *Bot) sendText(to Recipient, text string, opt *SendOptions) (*Message, error) {
	params := map[string]any{
		"chat_id": to.Recipient(),
		"text":    text,
	}
	b.embedSendOptions(params, opt)

	data, err := b.Raw("sendMessage", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

func (b *Bot) sendMedia(media Media, params map[string]any, files map[string]File) (*Message, error) {
	kind := media.MediaType()
	what := litefmt.PSprint("send", cases.Title(language.English).String(kind))

	if kind == "videoNote" {
		kind = "video_note"
	}

	sendFiles := map[string]File{kind: *media.MediaFile()}
	for k, v := range files {
		sendFiles[k] = v
	}

	data, err := b.sendFiles(what, sendFiles, params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

func (b *Bot) getMe() (*User, error) {
	data, err := b.Raw("getMe")
	if err != nil {
		return nil, err
	}
	defer ReleaseBuffer(data)

	var resp Response[*User]
	if err := b.json.NewEncoder(data).Encode(&resp); err != nil {
		return nil, wrapError(err)
	}

	return resp.Result, nil
}

func (b *Bot) getUpdates(offset, limit int, timeout time.Duration, allowed []string) ([]Update, error) {
	params := updates.AcquireParams()
	params.Offset = offset
	params.Timeout = int(timeout / time.Second)
	params.AllowedUpdates = allowed
	if limit != 0 {
		params.Limit = limit
	}
	data, err := b.Raw("getUpdates", params)
	if err != nil {
		updates.ReleaseParams(params)
		return nil, err
	}
	updates.ReleaseParams(params)
	defer ReleaseBuffer(data)

	var resp Response[[]Update]
	if err := b.json.NewDecoder(data).Decode(&resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

type extracts struct {
	Ok          bool               `json:"ok"`
	Code        int                `json:"error_code"`
	Description string             `json:"description"`
	Parameters  map[string]float64 `json:"parameters"`
}

// extractOk checks given result for error. If result is ok returns nil.
// In other cases it extracts API error. If error is not presented
// in errors.go, it will be prefixed with `unknown` keyword.
func extractOk(data *bytes.Buffer) error {
	var e extracts
	if err := defaultJson.Unmarshal(data.Bytes(), &e); err != nil {
		ReleaseBuffer(data)
		return err
	}
	if e.Ok {
		return nil
	}
	ReleaseBuffer(data)

	err := Err(e.Description)
	switch err {
	case nil:
	case ErrGroupMigrated:
		migratedTo, ok := e.Parameters["migrate_to_chat_id"]
		if !ok {
			return NewError(e.Code, e.Description)
		}

		return GroupError{
			err:        err.(*Error),
			MigratedTo: int64(migratedTo),
		}
	default:
		return err
	}

	if e.Code == 429 {
		retryAfter, ok := e.Parameters["retry_after"]
		if !ok {
			return NewError(e.Code, e.Description)
		}

		err = FloodError{
			err:        NewError(e.Code, e.Description),
			RetryAfter: int(retryAfter),
		}
	} else {
		err = fmt.Errorf("telegram: %s (%d)", e.Description, e.Code)
	}

	return err
}

// extractMessage extracts common Message result from given data.
// Should be called after extractOk or b.Raw() to handle possible errors.
//
// This method will automatically release the incoming Buffer
func extractMessage(data *bytes.Buffer) (*Message, error) {
	defer ReleaseBuffer(data)
	var resp Response[*Message]
	if err := defaultJson.Unmarshal(data.Bytes(), &resp); err != nil {
		var resp Response[bool]
		if err := defaultJson.NewDecoder(data).Decode(&resp); err != nil {
			return nil, wrapError(err)
		}
		if resp.Result {
			return nil, ErrTrueResult
		}
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

func indent(b []byte) string {
	buf := pool.NewBuffer()
	defer ReleaseBuffer(buf)
	_ = defaultJson.Indent(buf, b, "", "  ")
	return buf.String()
}

func (b *Bot) verboses(method string, payload any, data *bytes.Buffer) {
	body, _ := defaultJson.Marshal(payload)
	body = bytes.ReplaceAll(body, unsafeConvert.BytePointer(`\"`), unsafeConvert.BytePointer(`"`))
	body = bytes.ReplaceAll(body, unsafeConvert.BytePointer(`"{`), unsafeConvert.BytePointer(`{`))
	body = bytes.ReplaceAll(body, unsafeConvert.BytePointer(`}"`), unsafeConvert.BytePointer(`}`))

	b.logger.Printf(
		"[verbose] telebot: sent request\nMethod: %v\nParams: %v\nResponse: %v",
		method, indent(body), indent(data.Bytes()),
	)
}
