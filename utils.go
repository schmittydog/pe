package pe

// Copies an array from 0 -> index
func CopyArr(arr []int, index int) []int {
	newArr := make([]int, index+1)
	for i := 0; i <= index; i++ {
		newArr[i] = arr[i]
	}
	return newArr
}

// Converts a number to an array of decimal digits
// 12345 -> [1 2 3 4 5]
func NumToArr(n int) []int {
	arr := []int{}
	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	return arr
}

// Converts an array to a number assuming decimal
// [5 4 3 2 1] -> 54321
func ArrToNum(arr []int) int {
	num := 0
	for _, n := range arr {
		num *= 10
		num += n
	}
	return num
}

// Takes slice representation of number and mods it
func ArrMod(arr []int, mod int) int {
	ret := 0
	for _, n := range arr {
		ret *= 10
		ret += n
		ret %= mod
	}
	return ret
}
