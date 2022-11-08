package chapter03

import "github.com/gin-gonic/gin"

func Router(cp3 *gin.RouterGroup) {
	cp3.GET("/test_tpl", TestSyntaxTpl)
}
