package log4go

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/hashicorp/logutils"
)

var (
	defaultIdName = "x-log4go-id"
	loggerMu      sync.Mutex
	loggerMap     map[string]Logger
)

var (
	LogLevelTrace = logutils.LogLevel("TRCE") // trace log level
	LogLevelDebug = logutils.LogLevel("DBUG") // debug log level
	LogLevelInfo  = logutils.LogLevel("INFO") // info log level
	LogLevelWarn  = logutils.LogLevel("WARN") // warning log level
	LogLevelError = logutils.LogLevel("EROR") // error log level
	logLevel      = []logutils.LogLevel{
		LogLevelTrace,
		LogLevelDebug,
		LogLevelInfo,
		LogLevelWarn,
		LogLevelError,
	}
)

func init() {
	loggerMap = map[string]Logger{}
}

// Interface for logger, you can implement your own logger with this.
type Logger interface {
	Trace(ctx context.Context, format string, v ...interface{})
	Debug(ctx context.Context, format string, v ...interface{})
	Info(ctx context.Context, format string, v ...interface{})
	Warn(ctx context.Context, format string, v ...interface{})
	Error(ctx context.Context, format string, v ...interface{})
	SetFlags(flags int)
	SetFilter(filter *logutils.LevelFilter)
	GetFilter() (filter *logutils.LevelFilter)
}

// Options for logger.
type LoggerOpts struct {
	WithId bool
	IdName interface{}
}

// Create a named logger with specify options.
func NewLogger(name string, opts *LoggerOpts) (logger Logger) {
	var exists bool

	loggerMu.Lock()
	defer loggerMu.Unlock()

	if logger, exists = loggerMap[name]; !exists {
		filter := &logutils.LevelFilter{
			Levels:   logLevel,
			MinLevel: LogLevelInfo,
			Writer:   os.Stdout,
		}
		logger = &log4GoLogger{
			stdLogger: log.New(filter, name, log.LstdFlags),
			filter:    filter,
			opts:      formatOpts(opts),
		}
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
