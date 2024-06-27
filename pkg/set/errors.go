package set

import "fmt"

var _ error = (*EmptySetError)(nil)

// EmptySetError signal about empty Set
type EmptySetError struct {
}

func (e *EmptySetError) Error() string {
	return fmt.Sprintf("empty set")
}
