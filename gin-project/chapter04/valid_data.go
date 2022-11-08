package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"net/http"
)

type Article struct {
	Id      int     `form:"-"`
	Title   string  `form:"title" binding:"lenValid"`
	Content string  `form:"content" binding:"min=5"`
	Desc    string  `form:"desc"`
	Name    [][]int `binding:"dive,max=20,dive,required"` //嵌套验证
}

// LenValid 判断是不是大于6
var LenValid validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field()

	fmt.Println(data)

	return true
}

func ToValidData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/valid_data.html", nil)
}

func DoValidData(ctx *gin.Context) {
	var article Article
	if err := ctx.ShouldBind(&article); err != nil {
		fmt.Println(err)
		ctx.String(http.StatusNotFound, "fail", nil)
	}
	fmt.Println(article)
	ctx.String(http.StatusOK, "validation success", nil)
}
