package main

import "fmt"

// 竞争条件
func main() {
	var i int
	go func() {
		i++
	}()
	if i == 0 {
		fmt.Printf("i的值为%d\n", i)
	}
}
