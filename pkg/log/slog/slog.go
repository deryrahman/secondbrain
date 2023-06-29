package slog

import (
	"os"

	"github.com/deryrahman/secondbrain/pkg/errors"
	"github.com/deryrahman/secondbrain/pkg/log"
	"github.com/gookit/slog"
)

var (
	_ log.Logger                = (*logger)(nil)
	_ log.LoggerWithErrReceiver = (*logger)(nil)
)

type logger struct {
	*slog.SugaredLogger
}

func NewSLog(level slog.Level, setFns ...slog.SugaredLoggerFn) *logger {
	fns := []slog.SugaredLoggerFn{
		func(sl *slog.SugaredLogger) {
			sl.ReportCaller = false
			sl.CallerFlag = slog.CallerFlagFnlFcn
			sl.Formatter.(*slog.TextFormatter).EnableColor = true
		},
	}
	if len(setFns) > 0 {
		fns = append(fns, setFns...)
	}

	l := slog.NewSugaredLogger(os.Stdout, level, fns...)
	return &logger{l}
}

func (l *logger) Error(err error) {
	r := l.Record()
	if e, ok := err.(errors.ErrorCause); ok {
		r.Caller = e.At()
	}
	r.Logf(slog.ErrorLevel, err.Error())
}
