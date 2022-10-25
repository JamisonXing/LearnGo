package main

import (
	"gin-project/chapter01"
	"gin-project/chapter02"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//router := gin.New()

	/*	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello gin")
	})*/

	router.LoadHTMLGlob("template/**/*")
	//router.LoadH TMLFiles("index.html")

	//router.Static("/front", "static")
	router.StaticFS("/static", http.Dir("static"))

	router.GET("/", chapter01.Hello)
	router.GET("/user", chapter02.User)
	router.GET("/user_struct", chapter02.UserInforStruct)
	router.GET("/arr", chapter02.ArrController)
	router.GET("/arrStruct", chapter02.ArrStruct)

	router.GET("/param1/:id", chapter02.Param1)
	router.GET("/param2/*id", chapter02.Param2)
	router.GET("/query", chapter02.GetQueryData)
	router.GET("/query_arr", chapter02.GetQueryArrData)
	router.GET("/query_map", chapter02.GetQueryMapData)

	router.GET("/to_user_add", chapter02.ToUserAdd)
	router.POST("/do_user_add", chapter02.DoUserAdd)

	router.Run(":8080")
}
