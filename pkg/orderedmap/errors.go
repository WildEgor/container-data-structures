package orderedmap

import "fmt"

var _ error = (*KeyNotFoundError[any])(nil)

// KeyNotFoundError may be returned by functions in this package when they're called with keys that are not present in the map.
type KeyNotFoundError[K comparable] struct {
	MissingKey K
}

func (e *KeyNotFoundError[K]) Error() string {
	return fmt.Sprintf("missing key: %v", e.MissingKey)
}

var _ error = (*InvalidOptionError)(nil)

// InvalidOptionError may be returned
type InvalidOptionError struct{}

func (e *InvalidOptionError) Error() string {
	return fmt.Sprint("when using orderedmap.New[K,V]() with options, either provide one or several Options[K, V]; or a single integer which is then interpreted as a capacity hint, à la make(map[K]V, capacity).")
}

// errInvalidOption signal misconfiguration
func errInvalidOption() { panic(&InvalidOptionError{}) }
