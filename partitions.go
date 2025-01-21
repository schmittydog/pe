package main

func generatePartitions(n, index int, ch chan []int, arr []int) {
	if n == 0 {
		ch <- CopyArr(arr, index-1)
		if arr[0] == 1 && arr[len(arr)-1] == 1 {
			close(ch)
		}
		return
	}
	if n < 0 {
		return
	}
	start := n
	if index > 0 {
		start = arr[index-1]
	}
	for nxt := start; nxt > 0; nxt-- {
		arr[index] = nxt
		generatePartitions(n-nxt, index+1, ch, arr)
	}
}

func Partitions(n int) chan []int {
	arr := make([]int, n)
	ch := make(chan []int)
	go generatePartitions(n, 0, ch, arr)
	return ch
}

func generateAscDec(n, index int, ch chan []int, arr []int) {
	if index == n {
		ch <- CopyArr(arr, index-1)
		if arr[0] == 9 && arr[len(arr)-1] == 9 {
			close(ch)
		}
		return
	}
	start := 1
	if index > 0 {
		start = arr[index-1]
	}
	for nxt := start; nxt < 10; nxt++ {
		arr[index] = nxt
		generateAscDec(n, index+1, ch, arr)
	}
}

func AscendingDecimal(n int) chan []int {
	ch := make(chan []int)
	arr := make([]int, n)
	go generateAscDec(n, 0, ch, arr)
	return ch
}
