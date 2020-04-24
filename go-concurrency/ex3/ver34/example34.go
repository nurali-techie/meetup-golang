package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Event struct {
	msg  string
	wait chan bool
}

// demo generator, fanIn with maintain sequencing of event using channel. Check, struct Event{wail chan bool}
func main() {
	event1 := restaurantCreated()
	event2 := customerCreated()
	event3 := driverCreated()

	allEvent := fanIn(event1, event2, event3)

	var event Event
	for i := 1; i <= 12; i++ {
		event = <-allEvent
		fmt.Println(event.msg)
		event.wait <- true
	}

}

// fanIn pattern / aggregator pattern
func fanIn(event1, event2, event3 <-chan Event) <-chan Event {
	ch := make(chan Event)

	go func() {
		for {
			ch <- <-event1
		}
	}()
	go func() {
		for {
			ch <- <-event2
		}
	}()
	go func() {
		for {
			ch <- <-event3
		}
	}()

	return ch
}

// generator pattern / producer pattern
func restaurantCreated() <-chan Event {
	ch := make(chan Event)
	go func() {
		for i := 1; i <= 4; i++ {
			e := Event{msg: fmt.Sprintf("restaurnat-%d", i), wait: make(chan bool)}
			ch <- e
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
			<-e.wait
		}
	}()
	return ch
}

func customerCreated() <-chan Event {
	ch := make(chan Event)
	go func() {
		for i := 1; i <= 4; i++ {
			e := Event{msg: fmt.Sprintf("customer-%d", i), wait: make(chan bool)}
			ch <- e
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
			<-e.wait
		}
	}()
	return ch
}

func driverCreated() <-chan Event {
	ch := make(chan Event)
	go func() {
		for i := 1; i <= 4; i++ {
			e := Event{msg: fmt.Sprintf("driver-%d", i), wait: make(chan bool)}
			ch <- e
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
			<-e.wait
		}
	}()
	return ch
}
