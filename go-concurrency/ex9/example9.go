package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	fmt.Println(runtime.NumCPU())
}
