package log4go

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/logutils"
)

var defaultIdName = "x-log4go-id"
var loggerMap map[string]Logger

// trace log level
var LogLevelTrace = logutils.LogLevel("TRCE")

// debug log level
var LogLevelDebug = logutils.LogLevel("DBUG")

// info log level
var LogLevelInfo = logutils.LogLevel("INFO")

// warning log level
var LogLevelWarn = logutils.LogLevel("WARN")

// error log level
var LogLevelError = logutils.LogLevel("EROR")

var logLevel = []logutils.LogLevel{
	LogLevelTrace, LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError,
}

func init() {
	loggerMap = map[string]Logger{}
}

// Interface for logger, you can implement your own logger with this.
type Logger interface {
	Trace(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})

	TraceCtx(ctx context.Context, format string, v ...interface{})
	DebugCtx(ctx context.Context, format string, v ...interface{})
	InfoCtx(ctx context.Context, format string, v ...interface{})
	WarnCtx(ctx context.Context, format string, v ...interface{})
	ErrorCtx(ctx context.Context, format string, v ...interface{})

	SetFilter(filter *logutils.LevelFilter)
	GetFilter() (filter *logutils.LevelFilter)

	WithField(key string, val interface{}) Logger
	WithError(err error) Logger
}

// Options for logger.
type LoggerOpts struct {
	WithId bool
	IdName string
}

// Create a named logger with specify options.
func NewLogger(name string, opts *LoggerOpts) (logger Logger) {
	var exists bool
	if logger, exists = loggerMap[name]; !exists {
		filter := &logutils.LevelFilter{
			Levels:   logLevel,
			MinLevel: LogLevelInfo,
			Writer:   os.Stdout,
		}
		log.SetOutput(filter)
		logger = &log4GoLogger{filter: filter, opts: formatOpts(opts)}
		loggerMap[name] = logger
	}

	return loggerMap[name]
}

func formatOpts(opts *LoggerOpts) *LoggerOpts {
	if opts == nil {
		return &LoggerOpts{
			WithId: false,
		}
	}

	if opts.IdName == "" {
		opts.IdName = defaultIdName
	}

	return opts
}
