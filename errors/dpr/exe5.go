package dpr

import (
	"fmt"
)

func RecoverDemo() {
	fmt.Println("*** Recover Demo ***")
	f()
}

func f() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered, r=", r)
			// panic(r)
		}
	}()
	g(10)
	// g(1)
}

func g(no int) {
	if no > 2 {
		panic(no)
	}
	fmt.Println("from g, no=", no)
}
