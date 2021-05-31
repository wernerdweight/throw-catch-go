Throw - Catch for Go
====================================

This package brings basic throw-catch functionality for Go.
Please be aware that you should not use this the same way you would use throw-catch in other languages!
This package helps with handling of critical errors. Don't use it for things like input validation etc.
Be sure to read [Error handling and Go](https://blog.golang.org/error-handling-and-go) and [Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover). 

[![Build Status](https://www.travis-ci.com/wernerdweight/throw-catch-go.svg?branch=master)](https://www.travis-ci.com/wernerdweight/throw-catch-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/wernerdweight/throw-catch-go)](https://goreportcard.com/report/github.com/wernerdweight/throw-catch-go)
[![GoDoc](https://godoc.org/github.com/wernerdweight/throw-catch-go?status.svg)](https://godoc.org/github.com/wernerdweight/throw-catch-go)
[![go.dev](https://img.shields.io/badge/go.dev-pkg-007d9c.svg?style=flat)](https://pkg.go.dev/github.com/wernerdweight/throw-catch-go)


Installation
------------

### 1. Installation

```bash
go get github.com/wernerdweight/throw-catch-go
```

Configuration
------------

The package needs no configuration.

Usage
------------

**Basic usage**

```go
package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/wernerdweight/throw-catch-go/catch"
	"github.com/wernerdweight/throw-catch-go/throw"
	"log"
)

func errorHandler(err error) {
	// do whatever you want with the error here
	log.Printf("%+v", err)
	// it might be a good idea to log the error to Sentry for example
	sentry.CaptureException(err)
}

func doSomething() {
    throw.Throw("Something terrible happened")
	// nothing else will get executed
}

func main() {
	defer catch.Catch(errorHandler)
    doSomething()
}
```

**Nested errors**

```go
package main

import (
	"github.com/pkg/errors"
	"github.com/wernerdweight/throw-catch-go/catch"
	"github.com/wernerdweight/throw-catch-go/throw"
	"log"
)

func errorHandler(err error) {
	// do whatever you want with the error here
	log.Printf("%+v", err)
}

func doSomething() {
    var err = errors.New("Something that caused the terrible thing to happen")
	throw.ThrowNested(err, "Something terrible happened")
	// nothing else will get executed
}

func main() {
	defer catch.Catch(errorHandler)
    doSomething()
}
```

**Get the error outside of Catch**

```go
package main

import (
	"github.com/wernerdweight/throw-catch-go/catch"
	"github.com/wernerdweight/throw-catch-go/throw"
	"log"
)

func errorHandler(returnableErr *error) func(error) {
	return func (err error) {
		// you can get the error outside of the Catch method
		// (e.g. if SomeFunction is a gRPC procedure)
		*returnableErr = err
		// do whatever you want with the error here
		log.Printf("%+v", err)
	}
}

func doSomething() {
	throw.Throw("Something terrible happened")
	// nothing else will get executed
}

func SomeFunction() (err error) {
	defer catch.Catch(errorHandler(&err))
	doSomething()
	return
}
```

License
-------
This package is under the MIT license. See the complete license in the root directory of the bundle.
