package chapter02

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OutputJson(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "submit success",
		"html": "<b>Hello, world!</b>",
	})
}

func OutputAsciiJson(ctx *gin.Context) {
	ctx.AsciiJSON(http.StatusOK, gin.H{ //生成转义的非ASCII字符的ASCII-only JSON
		"code": 200,
		"tag":  "<br>",
		"msg":  "submit success",
		"html": "<b>Hello, world!</b>",
	})
}

func OutputJsonp(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, gin.H{ //使用JSONP向不同域的服务器请求数据，如果查询参数存在问题，则将回调添加到响应体中
		"code": 200,
		"tag":  "<br>",
		"msg":  "submit success",
		"html": "<b>Hello, world!</b>",
	})
}

func OutputPureJson(ctx *gin.Context) {
	ctx.PureJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "submit success",
		"html": "<b>Hello, world!</b>",
	})
}

func OutputSecureJson(ctx *gin.Context) {
	ctx.SecureJSON(http.StatusOK, gin.H{ //防止json劫持
		"code": 200,
		"tag":  "<br>",
		"msg":  "submit success",
		"html": "<b>Hello, world!</b>",
	})
}

func OutputXML(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "submit success",
		"html": "<b>Hello, world!</b>",
	})
}

func OutputYAML(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{ //
		"code": 200,
		"tag":  "<br>",
		"msg":  "submit success",
		"html": "<b>Hello, world!</b>",
	})
}
