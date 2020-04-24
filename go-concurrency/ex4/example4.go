package main

import (
	"fmt"
	"sync"
	"time"
)

// first_way
var total1 = 0
var lock = sync.Mutex{}

// second_way
var total2 = make(chan int)

// demo - Idiom 1: Don't communicate by sharing memory, share memory by communicating.
func main() {
	// first_way
	commByshareMemory()
	time.Sleep(time.Second)
	fmt.Println(total1)

	// second_way
	shareMemoryByComm()
	total2 <- 0
	time.Sleep(time.Second)
	fmt.Println(<-total2)
}

func shareMemoryByComm() {
	go func() {
		for i := 1; i <= 1000; i++ {
			counter2()
		}
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			counter2()
		}
	}()
}

func commByshareMemory() {
	go func() {
		for i := 1; i <= 1000; i++ {
			counter1()
		}
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			counter1()
		}
	}()
}

func counter1() {
	lock.Lock()
	total1 = total1 + 1
	lock.Unlock()
}

func counter2() {
	total := <-total2
	total = total + 1
	total2 <- total
}
