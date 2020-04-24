package main

import (
	"fmt"
	"runtime"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

// demo Daisy chain
func main() {
	prn("g1")
	const n = 100000 // 100k
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	prn("g2")
	go func(c chan int) { c <- 1 }(right)
	prn("g3")
	fmt.Println(<-leftmost)
	prn("g4")
}

func prn(name string) {
	fmt.Println(name, runtime.NumGoroutine())
}
