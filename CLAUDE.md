# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build -v ./...

# Test all
go test -v ./...

# Run a single test
go test -v -run TestName ./...
```

## Architecture

log4go is a Go logging library wrapping hashicorp/logutils with level filtering, named logger registry, context-based ID propagation, and structured field support.

### Core files

| File | Role |
|------|------|
| `init.go` | Log level constants, global logger registry (`loggerMap`), `LoggerOpts` struct |
| `log.go` | `Logger` interface, `log4GoLogger` implementation (wraps `log.Logger` + `LevelFilter`) |
| `default.go` | Package-level convenience functions delegating to the default logger |
| `pool.go` | `fieldsPool` — object pool for `[][2]string` slices to reduce GC pressure |

### Key design decisions

**Named logger registry** — `NewLogger(name, opts)` caches logger instances in `loggerMap` (protected by `loggerMu`). Same name always returns the same instance. Default logger uses `""` as name.

**Immutable field chaining** — `WithField(key, val)` clones the logger with an extra field appended; the original is unchanged. Fields are stored as `[][2]string` pairs and emitted as `[key=val]` prefixes.

**Context ID propagation** — When `LoggerOpts.WithId = true`, each log call extracts a request ID from context using `IdName` as the key and prepends it to the message. Enables distributed tracing without manual threading.

**Level filtering** — Five levels: `TRCE`, `DBUG`, `INFO`, `WARN`, `EROR`. Backed by `logutils.LevelFilter`; minimum level is configurable at runtime via `SetLevel` / `SetFilter`.

**Object pooling** — `fieldsPool` caches slices by capacity (rounded to multiples of 10) to reduce allocations in high-throughput scenarios.
