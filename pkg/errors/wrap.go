package errors

import "fmt"

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w", err)
}
