package telebot

import (
	"github.com/3JoB/ulib/litefmt"
	"github.com/rs/zerolog"
)

// The idea of Logger comes from https://github.com/tucnak/telebot/issues/619.
//
// The Logger interface allows you to customize log wrappers for TEP,
// which uses Zerolog-based wrappers by default.
type Logger interface {
	Debugf(format string, a ...any)
	Infof(format string, a ...any)
	Warnf(format string, a ...any)
	Errorf(format string, a ...any)
	Panicf(format string, a ...any)
	Println(v ...string)
}

type LoggerZerolog struct {
	l zerolog.Logger
}

func NewZeroLogger() Logger {
	return &LoggerZerolog{
		l: zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger(),
	}
}

func (z *LoggerZerolog) Debugf(format string, a ...any) {
	if len(a) > 0 {
		z.l.Debug().Msgf(format, a...)
		return
	}
	z.l.Debug().Msg(format)
}

func (z *LoggerZerolog) Infof(format string, a ...any) {
	if len(a) > 0 {
		z.l.Debug().Msgf(format, a...)
		return
	}
	z.l.Debug().Msg(format)
}

func (z *LoggerZerolog) Warnf(format string, a ...any) {
	if len(a) > 0 {
		z.l.Debug().Msgf(format, a...)
		return
	}
	z.l.Debug().Msg(format)
}

func (z *LoggerZerolog) Errorf(format string, a ...any) {
	if len(a) > 0 {
		z.l.Debug().Msgf(format, a...)
		return
	}
	z.l.Debug().Msg(format)
}

func (z *LoggerZerolog) Panicf(format string, a ...any) {
	if len(a) > 0 {
		z.l.Debug().Msgf(format, a...)
		return
	}
	z.l.Debug().Msg(format)
}

func (z *LoggerZerolog) Println(v ...string) {
	v = append(v, "\n")
	if e := z.l.Debug(); e.Enabled() {
		e.CallerSkipFrame(1).Msg(litefmt.PSprint(v...))
	}
}
