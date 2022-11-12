package chapter05

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Middleware01 第一种定义方式
func Middleware01(ctx *gin.Context) {
	start_time := time.Now()

	fmt.Println("custom middleware01--start")

	ctx.Next() //执行下一个中间件
	fmt.Println("custom middleware01--end")

	cost := time.Since(start_time) //计算请求耗时
	fmt.Println("request time is:", cost)
}

// Middleware02 第二种定义方式
func Middleware02() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("custom middleware02--start")

		/*		if true { //假设不满足某个条件
				ctx.Abort() //终止调用整个链条
			}*/
		ctx.Next()
		time.Sleep(time.Second * 2)
		fmt.Println("custom middleware02--end")
	}
}

func Middleware03() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("custom middleware03--start")
		ctx.Next()
		fmt.Println("custom middleware03--end")
	}
}
