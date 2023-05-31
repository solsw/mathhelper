package mathhelper

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Split splits 'n' into 'parts' as equal as possible integer parts.
// Slice with parts' lengths is returned.
func Split[T constraints.Integer](n, parts T) ([]T, error) {
	if n <= 0 {
		return nil, errors.New("wrong n")
	}
	if parts <= 0 {
		return nil, errors.New("wrong parts")
	}
	if n < parts {
		return nil, errors.New("n less than parts")
	}
	var tt []T
	for n > 0 {
		part := n / parts
		if n%parts > 0 {
			part++
		}
		tt = append(tt, part)
		n -= part
		parts--
	}
	return tt, nil
}
