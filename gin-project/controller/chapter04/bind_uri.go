package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindUri(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.String(http.StatusNotFound, "bind fail")
	}
	fmt.Println(user)
	ctx.String(http.StatusOK, "bind success")

}
