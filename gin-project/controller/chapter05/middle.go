package chapter05

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Middleware01 第一种定义方式
func Middleware01(ctx *gin.Context) {
	fmt.Println("custom middleware01--start")

	ctx.Next() //执行下一个中间件
	fmt.Println("custom middleware01--end")

}

// Middleware02 第二种定义方式
func Middleware02() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("custom middleware02--start")

		if true { //假设不满足某个条件
			ctx.Abort() //终止调用整个链条
		}
		ctx.Next()

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
