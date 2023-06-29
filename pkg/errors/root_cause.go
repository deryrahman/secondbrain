package errors

import (
	"runtime"
)

var _ RootCauseErr = (*rootCause)(nil)

type rootCause struct {
	caller *runtime.Frame
	msg    string
}

func RootCause(err error) error {
	return createRootCause(err)
}

func (r *rootCause) Error() string {
	return r.msg
}

func (r *rootCause) At() *runtime.Frame {
	return r.caller
}

func getCaller(callerSkip int) (fr runtime.Frame, ok bool) {
	pcs := make([]uintptr, 1) // alloc 1 times
	num := runtime.Callers(callerSkip, pcs)
	if num < 1 {
		return
	}

	f, _ := runtime.CallersFrames(pcs).Next()
	return f, f.PC != 0
}

func createRootCause(cause error) error {
	err := &rootCause{msg: cause.Error()}

	if caller, ok := getCaller(4); ok {
		err.caller = &caller
	}

	return err
}
