package chapter05

import "github.com/gin-gonic/gin"

func Router(cp05 *gin.RouterGroup) {
	cp05.GET("/auth_test", gin.BasicAuth(gin.Accounts{
		"zs": "123456",
		"ls": "12345",
		"ww": "1234",
	}), gin.WrapF(WrapFTest), AuthTest)

}
