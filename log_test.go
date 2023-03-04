package log4go_test

import (
	"bytes"
	"context"
	"io"
	"log"
	"testing"

	"github.com/liuliqiang/log4go"

	"github.com/hashicorp/logutils"
)

func TestLevelFilter_impl(t *testing.T) {
	var _ io.Writer = new(logutils.LevelFilter)
}

func TestLevelDebug(t *testing.T) {
	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelDebug,
		Writer:   buf,
	}
	logger := log4go.NewLogger("", nil)
	logger.SetFilter(filter)
	log.SetFlags(0x0)

	logger.Debug(context.Background(), "1")
	logger.Info(context.Background(), "3")
	logger.Warn(context.Background(), "5")
	logger.Error(context.Background(), "7")

	result := buf.String()
	expected := "[DBUG]1\n[INFO]3\n[WARN]5\n[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}

func TestLevelInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelInfo,
		Writer:   buf,
	}
	logger := log4go.NewLogger("", nil)
	logger.SetFilter(filter)
	log.SetFlags(0x0)

	logger.Debug(context.Background(), "1")
	logger.Info(context.Background(), "3")
	logger.Warn(context.Background(), "5")
	logger.Error(context.Background(), "7")

	result := buf.String()
	expected := "[INFO]3\n[WARN]5\n[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}

func TestLevelWarn(t *testing.T) {
	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelWarn,
		Writer:   buf,
	}
	logger := log4go.NewLogger("", nil)
	logger.SetFilter(filter)
	log.SetFlags(0x0)

	logger.Debug(context.Background(), "1")
	logger.Info(context.Background(), "3")
	logger.Warn(context.Background(), "5")
	logger.Error(context.Background(), "7")

	result := buf.String()
	expected := "[WARN]5\n[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}

func TestLevelError(t *testing.T) {
	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelError,
		Writer:   buf,
	}
	logger := log4go.NewLogger("", nil)
	logger.SetFilter(filter)
	log.SetFlags(0x0)

	logger.Debug(context.Background(), "1")
	logger.Info(context.Background(), "3")
	logger.Warn(context.Background(), "5")
	logger.Error(context.Background(), "7")

	result := buf.String()
	expected := "[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}

func TestNewLogger(t *testing.T) {
	logger1, buf1 := newLogger("logger1")
	logger1.Error(context.Background(), "1")

	logger2, buf2 := newLogger("logger2")
	logger2.Error(context.Background(), "2")

	if buf1.String() != "logger1[EROR]1\n" {
		t.Fatalf("bad logger1 output: %s", buf1.String())
	}
	if buf2.String() != "logger2[EROR]2\n" {
		t.Fatalf("bad logger2 output: %s", buf2.String())
	}
}

func newLogger(name string) (log4go.Logger, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelDebug,
		Writer:   buf,
	}
	logger := log4go.NewLogger(name, nil)
	logger.SetFlags(0x0)
	logger.SetFilter(filter)
	return logger, buf
}
