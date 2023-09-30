package middleware

import (
	"io"
	"os"

	"github.com/3JoB/unsafeConvert"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"

	tele "github.com/3JoB/telebot"
)

// Logger returns a middleware that logs incoming updates.
// If no custom logger provided, log.Default() will be used.
func Logger(writers ...io.Writer) tele.MiddlewareFunc {
	var w io.Writer
	if len(writers) > 0 {
		w = zerolog.MultiLevelWriter(writers...)
	} else {
		w = zerolog.ConsoleWriter{Out: os.Stdout}
	}
	l := zerolog.New(w).With().Timestamp().Logger()

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c *tele.Context) error {
			data, _ := json.MarshalIndent(c.Update(), "", "  ")
			l.Info().Msg(unsafeConvert.StringSlice(data))
			return next(c)
		}
	}
}
