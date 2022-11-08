package chapter03

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Id    int
	Title string
	Desc  string
}

func TestSyntaxTpl(ctx *gin.Context) {
	name := "jamison"
	arr := []int{1, 2, 3, 4}

	article := Article{
		Id:    0,
		Title: "go",
		Desc:  "Concurrency In Go",
	}
	m := map[any]any{
		"name":     name,
		"arr":      arr,
		"birthday": "2002-04-20",
		"article":  article,
	}
	ctx.HTML(http.StatusOK, "chapter03/test.html", m)
}
