package log4go_test

import (
	"bytes"
	"context"
	"log"
	"testing"

	"github.com/hashicorp/logutils"
	"github.com/liuliqiang/log4go"
)

func TestDefaultDebug(t *testing.T) {
	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelDebug,
		Writer:   buf,
	}
	log4go.SetFilter(filter)
	log.SetFlags(log.LstdFlags & 0x0)

	log4go.Debug(context.Background(), "1")
	log4go.Info(context.Background(), "3")
	log4go.Warn(context.Background(), "5")
	log4go.Error(context.Background(), "7")

	result := buf.String()
	expected := "[DBUG]1\n[INFO]3\n[WARN]5\n[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}

func TestDefaultInfo(t *testing.T) {
	log.SetFlags(log.LstdFlags & 0x0)

	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelInfo,
		Writer:   buf,
	}
	log4go.SetFilter(filter)

	log4go.Debug(context.Background(), "1")
	log4go.Info(context.Background(), "3")
	log4go.Warn(context.Background(), "5")
	log4go.Error(context.Background(), "7")

	result := buf.String()
	expected := "[INFO]3\n[WARN]5\n[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}

func TestDefaultWarn(t *testing.T) {
	log.SetFlags(log.LstdFlags & 0x0)

	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelWarn,
		Writer:   buf,
	}
	log4go.SetFilter(filter)

	log4go.Debug(context.Background(), "1")
	log4go.Info(context.Background(), "3")
	log4go.Warn(context.Background(), "5")
	log4go.Error(context.Background(), "7")

	result := buf.String()
	expected := "[WARN]5\n[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}

func TestDefaultError(t *testing.T) {
	log.SetFlags(log.LstdFlags & 0x0)

	buf := new(bytes.Buffer)
	filter := &logutils.LevelFilter{
		MinLevel: log4go.LogLevelError,
		Writer:   buf,
	}
	log4go.SetFilter(filter)

	log4go.Debug(context.Background(), "1")
	log4go.Info(context.Background(), "3")
	log4go.Warn(context.Background(), "5")
	log4go.Error(context.Background(), "7")

	result := buf.String()
	expected := "[EROR]7\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}
