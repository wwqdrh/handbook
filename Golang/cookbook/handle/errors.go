package handle

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// OriginalError returns the error original error
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError calls OriginalError and
// forwards the error along after wrapping.
func PassThroughError() error {
	err := OriginalError()
	// no need to check error
	// since this works with nil
	return errors.Wrap(err, "in passthrougherror")
}

// FinalDestination deals with the error
// and doesn't forward it
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// we log because an unexpected error occurred!
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}

// WrappedError demonstrates error wrapping and
// annotating an error
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

// ErrorTyped is a error we can check against
type ErrorTyped struct {
	error
}

// Wrap shows what happens when we wrap an error
func Wrap() {
	e := errors.New("standard error")

	fmt.Println("Regular Error - ", WrappedError(e))

	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))

	fmt.Println("Nil -", WrappedError(nil))

}

// Unwrap will unwrap an error and do
// type assertion to it
func Unwrap() {

	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	// we can handle many error types
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}

// StackTrace will print all the stack for
// the error
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}
