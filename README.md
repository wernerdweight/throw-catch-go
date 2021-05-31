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

**Use with defer**

```go
package main

import "github.com/wernerdweight/throw-catch-go"

func doSomething() {
    defer performance.CreateSpan("my-transaction", "doSomething").Finish()
    // put your code here, the transaction/span is created now
    // due to how defer works, the 'Finish' method will be called at the end
}

func main() {
    defer performance.CreateTransaction("my-transaction", "main").Finish()
    doSomething()
}
```

License
-------
This package is under the MIT license. See the complete license in the root directory of the bundle.
