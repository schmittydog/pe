package pe

// BSIsIn binsearches for int value in sorted array
func BSIsIn(arr []int, value int) bool {
	if len(arr) == 0 {
		return false
	}
	l, r := 0, len(arr) - 1
	for r > l + 1 {
		m := l + (r-l)/2
		if arr[m] == value {
			return true
		} else if arr[m] > value {
			r = m
		} else {
			l = m
		}
	}
	return arr[l] == value || arr[r] == value
}
