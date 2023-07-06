package errors

import "runtime"

type RootCauseError interface {
	Error() string
	At() *runtime.Frame
}

type JoinErr interface {
	Error() string
	Unwrap() []error
}
