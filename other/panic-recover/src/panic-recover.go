package src

import (
	"errors"
	"fmt"
)

func OuterFunc() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Recovered panic: %s\n", p)
		}
	}()
	innerFunc()
}

func innerFunc() {
	/*	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Recovered panic: %s\n", p)
		}
	}()*/

	panic(errors.New("an intended fatal error"))
}
