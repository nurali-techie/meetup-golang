package dpr

import (
	"fmt"
	"time"
)

func PanicDemo() {
	fmt.Println("** Panic Demo ***")
	fmt.Println("THERE ARE TWO DEMO, demo1(), demo2(), ENABLE ONE BY ONE")

	// demo1()
	// demo2()
}

func demo1() {
	fmt.Println("demo1()")
	first()
}

func demo2() {
	fmt.Println("demo2()")
	go first()
	time.Sleep(time.Second * 100)
	fmt.Println("after wait")
}

func first() {
	defer fmt.Println("plz call me")
	second()
}

func second() {
	panic(10)
}
