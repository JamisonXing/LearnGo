package main

import (
	"fmt"
	"math/rand"
)

// GenerateIntA 最简单的带缓存的生成器
func GenerateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		//生成随机数，传入通道
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func main() {
	ch := GenerateIntA()
	/*	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}*/
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
