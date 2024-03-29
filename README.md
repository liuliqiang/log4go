# log4go

[![Go Report Card](https://goreportcard.com/badge/github.com/liuliqiang/log4go)](https://goreportcard.com/report/github.com/liuliqiang/log4go)
[![GoDoc](https://pkg.go.dev/badge/github.com/liuliqiang/log4go?status.svg)](https://pkg.go.dev/github.com/liuliqiang/log4go?tab=doc)
[![codecov](https://codecov.io/gh/liuliqiang/log4go/branch/master/graph/badge.svg)](https://codecov.io/gh/liuliqiang/log4go)
[![Sourcegraph](https://sourcegraph.com/github.com/liuliqiang/log4go/-/badge.svg)](https://sourcegraph.com/github.com/liuliqiang/log4go?badge)
[![Open Source Helpers](https://www.codetriage.com/liuliqiang/log4go/badges/users.svg)](https://www.codetriage.com/liuliqiang/log4go)
[![Release](https://img.shields.io/github/release/liuliqiang/log4go.svg?style=flat-square)](https://github.com/liuliqiang/log4go/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/liuliqiang/log4go)](https://www.tickgit.com/browse?repo=github.com/liuliqiang/log4go)

A log library for golang based on hashicorp/logutils

## Why log4go

Yes, it another logging library for Go program language. Why do I create a new logging library for Go when there are so many popular logging library. such as:

- [hashicorp/logutils](https://github.com/hashicorp/logutils)
- [sirupsen/logrus](https://github.com/Sirupsen/logrus)
- [golang/glog](https://github.com/golang/glog)
- [so on...](https://github.com/avelino/awesome-go#logging)

For my daily use, i found several points are important for logging:

- level for logging shoud be easy to use
- content for logging should be friendly for human reading
- logs should be easy to check

so i create log4go

## Quick start

1. Step 01: get the library

    ```
    $ go get -u github.com/liuliqiang/log4go
    ```

2. Step 02: try in code:

    ```
    func main() {
    	log4go.Debug("I am debug log")
    	log4go.Info("Web server is started at %s:%d", "127.0.0.1", 80)
    	log4go.Warn("Get an empty http request")
    	log4go.Error("Failed to query record from db: %v", errors.New("db error"))
    }
    ```

3. Step 03: Run it!

    ```
    $ go run main.go
    2019/08/10 00:02:18 [INFO]Web server is started at 127.0.0.1:80
    2019/08/10 00:02:18 [WARN]Get an empty http request
    2019/08/10 00:02:18 [EROR]Failed to query record from db: db error
    ```

## Dynamic change log level

Just add a endpoint(such as http/grpc/...) to change log level such as:

```
http.HandleFunc("/logs", func(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Query()["level"][0] {
	case "debug":
		log4go.SetLevel(log4go.LogLevelDebug)
	case "info":
		log4go.SetLevel(log4go.LogLevelInfo)
	case "warning":
		log4go.SetLevel(log4go.LogLevelWarn)
	case "error":
		log4go.SetLevel(log4go.LogLevelError)
	}
	return
})
```

## Learn more...

- To be continue...
- More examples at [Examples](./examples)

