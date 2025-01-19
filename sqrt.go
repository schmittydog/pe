package pe

import (
	"math"
)

// Returns the integer square root of n
func IntSqrt(n int) int {
	f := float64(n)
	froot := math.Sqrt(f)
	return int(froot)
}
