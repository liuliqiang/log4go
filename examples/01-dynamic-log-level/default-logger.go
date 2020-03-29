package main

import (
	"net/http"
	"time"

	"github.com/liuliqiang/log4go"
)

func main() {
	var err error

	go func() {
		for {
			log4go.Debug("I am debug log")
			log4go.Info("I am info log")
			time.Sleep(time.Second * 5)
		}
	}()

	http.HandleFunc("/logs", func(resp http.ResponseWriter, req *http.Request) {
		var queryParam =  req.URL.Query()
		if logLevels, exists := queryParam["level"]; !exists  || len(logLevels) != 1{
			resp.Write([]byte("log level not found"))
		} else {
			switch logLevels[0] {
			case "debug":
				log4go.SetLevel(log4go.LogLevelDebug)
			case "info":
				log4go.SetLevel(log4go.LogLevelInfo)
			case "warning":
				log4go.SetLevel(log4go.LogLevelWarn)
			case "error":
				log4go.SetLevel(log4go.LogLevelError)
			default:
				resp.Write([]byte("bad log level " + logLevels[0]))
			}
		}
		return
	})
	log4go.Info("Server start!")
	if err = http.ListenAndServe(":9091", nil); err != nil {
		panic(err)
	}
	log4go.Info("Server Exit!")
}
