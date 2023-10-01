package telebot

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/ulib/pool"
	"github.com/3JoB/unsafeConvert"
	"github.com/goccy/go-json"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Raw lets you call any method of Bot API manually.
// It also handles API errors, so you only need to unwrap
// result field from json data.
//
// If you do not pass in the payload value, you need to define the type of T.
//
//		Example:
//	    	Raw[int](b, "getMe")
//
// It now returns a *bytes.Buffer, which will be automatically returned to the
// pool most of the time, but if you call a method such as Raw alone that returns
// a Buffer pointer, please use the ReleaseBuffer() method to save it back to the pool.
func Raw[T any](b *Bot, method string, payload ...T) (*bytes.Buffer, error) {
	url := b.buildUrl(method)
	req := b.client.AcquireRequest()
	buf := pool.NewBuffer()
	req.SetRequestURI(url)
	req.SetWriter(buf)
	req.MethodPOST()
	if len(payload) > 0 {
		if err := req.WriteJson(payload[0]); err != nil {
			ReleaseBuffer(buf)
			return nil, err
		}
	}

	resp, err := req.Do()
	if err != nil {
		ReleaseBuffer(buf)
		return nil, wrapError(err)
	}
	//buf.Write(resp.Bytes())
	defer resp.Release()

	if b.verbose {
		verbose(method, payload, buf)
	}

	// returning data as well
	return buf, extractOk(buf)
}

// Raw lets you call any method of Bot API manually.
// It also handles API errors, so you only need to unwrap
// result field from json data.
//
// It now returns a *bytes.Buffer, which will be automatically returned to the
// pool most of the time, but if you call a method such as Raw alone that returns
// a Buffer pointer, please use the ReleaseBuffer() method to save it back to the pool.
func (b *Bot) Raw(method string, payload ...any) (*bytes.Buffer, error) {
	url := b.buildUrl(method)

	// Cancel the request immediately without waiting for the timeout  when bot is about to stop.
	// This may become important if doing long polling with long timeout.
	/*exit := make(chan struct{})
	defer close(exit)
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		select {
		case <-b.stopClient:
			cancel()
		case <-exit:
		}
	}()*/
	buf := pool.NewBuffer()
	req := b.client.AcquireRequest()

	req.SetRequestURI(url)
	req.MethodPOST()

	if len(payload) > 0 {
		if payload[0] != nil {
			if err := req.WriteJson(payload[0]); err != nil {
				return nil, err
			}
		}
	}

	resp, err := req.Do()
	if err != nil {
		ReleaseBuffer(buf)
		return nil, wrapError(err)
	}
	buf.Write(resp.Bytes())
	defer resp.Release()
	if b.verbose {
		verbose(method, payload, buf)
	}

	// returning data as well
	return buf, extractOk(buf)
}

func (b *Bot) buildUrl(method string) string {
	return litefmt.PSprint(b.URL, "/bot", b.Token, "/", method)
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
		return Raw(b, method, params)
	}

	pipeReader, pipeWriter := io.Pipe()
	writer := multipart.NewWriter(pipeWriter)

	go func() {
		defer pipeWriter.Close()

		for field, file := range rawFiles {
			if err := addFileToWriter(writer, files[field].fileName, field, file); err != nil {
				_ = pipeWriter.CloseWithError(err)
				return
			}
		}
		for field, value := range params {
			if err := writer.WriteField(field, value.(string)); err != nil {
				_ = pipeWriter.CloseWithError(err)
				return
			}
		}
		if err := writer.Close(); err != nil {
			_ = pipeWriter.CloseWithError(err)
			return
		}
	}()

	url := b.buildUrl(method)
	req := b.client.AcquireRequest()
	req.SetRequestURI(url)
	buf := pool.NewBuffer()

	if err := req.WriteFile(writer.FormDataContentType(), pipeReader); err != nil {
		err = wrapError(err)
		_ = pipeReader.CloseWithError(err)
		req.Release()
		ReleaseBuffer(buf)
		return nil, err
	}
	
	resp, err := req.Do()
	if err != nil {
		err = wrapError(err)
		_ = pipeReader.CloseWithError(err)
		ReleaseBuffer(buf)
		return nil, err
	}
	defer resp.Release()

	if resp.IsStatusCode(500) {
		return nil, ErrInternal
	}

	return buf, extractOk(buf)
}

func addFileToWriter(writer *multipart.Writer, filename, field string, file any) error {
	var reader io.Reader
	if r, ok := file.(io.Reader); ok {
		reader = r
	} else if path, ok := file.(string); ok {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		reader = f
	} else {
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

	data, err := Raw(b, "sendMessage", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

func (b *Bot) sendMedia(media Media, params map[string]any, files map[string]File) (*Message, error) {
	kind := media.MediaType()
	what := "send" + cases.Title(language.English).String(kind)

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
	data, err := Raw[bool](b, "getMe")
	defer ReleaseBuffer(data)
	if err != nil {
		return nil, err
	}

	var resp Response[*User]
	if err := b.json.Unmarshal(data.Bytes(), &resp); err != nil {
		return nil, wrapError(err)
	}

	return resp.Result, nil
}

func (b *Bot) getUpdates(offset, limit int, timeout time.Duration, allowed []string) ([]*Update, error) {
	params := map[string]any{
		"offset":  offset,
		"timeout": int(timeout / time.Second),
	}
	params["allowed_updates"] = allowed

	if limit != 0 {
		params["limit"] = limit
	}

	data, err := Raw(b, "getUpdates", params)
	if err != nil {
		return nil, err
	}
	defer ReleaseBuffer(data)

	var resp Response[[]*Update]
	if err := b.json.NewDecoder(data).Decode(&resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

type extracts struct {
	Ok          bool           `json:"ok"`
	Code        int            `json:"error_code"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
}

// extractOk checks given result for error. If result is ok returns nil.
// In other cases it extracts API error. If error is not presented
// in errors.go, it will be prefixed with `unknown` keyword.
func extractOk(data *bytes.Buffer) error {
	var e extracts
	if err := json.Unmarshal(data.Bytes(), &e); err != nil {
		return err
	}
	if e.Ok {
		return nil
	}

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
			MigratedTo: int64(migratedTo.(float64)),
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
			RetryAfter: int(retryAfter.(float64)),
		}
	} else {
		err = fmt.Errorf("telegram: %s (%d)", e.Description, e.Code)
	}

	return err
}

// extractMessage extracts common Message result from given data.
// Should be called after extractOk or b.Raw() to handle possible errors.
func extractMessage(data *bytes.Buffer) (*Message, error) {
	defer pool.ReleaseBuffer(data)
	var resp Response[*Message]
	if err := json.Unmarshal(data.Bytes(), &resp); err != nil {
		var resp Response[bool]
		if err := json.NewDecoder(data).Decode(&resp); err != nil {
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
	_ = json.Indent(buf, b, "", "  ")
	return buf.String()
}

func verbose(method string, payload any, data *bytes.Buffer) {
	body, _ := json.Marshal(payload)
	body = bytes.ReplaceAll(body, unsafeConvert.ByteSlice(`\"`), unsafeConvert.ByteSlice(`"`))
	body = bytes.ReplaceAll(body, unsafeConvert.ByteSlice(`"{`), unsafeConvert.ByteSlice(`{`))
	body = bytes.ReplaceAll(body, unsafeConvert.ByteSlice(`}"`), unsafeConvert.ByteSlice(`}`))

	log.Printf(
		"[verbose] telebot: sent request\nMethod: %v\nParams: %v\nResponse: %v",
		method, indent(body), indent(data.Bytes()),
	)
}
