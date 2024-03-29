package db

import (
	"context"
	"io"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/darabuchi/log"
	"github.com/gookit/color"
	"gorm.io/gorm/logger"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	l := &Logger{
		logger: log.Clone().SetCallerDepth(4),
	}
	l.LogMode(logger.Info)
	return l
}

func NewDiscard() *Logger {
	l := &Logger{
		logger: log.Clone().SetCallerDepth(4).SetOutput(io.Discard),
	}
	l.LogMode(logger.Info)
	return l
}

func (l *Logger) LogMode(logLevel logger.LogLevel) logger.Interface {
	switch logLevel {
	case logger.Silent:
		l.logger.SetLevel(log.TraceLevel)
	case logger.Error:
		l.logger.SetLevel(log.ErrorLevel)
	case logger.Warn:
		l.logger.SetLevel(log.WarnLevel)
	case logger.Info:
		l.logger.SetLevel(log.InfoLevel)
	default:
		l.logger.SetLevel(log.DebugLevel)
	}
	return l
}

func (l *Logger) Info(ctx context.Context, s string, i ...interface{}) {
	l.logger.Infof(s, i...)
}

func (l *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.logger.Warnf(s, i...)
}

func (l *Logger) Error(ctx context.Context, s string, i ...interface{}) {
	l.logger.Errorf(s, i...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	var callerName string
	pc, file, callerLine, ok := runtime.Caller(3)
	if ok {
		callerName = runtime.FuncForPC(pc).Name()
	}

	callerDir, callerFunc := log.SplitPackageName(callerName)

	sql, rowsAffected := fc()
	l.logger.Infof("%s %s %s %s",
		color.Yellow.Sprintf("%s:%d %s", path.Join(callerDir, path.Base(file)), callerLine, callerFunc),
		color.Blue.Sprintf("[%s]", time.Since(begin)),
		strings.ReplaceAll(sql, "\n", " "),
		color.Blue.Sprintf("[%d rows]", rowsAffected),
	)
}
