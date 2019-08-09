package log4go

import (
	"context"

	"github.com/hashicorp/logutils"
)

func DefaultLogger() (logger Logger) {
	return NewLogger("", nil)
}

func Debug(format string, v ...interface{}) {
	DefaultLogger().Debug(context.Background(), format, v...)
}

func Info(format string, v ...interface{}) {
	DefaultLogger().Info(context.Background(), format, v...)
}

func Warn(format string, v ...interface{}) {
	DefaultLogger().Warn(context.Background(), format, v...)
}

func Error(format string, v ...interface{}) {
	DefaultLogger().Error(context.Background(), format, v...)
}

func SetFilter(filter *logutils.LevelFilter) {
	DefaultLogger().SetFilter(filter)
}
