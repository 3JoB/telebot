package crare

import (
	"io"

	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
	"github.com/rs/zerolog"
)

// The idea of Logger comes from https://github.com/tucnak/telebot/issues/619.
//
// The Logger interface allows you to customize log wrappers for TEP,
// which uses Zerolog-based wrappers by default.
type Logger interface {
	Println(a ...any)
	Panicf(format string, a ...any)
	Printf(format string, a ...any)
	OnError(error, *Context)
}

type LoggerZerolog struct {
	l zerolog.Logger
}

func NewZeroLogger(writers ...io.Writer) Logger {
	var w io.Writer
	if len(writers) > 0 {
		w = zerolog.MultiLevelWriter(writers...)
	} else {
		w = zerolog.NewConsoleWriter()
	}
	return &LoggerZerolog{
		l: zerolog.New(w).With().Timestamp().Logger(),
	}
}

func (z *LoggerZerolog) Println(v ...any) {
	if len(v) > 0 {
		v = append(v, "\n")
	}
	z.l.Print(v)
}

func (z *LoggerZerolog) Printf(format string, a ...any) {
	z.l.Printf(format, a...)
}

func (z *LoggerZerolog) Panicf(format string, a ...any) {
	if len(a) > 0 {
		z.l.Debug().Msgf(format, a...)
		return
	}
	z.l.Debug().Msg(format)
}

func (z *LoggerZerolog) OnError(err error, c *Context) {
	var message string
	if c != nil {
		message = litefmt.PSprint(unsafeConvert.Itoa(c.Update().ID), " ", err.Error())
	} else {
		message = err.Error()
	}
	z.doPrint(message)
}

func (z *LoggerZerolog) doPrint(v string) {
	if e := z.l.Debug(); e.Enabled() {
		e.CallerSkipFrame(1).Msg(v)
	}
}
