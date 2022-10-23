package main

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	//ctx.String(200, "hello gin")
	name := "xxj"
	ctx.HTML(200, "index/index.html", name)
}

func User(ctx *gin.Context) {
	ctx.HTML(200, "user/user.html", nil)
}

func main() {
	router := gin.Default()
	//router := gin.New()

	/*	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello gin")
	})*/

	router.LoadHTMLGlob("template/**/*")
	//router.LoadHTMLFiles("index.html")

	router.GET("/", Hello)
	router.GET("/user", User)

	router.Run(":9000")
}
