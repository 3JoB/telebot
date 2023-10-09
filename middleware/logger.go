package middleware

import (
	"github.com/3JoB/unsafeConvert"

	tele "github.com/3JoB/telebot/v2"
)

// Logger returns a middleware that logs incoming updates.
// If no custom logger provided, log.Default() will be used.
func Logger(logger tele.Logger) tele.HandlerFunc {
	return func(c *tele.Context) error {
		data, _ := c.Bot().Json().MarshalIndent(c.Update(), "", "  ")
		logger.Println(unsafeConvert.StringSlice(data))
		return c.Next()
	}
}
