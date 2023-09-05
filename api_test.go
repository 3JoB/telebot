package telebot

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/3JoB/unsafeConvert"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testPayload implements json.Marshaler
// to test json encoding error behaviour.
type testPayload struct{}

func (testPayload) MarshalJSON() ([]byte, error) {
	return nil, errors.New("test error")
}

func testRawServer(w http.ResponseWriter, r *http.Request) {
	switch {
	// causes EOF error on ioutil.ReadAll
	case strings.HasSuffix(r.URL.Path, "/testReadError"):
		// tells the body is 1 byte length but actually it's 0
		w.Header().Set("Content-Length", "1")

	// returns unknown telegram error
	case strings.HasSuffix(r.URL.Path, "/testUnknownError"):
		data, _ := json.Marshal(struct {
			Ok          bool   `json:"ok"`
			Code        int    `json:"error_code"`
			Description string `json:"description"`
		}{
			Ok:          false,
			Code:        400,
			Description: "unknown error",
		})

		w.WriteHeader(400)
		w.Write(data)
	}
}

func TestRaw(t *testing.T) {}

func TestExtractOk(t *testing.T) {
	data := unsafeConvert.ByteSlice(`{"ok": true, "result": {}}`)
	require.NoError(t, extractOk(data))

	data = unsafeConvert.ByteSlice(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: reply message not found"
	}`)
	assert.EqualError(t, extractOk(data), ErrNotFoundToReply.Error())

	data = unsafeConvert.ByteSlice(`{
		"ok": false,
		"error_code": 429,
		"description": "Too Many Requests: retry after 8",
		"parameters": {"retry_after": 8}
	}`)
	assert.Equal(t, FloodError{
		err:        NewError(429, "Too Many Requests: retry after 8"),
		RetryAfter: 8,
	}, extractOk(data))

	data = unsafeConvert.ByteSlice(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: group chat was upgraded to a supergroup chat",
		"parameters": {"migrate_to_chat_id": -100123456789}
	}`)
	assert.Equal(t, GroupError{
		err:        ErrGroupMigrated,
		MigratedTo: -100123456789,
	}, extractOk(data))
}

func TestExtractMessage(t *testing.T) {
	data := unsafeConvert.ByteSlice(`{"ok":true,"result":true}`)
	_, err := extractMessage(data)
	assert.Equal(t, ErrTrueResult, err)

	data = unsafeConvert.ByteSlice(`{"ok":true,"result":{"foo":"bar"}}`)
	_, err = extractMessage(data)
	require.NoError(t, err)
}
