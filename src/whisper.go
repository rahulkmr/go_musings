package main

import "fmt"

func whisper(left, right chan int) {
	right <- 1 + <-left
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	left := leftmost
	right := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go whisper(left, right)
		left = right
	}
	leftmost <- 1
	fmt.Println(<-right)
}
