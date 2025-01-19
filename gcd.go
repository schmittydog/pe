package pe

// Returns GCD of two integers
func GCD(a, b int) int {
	if b%a == 0 {
		return a
	}
	return GCD(b%a,a)
}
