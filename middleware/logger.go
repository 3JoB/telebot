package middleware

import (
	"github.com/goccy/go-json"
	"log"

	tele "telebot"
)

// Logger returns a middleware that logs incoming updates.
// If no custom logger provided, log.Default() will be used.
func Logger(logger ...*log.Logger) tele.MiddlewareFunc {
	var l *log.Logger
	if len(logger) > 0 {
		l = logger[0]
	} else {
		l = log.Default()
	}

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			data, _ := json.MarshalIndent(c.Update(), "", "  ")
			l.Println(tele.String(data))
			return next(c)
		}
	}
}
