package main

import (
	"fmt"
	"time"
)

//我们可以使用 context 包的 func WithCancel() 函数取消操作，从而避免 goroutine 泄露。

func main() {
	gen := func() <-chan int {
		dst := make(chan int)
		go func() {
			var n int
			for {
				dst <- n
				n++
			}
		}()
		return dst
	}
	for v := range gen() {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
	time.Sleep(1 * time.Second)
}

/*
阅读上面这段代码，我们创建一个 gen() 函数，
启动一个 goroutine 生成整数，循环调用 gen()
函数输出生成的整数，当整数值为 5 时，停止循环，
从输出结果看，没有发现问题。但是，实际上该段代码
会导致 goroutine 泄露，因为 gen() 函数一直在无限循环
*/
