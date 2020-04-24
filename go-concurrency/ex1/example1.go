package main

import (
	"fmt"
)

// demo - goroutine and channel
func main() {
	ch := make(chan int)

	fmt.Print("Hello")
	go greet(ch)
	<-ch
}

func greet(ch chan int) {
	fmt.Print("World")
	ch <- 0
}
