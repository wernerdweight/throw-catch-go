package throw

import (
	"github.com/pkg/errors"
)

// Throw issues a panic with custom error message
func Throw(message string) {
	panic(errors.New(message))
}

// ThrowNested issues panic with wrapped error
func ThrowNested(err error, message string) {
	panic(errors.Wrap(err, message))
}
