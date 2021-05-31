package catch

import "github.com/pkg/errors"

// Catch recovers from thrown errors (panics)
func Catch(handler func(error)) {
	if r := recover(); nil != r {
		var stackedErr = errors.WithStack(r.(error))
		handler(stackedErr)
	}
}
