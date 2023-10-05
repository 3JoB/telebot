package middleware

import (
	"errors"

	tele "github.com/3JoB/telebot/v2"
)

// AutoRespond returns a middleware that automatically responds
// to every callback.
func AutoRespond(c *tele.Context) error {
	if c.Callback() != nil {
		defer c.Respond()
	}
	return c.Next()
}

// IgnoreVia returns a middleware that ignores all the
// "sent via" messages.
func IgnoreVia(c *tele.Context) error {
	if msg := c.Message(); msg != nil && msg.Via != nil {
		return nil
	}
	return c.Next()
}

// Recover returns a middleware that recovers a panic happened in
// the handler.
func Recover(onError ...func(error)) tele.HandlerFunc {
	return func(c *tele.Context) error {
		var f func(error)
		if len(onError) > 0 {
			f = onError[0]
		} else {
			f = func(err error) {
				c.Bot().OnError(err, nil)
			}
		}

		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					f(err)
				} else if s, ok := r.(string); ok {
					f(errors.New(s))
				}
			}
		}()

		return c.Next()
	}
}
