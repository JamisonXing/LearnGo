package allRouter

import (
	"gin-project/controller/chapter01"
	"gin-project/controller/chapter02"
	"gin-project/controller/chapter03"
	"gin-project/controller/chapter04"
	"gin-project/controller/chapter05"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	cp01 := router.Group("/chapter01")
	cp01.Use(chapter05.Middleware01) // 路由组中间件
	cp02 := router.Group("/chapter02")
	cp03 := router.Group("/chapter03")
	cp04 := router.Group("/chapter04")
	cp05 := router.Group("/chapter05")

	chapter01.Router(cp01)
	chapter02.Router(cp02)
	chapter03.Router(cp03)
	chapter04.Router(cp04)
	chapter04.Router(cp05)
}
