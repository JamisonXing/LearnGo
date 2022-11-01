package main

import (
	"fmt"
	"sync"
	"time"
)

//p50
//创建join点（go的并发模型fork-join），而不是使用sleep增加goroutine执行的概率

func main() {
	var wg sync.WaitGroup
	sayhello := func() {
		defer wg.Done()
		fmt.Print("hello")
	}
	wg.Add(1)
	go sayhello()
	wg.Wait()

	//而不是下面这种写法
	sayhello01 := func() {
		defer wg.Done()
		fmt.Print("hello01")
	}
	go sayhello01()
	time.Sleep(1 * time.Second) //仅仅增加了概率
}
