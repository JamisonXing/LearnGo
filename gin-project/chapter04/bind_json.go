package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToBindJson(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_json.html", nil)
}

func DoBindJson(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusNotFound,
			"msg":  "bind fail",
		})
		fmt.Println("bind fail: ", err)
	} else {
		fmt.Println(user)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "bind success",
		})
	}
}
