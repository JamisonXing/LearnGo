package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 9 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Second)
}
