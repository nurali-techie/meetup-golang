package main

import (
	"fmt"
	"math/rand"
	"time"
)

// demo generator, fanIn pattern with select statement
func main() {
	event1 := restaurantCreated()
	event2 := customerCreated()
	event3 := driverCreated()

	allEvent := fanIn(event1, event2, event3)

	var event string
	for i := 1; i <= 12; i++ {
		event = <-allEvent
		fmt.Println(event)
	}

}

// fanIn pattern / aggregator pattern
func fanIn(event1, event2, event3 <-chan string) <-chan string {
	ch := make(chan string)

	// go func() { for { ch <- <-event1}} ()
	// go func() { for { ch <- <-event2}} ()
	// go func() { for { ch <- <-event3}} ()

	go func() {
		for {
			select {
			case s := <-event1:
				ch <- s
			case s := <-event2:
				ch <- s
			case s := <-event3:
				ch <- s
			}
		}
	}()

	return ch
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

func customerCreated() <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= 4; i++ {
			ch <- fmt.Sprintf("customer-%d", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return ch
}

func driverCreated() <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= 4; i++ {
			ch <- fmt.Sprintf("driver-%d", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return ch
}
