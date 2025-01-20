package pe

import (
	"log"
)

// Returns GCD of two integers
func GCD(a, b int) int {
	if b%a == 0 {
		return a
	}
	return GCD(b%a, a)
}

// Returns GCD of a slice of numbers
func GCDArray(arr []int) int {
	if len(arr) == 0 {
		log.Fatalln("Send empty slice")
	}
	gcd := arr[0]
	for i := 1; i < len(arr); i++ {
		gcd = GCD(gcd, arr[i])
	}
	return gcd
}
