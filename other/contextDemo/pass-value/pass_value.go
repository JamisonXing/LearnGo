package main

import (
	"context"
	"fmt"
	"time"
)

// 使用 context 包的 func WithValue() 函数传递数据。
func main() {
	ctx := context.WithValue(context.Background(), "ctxKey1", "ctxVal")
	go func(ctx context.Context) {
		//读取ctx的value
		data, ok := ctx.Value("ctxKey1").(string)
		if ok {
			fmt.Printf("sub goroutine get value from parent goroutine, val=%s\n", data)
		}
	}(ctx)
	time.Sleep(1 * time.Second)
}
