package main

import (
	"errors"

	"github.com/liuliqiang/log4go"
)

func main() {
	log4go.Trace("I am trace log")
	log4go.Debug("I am debug log")
	log4go.Info("Web server is started at %s:%d", "127.0.0.1", 80)
	log4go.Warn("Get an empty http request")
	log4go.Error("Failed to query record from db: %v", errors.New("db error"))
}
