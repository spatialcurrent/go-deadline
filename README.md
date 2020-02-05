[![CircleCI](https://circleci.com/gh/spatialcurrent/go-deadline/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-deadline/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-deadline)](https://goreportcard.com/report/spatialcurrent/go-deadline)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-deadline?status.svg)](https://godoc.org/github.com/spatialcurrent/go-deadline) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-deadline/blob/master/LICENSE)

# go-deadline

## Description

**go-deadline** is a library to create deadlines for goroutines and programs.  This package is used as a safeguard to prevent a goroutine, test, or program from exhausting resources or otherwise running beyond the expected duration of time.  You can also use this library for automatically introducing some chaos into your container environment to test failover and network resilience.

# Usage

**Go**

You can import **go-deadline** as a library with:

```go
import (
  "time"
  "github.com/spatialcurrent/go-deadline/pkg/deadline"
)
```

You can create a deadline near the start of your program.

```go
func main() {

  ...

  d, err := deadline.New(5*time.Second, deadline.ExitError)
  if err != nil {
    return fmt.Errorf("error creating deadline: %w", err)
  }
  err := d.Start()
  if err != nil {
    return fmt.Errorf("error starting deadline: %w", err)
  }
  // deadline is no running in a separate goroutine

  ...
}
```

If you do not care to handle errors yourself, you can use the `deadline.MustStart` function with the default `deadline.ExitError` function.

```go
func main() {
  ...
  deadline.MustStart(context.Background(), 5*time.Second, deadline.ExitError)
  ...
}
```

Alternatively, if you wish to post a custom error to `stderr`, you can provide a custom function as below.

```go
func main() {
  ...
  deadline.MustStart(context.Background(), 5*time.Second, func(ctx context.Context) {
    fmt.Fprintln(os.Stderr, "deadline reached")
    os.Exit(1)
  })
  ...
}
```

See [deadline](https://godoc.org/github.com/spatialcurrent/go-deadline/pkg/deadline) in GoDoc for further API documentation.

# Testing

To run Go tests use `make test_go` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-deadline/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
