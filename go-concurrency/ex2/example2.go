package main

import (
	"fmt"
	"math/rand"
	"time"
)

// demo select statement
func main() {
	rand.Intn(1e3)

	ch1 := make(chan string)
	go greet1(ch1)

	ch2 := make(chan string)
	go greet2(ch2)

	time.Sleep(time.Millisecond * 500)
	for i := 1; i <= 4; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("from ch1,", msg)
		case msg := <-ch2:
			fmt.Println("from ch2,", msg)
		default:
			fmt.Println("from default")
		}
	}
}

func greet1(ch chan string) {
	for i := 1; i <= 2; i++ {
		ch <- fmt.Sprintf("%s-%d", "Hello World", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	}
}

func greet2(ch chan string) {
	for i := 1; i <= 2; i++ {
		ch <- fmt.Sprintf("%s-%d", "Hello Go", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	}
}
