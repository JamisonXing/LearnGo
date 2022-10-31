package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	go func() {
		i++
	}()
	time.Sleep(1 * time.Second) //不优雅
	if i == 0 {
		fmt.Printf("i的值为%d\n", i)
	}
}
