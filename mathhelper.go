// Package mathhelper contains 'math' package helpers.
package mathhelper

import (
	"math"

	"golang.org/x/exp/constraints"
)

// IsEven reports whether 'i' is even.
func IsEven[T constraints.Integer](i T) bool {
	return (i & 1) == 0
}

// Abs returns the absolute value of 'x'.
func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Frac returns the fractional part of 'f'.
func Frac[T constraints.Float](f T) T {
	_, r := math.Modf(float64(f))
	return T(r)
}

// IntMin returns the smaller of 'm' or 'n'.
func IntMin[T constraints.Integer](m, n T) T {
	if m < n {
		return m
	}
	return n
}

// IntMax returns the larger of 'm' or 'n'.
func IntMax[T constraints.Integer](m, n T) T {
	if m > n {
		return m
	}
	return n
}
