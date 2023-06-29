package errors

import "runtime"

type RootCauseErr interface {
	Error() string
	At() *runtime.Frame
}

type JoinErr interface {
	Error() string
	Unwrap() []error
}
