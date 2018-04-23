package err

import (
	"errors"
	"fmt"
)

func SimpleErrorDemo() {
	fmt.Println("*** Simple Error Demo ***")

	err := simple()
	if err != nil {
		fmt.Println(err) // println will call err.Error() method rather err.String()
	}
}

func simple() error {
	err := errors.New("first error")
	// err := fmt.Errorf("second error, with something %d", 100)
	return err
}

// error interface defined by go
// type error interface {
// 	Error() string
// }
