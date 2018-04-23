package dpr

import (
	"fmt"
)

func DeferDemo() {
	fmt.Println("*** Defer Demo ***")

	rule1()

	rule2()

	result := rule3()
	fmt.Println("result=", result)
}

func rule3() (result int) {
	// Deferred functions may read and assign to the returning function's named return values.
	fmt.Println("rule3()")
	result = 10
	defer func() {
		result = 11
	}()
	return result
}

func rule2() {
	// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	fmt.Println("rule2()")
	i := 10
	defer fmt.Println("first defer, i=", i)
	i++
	defer fmt.Println("second defer, i=", i)
	i++
	defer fmt.Println("third defer, i=", i)
}

func rule1() {
	// A deferred function's arguments are evaluated when the defer statement is evaluated.
	fmt.Println("rule1()")
	i := 10
	defer fmt.Println("defer, i=", i)
	i++
	fmt.Println("end, i=", i)
}
