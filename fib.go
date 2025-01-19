package pe

// Fibs returns the first n fibonacci numbers
func Fibs(n int) []int {
	arr := []int{1, 2}
	if n < 3 {
		return arr[:n]
	}
	x, y := 1, 2
	for i := 0; i < n-2; i++ {
		x, y = y, x+y
		arr = append(arr, y)
	}
	return arr
}
