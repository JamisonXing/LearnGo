package main

import (
	"fmt"
	"sync"
)

// p51
func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

/*
result:
good day
good day
good day
*/

/*
并没有像我们想的那样打印三个不同的值
*/
