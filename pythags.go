package pe

import (
	"sort"
)

// PythagPrims returns a chan of []int of all pythag primitives
// containing legs a, b less than m
// where a < b < c
func PythagPrims(m int) chan []int {
	ch := make(chan []int)
	go pythagHelper(m, ch)
	return ch
}

func pythagHelper(m int, ch chan []int) {
	for i := 1; i*i+2*i <= m; i += 2 {
		for j := i + 2; i*j <= m; j += 2 {
			if GCD(i, j) != 1 {
				continue
			}
			arr := []int{i * j, (j*j - i*i) / 2, (j*j + i*i) / 2}
			sort.Ints(arr)
			if arr[1] > m {
				break
			}
			ch <- arr
		}
	}
	close(ch)
}
