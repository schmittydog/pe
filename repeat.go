package pe

func Repeats(n int) bool {
	for n%2 == 0 {
		n /= 2
	}
	for n%5 == 0 {
		n /= 5
	}
	return n != 1
}

func RepeatCount(n int) int {
	if !Repeats(n) {
		return 0
	}

	start := 10
	for start < n {
		start *= 10
	}
	m := map[int]int{start: 0}
	index := 0
	for {
		index += 1
		start %= n
		start *= 10
		if _, ok := m[start]; ok {
			return index - m[start]
		}
		m[start] = index
	}
	return 0
}
