package errors

import (
	"runtime"
)

var _ RootCauseError = (*rootCauseError)(nil)

type rootCauseError struct {
	caller *runtime.Frame
	err    error
}

func RootCause(err error) error {
	if err == nil {
		return nil
	}
	return createRootCause(err)
}

func (r *rootCauseError) Error() string {
	return r.err.Error()
}

func (r *rootCauseError) At() *runtime.Frame {
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
	err := &rootCauseError{err: cause}

	lvl := 4 // stack level
	if caller, ok := getCaller(lvl); ok {
		err.caller = &caller
	}

	return err
}
