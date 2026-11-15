package main

func concurrentFib(n int) []int {
	chanInt := make(chan int)
	sliceInt := make([]int, 0)
	go fibonacci(n, chanInt)
	for fib := range chanInt {
		sliceInt = append(sliceInt, fib)
	}
	return sliceInt
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
