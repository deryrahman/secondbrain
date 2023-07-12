package errors

import "runtime"

type RootCauseError interface {
	Error() string
	StatusCode() int
	At() *runtime.Frame
}

type JoinError interface {
	Error() string
	Unwrap() []error
}
