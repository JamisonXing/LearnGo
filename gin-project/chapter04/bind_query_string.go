package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindQueryString(ctx *gin.Context) {
	var user User
	ctx.ShouldBind(&user)
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.String(http.StatusNotFound, "bind fail")
	} else {
		fmt.Println(user)
		ctx.String(http.StatusOK, "bind success")
	}
}
