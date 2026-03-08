package log4go

import (
	"context"
	"fmt"
	"log"
	"runtime"
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

func (l *log4GoLogger) SetWithCaller(enabled bool) {
	l.opts.WithCaller = enabled
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
	line := fmt.Sprintf(format, v...)

	// Build extra tags to insert after the level tag: [caller][id][fields...]
	var extraTags string
	if l.opts.WithCaller {
		extraTags += "[" + callerInfo() + "]"
	}
	if l.opts.WithId {
		extraTags += "[" + getIdFromContext(ctx, l.opts.IdName) + "]"
	}
	for _, field := range l.fields {
		extraTags += "[" + field[0] + "=" + field[1] + "]"
	}

	logStr := line
	if extraTags != "" {
		if x := strings.Index(line, "["); x >= 0 {
			if y := strings.Index(line[x:], "]"); y >= 0 {
				logStr = line[:x+y+1] + extraTags + line[x+y+1:]
			}
		}
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.stdLogger.Output(4, logStr)
}

// callerInfo returns "file.go:line" of the first call site outside this package.
func callerInfo() string {
	pcs := make([]uintptr, 10)
	// skip: runtime.Callers, callerInfo, printf
	n := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, "github.com/liuliqiang/log4go.") {
			file := frame.File
			if idx := strings.LastIndex(file, "/"); idx >= 0 {
				file = file[idx+1:]
			}
			return fmt.Sprintf("%s:%d", file, frame.Line)
		}
		if !more {
			break
		}
	}
	return "??:0"
}

func getIdFromContext(ctx context.Context, idName interface{}) string {
	id, ok := ctx.Value(idName).(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(id)
}
