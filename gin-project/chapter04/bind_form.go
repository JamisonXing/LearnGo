package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
	Addr string `form:"addr" json:"addr"`
}

func ToBindForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_form.html", nil)
}

func DoBindForm(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.String(http.StatusNotFound, "bind fail")
	} else {
		fmt.Println(user)
		ctx.String(http.StatusOK, "bind success")
	}
}
