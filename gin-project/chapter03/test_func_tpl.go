package chapter03

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func FuncTpl(ctx *gin.Context) {
	//fmt.Sprintf()是有返回值的,主要用于格式化，通常用于数据写到数据库或者日志里面
	user := fmt.Sprintf("name:%s,age:%d,addr:%s", "jamsion", 20, "xxx")
	fmt.Println(user)
	name := "jamison"
	age := 18
	sliceData := []string{"ZhangSan", "LiSi", "WangWu"}

	mapData := map[string]string{
		"name": "jamison",
		"age":  "34",
	}

	timeNow := time.Now()

	m := map[string]any{
		"name":      name,
		"age":       age,
		"sliceData": sliceData,
		"mapData":   mapData,
		"timeNow":   timeNow,
	}
	ctx.HTML(http.StatusOK, "chapter03/test_func_tpl.html", m)
}

// AddNum 第一步：定义函数
func AddNum(num1, num2 int) int {
	return num1 + num2
}

// SubStr 截取字符串（字符串过长显式不下） 只显示7个字符
func SubStr(str string, n int) string {
	strLen := len(str)
	if n >= strLen {
		return str
	}
	subStr := str[0:n]
	return subStr + "..."
}

func Safe(str string) template.HTML { //将字符串转义为html
	return template.HTML(str)
}
