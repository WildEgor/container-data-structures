package stack

import "fmt"

var _ error = (*InvalidOptionError)(nil)

// InvalidOptionError may be returned
type InvalidOptionError struct{}

func (e *InvalidOptionError) Error() string {
	return fmt.Sprint("")
}

// errInvalidOption signal misconfiguration
func errInvalidOption() { panic(&InvalidOptionError{}) }
