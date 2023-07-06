package errors

import (
	"errors"
	"fmt"
)

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w", err)
}
