package crare

import (
	"testing"

	"github.com/3JoB/ulib/pool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractOk(t *testing.T) {
	buf := pool.NewBuffer()
	data := `{"ok": true, "result": {}}`
	buf.WriteString(data)
	require.NoError(t, extractOk(buf))
	buf.Reset()

	data = `{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: reply message not found"
	}`
	buf.WriteString(data)

	assert.EqualError(t, extractOk(buf), ErrNotFoundToReply.Error())
	buf.Reset()

	data = `{
		"ok": false,
		"error_code": 429,
		"description": "Too Many Requests: retry after 8",
		"parameters": {"retry_after": 8}
	}`
	buf.WriteString(data)
	assert.Equal(t, FloodError{
		err:        NewError(429, "Too Many Requests: retry after 8"),
		RetryAfter: 8,
	}, extractOk(buf))
	buf.Reset()

	data = `{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: group chat was upgraded to a supergroup chat",
		"parameters": {"migrate_to_chat_id": -100123456789}
	}`
	buf.WriteString(data)
	assert.Equal(t, GroupError{
		err:        ErrGroupMigrated,
		MigratedTo: -100123456789,
	}, extractOk(buf))
}

func TestExtractMessage(t *testing.T) {
	buf := pool.NewBuffer()
	data := `{"ok":true,"result":true}`
	buf.WriteString(data)
	_, err := extractMessage(buf)
	assert.Equal(t, ErrTrueResult, err)
	buf.Reset()

	data = `{"ok":true,"result":{"foo":"bar"}}`
	buf.WriteString(data)
	_, err = extractMessage(buf)
	require.NoError(t, err)
}
