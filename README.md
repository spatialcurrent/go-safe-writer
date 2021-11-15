[![CircleCI](https://circleci.com/gh/spatialcurrent/go-safe-writer/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-safe-writer/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-safe-writer)](https://goreportcard.com/report/spatialcurrent/go-safe-writer)  [![PkgGoDev](https://pkg.go.dev/badge/github.com/spatialcurrent/go-safe-writer)](https://pkg.go.dev/github.com/spatialcurrent/go-safe-writer) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-safe-writer/blob/master/LICENSE.md)

# go-safe-writer

## Description

**go-safe-writer** is used to wrap a writer to enable it to be safely used by concurrent goroutines.  The `NewSafeWriter` is used to create a `SafeWriter`, which uses an internal `sync.Mutex` to ensure only one goroutine is writing at a time.  The `NewSafeWriteCloser` is used to create a `SafeWriterCloser`.  The `SafeWriteCloser` also locks the mutex when calling the underlying close function.  The `NewSafeWriteFlusher` is used to create a `SafeWriteFlusher`.  The `SafeWriteFlusher` also locks the mutex when calling the underlying `Flush() error` function.  The `NewSafeWriteFlushCloser` is used to create a `SafeWriteFlushCloser` that locks the mutex when calling `Write`, `Flush`, and `Closer`.

# Usage

**Go**

You can import **go-safe-writer** as a library with:

```go
import (
  "github.com/spatialcurrent/go-safe-writer/pkg/gsw"
)
```

Use the `NewSafeWriter` function to create a safe `SafeWriter` that implements `io.Writer`.

```go
return gsw.NewSafeWriter(writer)
```

Use the `NewSafeWriteCloser` function to create a safe `SafeWriteCloser` that implements `io.WriteCloser`.

```go
return gsw.NewSafeWriteCloser(writer)
```

If you choose to lock the writer yourself, you need to use the `WriteUnsafe` method otherwise you will enter a deadlock.

```go
gsw.NewSafeWriter(writer)
writer.Lock()
writer.WriteUnsafe([]byte("hello"))
writer.WriteUnsafe([]byte("world"))
writer.Unlock()
```

See [gsw](https://pkg.go.dev/github.com/spatialcurrent/go-safe-writer/pkg/gsw) for information on how to use Go API.

# Testing

To run Go tests use `make test_go` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-safe-writer/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
