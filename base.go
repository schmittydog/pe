package pe

func NumToArr(n int) []int {
	arr := []int{}
	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	return arr
}

func ArrToNum(arr []int) int {
	num := 0
	for _, n := range arr {
		num *= 10
		num += n
	}
	return num
}

func Base(b, n int) []int {
	arr := []int{}
	for n > 0 {
		arr = append(arr, n%b)
		n /= b
	}
	return arr
}
		
