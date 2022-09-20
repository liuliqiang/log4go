package log4go

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hashicorp/logutils"
)

const (
	builtInKeyError = "__liuliqiang_log4go_key_err"
)

type defaultField struct {
	key string
	val interface{}
}
type log4GoLogger struct {
	mutex  sync.Mutex
	filter *logutils.LevelFilter
	fields []defaultField
	opts   *LoggerOpts
}

func (l *log4GoLogger) WithError(err error) Logger {
	// bug(@liuliqiang): multi logger with the same fields
	l.fields = append(l.fields, defaultField{
		key: builtInKeyError,
		val: err,
	})

	return l
}

func (l *log4GoLogger) WithField(key string, val interface{}) Logger {
	// bug(@liuliqiang): multi logger with the same fields
	l.fields = append(l.fields, defaultField{
		key: key,
		val: val,
	})

	return l
}

func (l *log4GoLogger) GetFilter() (filter *logutils.LevelFilter) {
	return l.filter
}

func (l *log4GoLogger) SetFilter(filter *logutils.LevelFilter) {
	l.filter = &logutils.LevelFilter{
		Levels:   logLevel,
		MinLevel: LogLevelInfo,
		Writer:   filter.Writer,
	}

	switch filter.MinLevel {
	case LogLevelTrace:
		fallthrough
	case LogLevelDebug:
		fallthrough
	case LogLevelInfo:
		fallthrough
	case LogLevelWarn:
		fallthrough
	case LogLevelError:
		l.filter.MinLevel = filter.MinLevel
	default:
		l.filter.MinLevel = LogLevelInfo
	}

	log.SetOutput(l.filter)
}

func (l *log4GoLogger) Trace(ctx context.Context, format string, v ...interface{}) {
	format = "[TRCE]" + format
	l.printf(ctx, format, v...)
}

func (l *log4GoLogger) Debug(ctx context.Context, format string, v ...interface{}) {
	format = "[DBUG]" + format
	l.printf(ctx, format, v...)
}

func (l *log4GoLogger) Info(ctx context.Context, format string, v ...interface{}) {
	format = "[INFO]" + format
	l.printf(ctx, format, v...)
}

func (l *log4GoLogger) Warn(ctx context.Context, format string, v ...interface{}) {
	format = "[WARN]" + format
	l.printf(ctx, format, v...)
}

func (l *log4GoLogger) Error(ctx context.Context, format string, v ...interface{}) {
	format = "[EROR]" + format
	l.printf(ctx, format, v...)
}

func (l *log4GoLogger) printf(ctx context.Context, format string, v ...interface{}) {
	var id string
	var line string
	var x, y int

	line = fmt.Sprintf(format, v...)

	var logStr string
	if l.opts.WithId {
		id = getIdFromContext(ctx, l.opts.IdName)
		var tag = "[" + id + "]"
		// Check for a log level
		x = strings.Index(line, "[")
		if x >= 0 {
			y = strings.Index(line[x:], "]")
			if y >= 0 {
				logStr = line[:x+y+1] + tag + line[x+y+1:]
			}
		}
	} else {
		logStr = line
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
	log.Output(4, logStr)
}

func getIdFromContext(ctx context.Context, idName string) string {
	id, ok := ctx.Value(idName).(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(id)
}
