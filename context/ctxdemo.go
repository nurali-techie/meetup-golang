package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Context Demo")

	test1()
	test2()
	test3()
	test4()
}

func test1() {
	log.Println("test1: waitAndWatch() prints 'over' after 5 seconds")
	ctx := context.Background()
	waitAndWatch(ctx)
}

func test2() {
	log.Println("test2: waitAndWatch() prints 'ctx.Done()' after 2 seconds")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	waitAndWatch(ctx)
}

func test3() {
	log.Println("test3: waitAndWatch() prints 'ctx.Done()' if 'enter' pressed before 5 seconds otherwise prints 'over' after 5 seconds")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	waitAndWatch(ctx)
}

func test4() {
	log.Println("test4: waitAndWatch() prints 'ctx.Done()' if 'enter' pressed before 3 seconds (due to cancel_context) otherwise prints 'ctx.Done()' after 3 seconds (due to timeout_context)")
	ctx := context.Background()
	ctx, cancel1 := context.WithTimeout(ctx, 3*time.Second)
	defer cancel1()
	ctx, cancel2 := context.WithCancel(ctx)
	defer cancel2()

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel1()
	}()

	waitAndWatch(ctx)
}

func waitAndWatch(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("ctx.Done()")
	case <-time.After(5 * time.Second):
		log.Println("over")
	}
}
