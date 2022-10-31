package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//我们可以使用 context 包的 func WithTimeout() 函数设置超时时间，从而避免请求阻塞。
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
