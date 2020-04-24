package main

import (
	"fmt"
	"math/rand"
	"time"
)

// demo generator, fanIn and fanOut patterns
func main() {
	notify1, notify2 := fanOut()

	go func() {
		for i := 1; i <= 6; i++ {
			fmt.Println(<-notify1)
		}
	}()
	go func() {
		for i := 1; i <= 6; i++ {
			fmt.Println(<-notify2)
		}
	}()

	time.Sleep(4 * time.Second)
}

// fanOut pattern
func fanOut() (notify1, notify2 chan string) {
	notify1 = make(chan string)
	notify2 = make(chan string)

	event1 := restaurantCreated()
	event2 := customerCreated()
	event3 := driverCreated()

	allEvent := fanIn(event1, event2, event3)

	go func() {
		var event string
		for i := 1; i <= 12; i++ {
			event = <-allEvent
			if i%2 == 0 {
				email := fmt.Sprintf("android-%s", event)
				notify1 <- email
			} else {
				email := fmt.Sprintf("apple-%s", event)
				notify2 <- email
			}
		}
	}()

	return
}

// fanIn pattern / aggregator pattern
func fanIn(event1, event2, event3 <-chan string) <-chan string {
	ch := make(chan string)

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
