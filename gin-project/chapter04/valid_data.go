package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Id      int    `form:"-"`
	Title   string `form:"title" binding:"required"`
	Content string `form:"content"`
	Desc    string `form:"desc"`
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
