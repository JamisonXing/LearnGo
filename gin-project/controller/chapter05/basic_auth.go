package chapter05

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 私有数据
var mapData map[string]any = map[string]any{
	"zs": gin.H{"age": 18, "addr": "zs--xx"},
	"ls": gin.H{"age": 19, "addr": "ls--xx"},
	"ww": gin.H{"age": 20, "addr": "ww--xx"},
}

func AuthTest(ctx *gin.Context) {
	userName := ctx.Query("user_name")
	if data, ok := mapData[userName]; ok {
		ctx.JSON(http.StatusOK, gin.H{"user": userName, "data": data})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"user": userName, "data": "no data"})
	}
}

func WrapFTest(w http.ResponseWriter, r *http.Request) {

}
