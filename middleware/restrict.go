package middleware

import tele "github.com/3JoB/telebot"

// RestrictConfig defines config for Restrict middleware.
type RestrictConfig struct {
	// Chats is a list of chats that are going to be affected
	// by either In or Out function.
	Chats []int64

	// In defines a function that will be called if the chat
	// of an update will be found in the Chats list.
	In tele.HandlerFunc

	// Out defines a function that will be called if the chat
	// of an update will NOT be found in the Chats list.
	Out tele.HandlerFunc
}

// Restrict returns a middleware that handles a list of provided
// chats with the logic defined by In and Out functions.
// If the chat is found in the Chats field, In function will be called,
// otherwise Out function will be called.
func Restrict(v RestrictConfig) tele.HandlerFunc {
	return func(c *tele.Context) error {
		for _, chat := range v.Chats {
			if chat == c.Sender().ID {
				return c.Next()
			}
		}
		return nil
	}
}

// Blacklist returns a middleware that skips the update for users
// specified in the chats field.
func Blacklist(chats ...int64) tele.HandlerFunc {
	return func(ctx *tele.Context) error {
		return Restrict(RestrictConfig{
			Chats: chats,
		})(ctx)
	}
}

// Whitelist returns a middleware that skips the update for users
// NOT specified in the chats field.
func Whitelist(chats ...int64) tele.HandlerFunc {
	return func(ctx *tele.Context) error {
		return Restrict(RestrictConfig{
			Chats: chats,
		})(ctx)
	}
}
