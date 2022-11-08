package chapter01

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	//ctx.String(200, "hello gin")
	name := "xxj"
	ctx.HTML(200, "index/index.html", name)
}
