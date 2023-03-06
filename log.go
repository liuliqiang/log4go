package log4go

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hashicorp/logutils"
)

type log4GoLogger struct {
	fields    [][2]string // add pool support
	stdLogger *log.Logger
	mutex     sync.Mutex
	filter    *logutils.LevelFilter
	opts      *LoggerOpts
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

	l.stdLogger.SetOutput(l.filter)
}

func (l *log4GoLogger) SetFlags(flag int) {
	l.stdLogger.SetFlags(flag)
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

func (l *log4GoLogger) clone() *log4GoLogger {
	newLog := &log4GoLogger{
		stdLogger: l.stdLogger,
		filter:    l.filter,
		opts:      l.opts,
	}
	newLog.fields = fieldPool.GetFields(cap(l.fields))
	newLog.fields = append(newLog.fields, l.fields...)
	return newLog
}

func (l *log4GoLogger) WithField(field, val string) Logger {
	newLogger := l.clone() // copy on write?
	newLogger.fields = append(newLogger.fields, [2]string{field, val})

	return newLogger
}

func (l *log4GoLogger) printf(ctx context.Context, format string, v ...interface{}) {
	var id string
	var line string
	var x, y int

	line = fmt.Sprintf(format, v...)

	var logStr string
	if l.opts.WithId || len(l.fields) > 0 {
		id = getIdFromContext(ctx, l.opts.IdName)
		var tag string
		if l.opts.WithId {
			tag = "[" + id + "]"
		}
		var fieldStr string
		for _, field := range l.fields {
			fieldStr += "[" + field[0] + "=" + field[1] + "]"
		}
		// Check for a log level
		x = strings.Index(line, "[")
		if x >= 0 {
			y = strings.Index(line[x:], "]")
			if y >= 0 {
				logStr = line[:x+y+1] + tag + fieldStr + line[x+y+1:]
			}
		}
	} else {
		logStr = line
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.stdLogger.Output(4, logStr)
}

func getIdFromContext(ctx context.Context, idName interface{}) string {
	id, ok := ctx.Value(idName).(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(id)
}
