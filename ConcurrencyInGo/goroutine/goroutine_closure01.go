package main

import (
	"fmt"
	"sync"
)

// 正确方法是将salutation的副本传递到闭包中
func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

/*
result:
good day
hello
greetings
*/

/*
并没有像我们想的那样顺序打印
*/
