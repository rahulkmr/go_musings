package main

import "fmt"

func setup_channels(left chan int, num int) chan int {
	for i := 0; i < num; i++ {
		right := make(chan int)
		go func(left, right chan int) { right <- 1 + <-left }(left, right)
		left = right
	}
	return left
}

func main() {
	leftmost := make(chan int)
	rightmost := setup_channels(leftmost, 100000)
	leftmost <- 1
	fmt.Println(<-rightmost)
}
