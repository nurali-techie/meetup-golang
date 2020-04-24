package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player("ping", table)
	go player("pong", table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(name string, table chan int) {
	for {
		ball := <-table
		ball++
		fmt.Println(name)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
