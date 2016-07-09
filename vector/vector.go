// Package vector implements vector operations in mathematics.
package vector

import "math"

// Vector is a float64 slice.
type Vector []float64

// Len returns the length of v.
func (v Vector) Len() int {
	return len(v)
}

// Norm returns the vector norm of v.
func (v Vector) Norm() float64 {
	return math.Sqrt(float64(Dot(v, v)))
}

// New returns a new vector with the size.
func New(size int) Vector {
	return make([]float64, size)
}

// Dot returns the inner product of v1 and v2. If the size mismatches,
// it assumes that the smaller vector has additional 0s.
func Dot(v1, v2 Vector) float64 {
	smaller := v1.Len()
	if len(v2) < smaller {
		smaller = len(v2)
	}
	result := 0.0
	for i := 0; i < smaller; i++ {
		result += v1[i] * v2[i]
	}
	return result
}
