package telebot

import (
	"testing"

	"github.com/3JoB/unsafeConvert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
