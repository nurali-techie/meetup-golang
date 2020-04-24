package main

import (
	"fmt"
	"math/rand"
	"time"
)

// demo generator pattern
func main() {
	event1 := restaurantCreated()

	var event string
	for i := 1; i <= 4; i++ {
		event = <-event1
		fmt.Println(event)
	}

}

// generator pattern / producer pattern
func restaurantCreated() <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= 4; i++ {
			ch <- fmt.Sprintf("restaurnat-%d", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return ch
}
