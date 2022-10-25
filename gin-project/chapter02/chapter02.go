package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}

func User(ctx *gin.Context) {
	name := "xxj"
	ctx.HTML(200, "user/user.html", name)
}

// UserInforStruct 结构体渲染
func UserInforStruct(ctx *gin.Context) {
	userInfo := UserInfo{
		Id:   1,
		Name: "xxj",
		Age:  20,
		Addr: "AUST",
	}
	userInfo1 := UserInfo{
		Id:   1,
		Name: "xxj",
		Age:  20,
		Addr: "AUST",
	}

	userMap := make(map[int]UserInfo)
	userMap[0], userMap[1] = userInfo, userInfo1
	ctx.HTML(http.StatusOK, "chapter02/user_info.html", userMap)
}

// ArrController 数组渲染
func ArrController(ctx *gin.Context) {
	arr := [3]int{1, 2}
	ctx.HTML(http.StatusOK, "chapter02/arr.html", arr)
}

// ArrStruct 结构体数组
func ArrStruct(ctx *gin.Context) {
	arrStruct := [3]UserInfo{{Id: 1, Name: "xxj", Age: 12, Addr: "xxx"},
		{Id: 2, Name: "xxj1", Age: 13, Addr: "xxx1"},
		{Id: 3, Name: "xxj2", Age: 14, Addr: "xxx2"}}
	ctx.HTML(http.StatusOK, "chapter02/arr-struct.html", arrStruct)
}

// Param1 在路径上直接加上参数值
func Param1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello, %s", id)
}
func Param2(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello, %s", id)
}

func GetQueryData(ctx *gin.Context) {
	id := ctx.Query("id")

	name := ctx.DefaultQuery("name", "defaultVal")
	ctx.String(http.StatusOK, "hello, %s,%s", id, name)
}

func GetQueryArrData(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	ctx.String(http.StatusOK, "xxj, %s", ids)
}

func GetQueryMapData(ctx *gin.Context) {
	user := ctx.QueryMap("user")
	ctx.String(http.StatusOK, "xxj, %s", user)
}

func ToUserAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/user_add.html", nil)
}
func DoUserAdd(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	favoriteArr := ctx.PostFormArray("favorite")
	userMap := ctx.PostFormMap("user")
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(favoriteArr)
	fmt.Println(userMap)
	ctx.String(http.StatusOK, "添加成功")
}
