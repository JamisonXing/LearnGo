package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ToUpload1(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/test_upload1.html", nil)
}
func DoUpload1(ctx *gin.Context) { //单文件上传
	file, _ := ctx.FormFile("file")
	fmt.Println(file.Filename)

	//防止文件重复,加上那个时间戳
	time_unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int, 10)

	//保存
	dst := "upload/" + time_unix_str + "-" + file.Filename
	ctx.SaveUploadedFile(file, dst)

	ctx.String(http.StatusOK, "上传成功")
}

func ToUpload2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/test_upload2.html", nil)
}
func DoUpload2(ctx *gin.Context) { //多文件上传
	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		//防止文件重复,加上那个时间戳
		time_unix_int := time.Now().Unix()
		time_unix_str := strconv.FormatInt(time_unix_int, 10)
		//保存
		dst := "upload/" + time_unix_str + "-" + file.Filename
		ctx.SaveUploadedFile(file, dst)
		ctx.String(http.StatusOK, file.Filename+"上传成功"+"\n")
	}

}
