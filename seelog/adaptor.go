package seelog

import (
	"fmt"
	"strings"

	"github.com/cihub/seelog"
)

type seeLogger struct {
	slog seelog.LoggerInterface
}

func LoggerFromConfigAsFile(file string) (*seeLogger, error) {
	logger, err := seelog.LoggerFromConfigAsFile(file)
	if err != nil {
		return nil, err
	}
	return Logger(logger), nil
}

func LoggerFromConfigAsString(data string) (*seeLogger, error) {
	logger, err := seelog.LoggerFromConfigAsString(data)
	if err != nil {
		return nil, err
	}
	return Logger(logger), nil
}

func LoggerFromConfigAsBytes(data []byte) (*seeLogger, error) {
	logger, err := seelog.LoggerFromConfigAsBytes(data)
	if err != nil {
		return nil, err
	}
	return Logger(logger), nil
}

func Logger(slog seelog.LoggerInterface) *seeLogger {
	l := &seeLogger{slog}
	l.slog.SetAdditionalStackDepth(1)
	return l
}

func (l *seeLogger) Debug(format interface{}, args ...interface{}) {
	switch f := format.(type) {
	case string:
		l.slog.Debugf(f, args...)
	default:
		l.slog.Debugf(fmt.Sprint(format)+strings.Repeat(" %v", len(args)), args...)
	}

}

func (l *seeLogger) Info(format interface{}, args ...interface{}) {
	switch f := format.(type) {
	case string:
		l.slog.Infof(f, args...)
	default:
		l.slog.Infof(fmt.Sprint(format)+strings.Repeat(" %v", len(args)), args...)
	}
}

func (l *seeLogger) Warn(format interface{}, args ...interface{}) error {
	switch f := format.(type) {
	case string:
		return l.slog.Warnf(f, args...)
	default:
		return l.slog.Warnf(fmt.Sprint(format)+strings.Repeat(" %v", len(args)), args...)
	}
}

func (l *seeLogger) Error(format interface{}, args ...interface{}) error {
	switch f := format.(type) {
	case string:
		return l.slog.Errorf(f, args...)
	default:
		return l.slog.Errorf(fmt.Sprint(format)+strings.Repeat(" %v", len(args)), args...)
	}
}

func (l *seeLogger) Critical(format interface{}, args ...interface{}) error {
	switch f := format.(type) {
	case string:
		return l.slog.Criticalf(f, args...)
	default:
		return l.slog.Criticalf(fmt.Sprint(format)+strings.Repeat(" %v", len(args)), args...)
	}
}
