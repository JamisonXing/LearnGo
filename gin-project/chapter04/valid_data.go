package chapter04

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"net/http"
)

//自带验证器和自定义验证器
/*type Article struct {
	Id      int     `form:"-"`
	Title   string  `form:"title" binding:"lenValid"`
	Content string  `form:"content" binding:"min=5"`
	Desc    string  `form:"desc"`
	Name    [][]int `binding:"dive,max=20,dive,required"` //嵌套验证
}*/

type Article struct { // valid for beego
	Id      int    `form:"-" valid:"Min(0)"`
	Title   string `form:"title" valid:"Required"`
	Content string `form:"content"`
	Desc    string `form:"desc"`
	Email   string `form:"email" valid:"Email"`
}

// LenValid 自定义验证器
var LenValid validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field().Interface().(string)

	fmt.Println(data)
	if len(data) > 6 {
		return true
	}
	return true
}

func ToValidData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/valid_data.html", nil)
}

//binding
/*func DoValidData(ctx *gin.Context) {
	var article Article
	if err := ctx.ShouldBind(&article); err != nil {
		fmt.Println(err)
		ctx.String(http.StatusNotFound, "fail", nil)
	}
	fmt.Println(article)

	ctx.String(http.StatusOK, "validation success", nil)
}
*/

// beego 对结构体数据进行校验
func DoValidData(ctx *gin.Context) {
	var article Article
	if err := ctx.ShouldBind(&article); err != nil {
		fmt.Println(err)
		ctx.String(http.StatusNotFound, "fail", nil)
	}
	fmt.Println(article)

	//自定义errMessage
	var MessageTmpls = map[string]string{
		"Required": "不能为空",
		"Min":      "最小值:%d",
		"Max":      "最大值:%d",
		"Range":    "取值范围%d to %d",
		"Email":    "邮箱格式不正确",
	}
	validation.SetDefaultMessage(MessageTmpls)

	keyMapping := map[string]interface{}{
		"Title.Required.": "标题",
		"Content.Min.":    "内容",
		"Email.Email.":    "邮箱",
	}

	//初始化验证器
	valid := validation.Validation{}
	if b, err := valid.Valid(&article); err != nil {
		fmt.Println("fail:", err)
	} else if !b {
		for _, err2 := range valid.Errors {
			fmt.Println(err2.Key) //哪个字段发生错误
			key := keyMapping[err2.Key]
			fmt.Println(err2.Message) //错误信息
			ctx.String(http.StatusOK, key.(string)+err2.Message)
		}

	}
}
