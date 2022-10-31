package main

import (
	"fmt"
	"sync"
)

// p18，解决的数据竞争，没有解决竞争条件
func main() {
	var memoryAccess sync.Mutex
	var data int
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("i的值为%d\n", data)
	} else {
		fmt.Printf("i的值为%d\n", data)
	}
	memoryAccess.Unlock()
}
