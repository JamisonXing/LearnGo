package chapter05

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Middleware01 第一种定义方式
func Middleware01(ctx *gin.Context) {
	fmt.Println("custom middleware01")
}

// Middleware02 第二种定义方式
func Middleware02() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("custom middleware02")
	}
}
