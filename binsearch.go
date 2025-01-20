package pe

// BSIsIn binsearches for int value in sorted array
// ([]int{1,3,7,33,67,101,199}, 3)) -> true
func BSIsIn(arr []int, value int) bool {
	if len(arr) == 0 {
		return false
	}
	l, r := 0, len(arr) - 1
	for r > l {
		m := (r + l) / 2
		if arr[m] < value {
			l = m + 1	
		} else {
			r = m
		}
	}
	return arr[l] == value
}


// BSRight bisect right
func BSRight(arr []int, value int) int {
        if len(arr) == 0 {
                return 0
        }
        l, r := 0, len(arr)
        for r > l {
                m := (r + l) / 2
                if arr[m] >= value {
                        r = m - 1
                } else {
                        l = m
                }
        }
        return l
}

// BSLeft bisect left
func BSLeft(arr []int, value int) int {
        if len(arr) == 0 {
                return 0
        }
        l, r := 0, len(arr)-1
        for r > l {
                m := (r + l) / 2
                if arr[m] < value {
                        l = m + 1
                } else {
                        r = m
                }
        }
        return l
}

