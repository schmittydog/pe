package pe

func Base(b, n int) []int {
	arr := []int{}
	for n > 0 {
		arr = append(arr, n%b)
		n /= b
	}
	return arr
}
