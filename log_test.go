package log4go_test

import (
	"bytes"
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
	log.SetFlags(log.LstdFlags & 0x0)

	logger.Debug("1")
	logger.Info("3")
	logger.Warn("5")
	logger.Error("7")

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
	log.SetFlags(log.LstdFlags & 0x0)

	logger.Debug("1")
	logger.Info("3")
	logger.Warn("5")
	logger.Error("7")

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
	log.SetFlags(log.LstdFlags & 0x0)

	logger.Debug("1")
	logger.Info("3")
	logger.Warn("5")
	logger.Error("7")

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
	log.SetFlags(log.LstdFlags & 0x0)

	logger.Debug("1")
	logger.Info("3")
	logger.Warn("5")
	logger.Error("7")

	result := buf.String()
	expected := "[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}
