package errors

var _ JoinErr = (*joinError)(nil)

type joinError struct {
	errs []error
}

func Join(errs ...error) error {
	e := &joinError{
		errs: []error{},
	}
	for _, err := range errs {
		if err != nil {
			if joinErr, ok := err.(*joinError); ok {
				e.errs = append(e.errs, Join(joinErr.errs...))
			} else {
				e.errs = append(e.errs, err)
			}
		}
	}
	return e
}

func (e *joinError) Error() string {
	var b []byte
	for i, err := range e.errs {
		if i > 0 {
			b = append(b, '\n')
		}
		b = append(b, err.Error()...)
	}
	return string(b)
}

func (e *joinError) Unwrap() []error {
	return e.errs
}
