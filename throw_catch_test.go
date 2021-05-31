package throw_catch_go_test

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/wernerdweight/throw-catch-go/catch"
	"github.com/wernerdweight/throw-catch-go/throw"
	"testing"
)

const (
	testErrorMessage       = "Test error message"
	testNestedErrorMessage = "Nested error message"
)

func throwTestHandler(t *testing.T) func(error) {
	assertion := assert.New(t)
	return func(err error) {
		assertion.Equal(testErrorMessage, err.Error())
	}
}

func throwNestedTestHandler(t *testing.T) func(error) {
	assertion := assert.New(t)
	return func(err error) {
		assertion.Equal(testErrorMessage, errors.Cause(err).Error())
		assertion.Equal(fmt.Sprintf("%s: %s", testNestedErrorMessage, testErrorMessage), err.Error())
	}
}

func throwTestPanic() {
	throw.Throw(testErrorMessage)
}

func throwNestedTestPanic() {
	throw.ThrowNested(errors.New(testErrorMessage), testNestedErrorMessage)
}

func TestThrowCatch(t *testing.T) {
	defer catch.Catch(throwTestHandler(t))
	throwTestPanic()
}

func TestThrowNestedCatch(t *testing.T) {
	defer catch.Catch(throwNestedTestHandler(t))
	throwNestedTestPanic()
}
