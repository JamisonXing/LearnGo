package chapter01

import "github.com/gin-gonic/gin"

func Router(cp01 *gin.RouterGroup) {
	cp01.GET("/", Hello)
}
