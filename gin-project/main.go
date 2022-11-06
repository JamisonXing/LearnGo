package main

import (
	"gin-project/chapter01"
	"gin-project/chapter02"
	"gin-project/chapter03"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	//router := gin.New()

	//自定义模板函数 第二步 setMap
	router.SetFuncMap(template.FuncMap{
		"add":    chapter03.AddNum,
		"subStr": chapter03.SubStr,
		"safe":   chapter03.Safe,
	})

	/*	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello gin")
	})*/

	router.LoadHTMLGlob("template/**/*")
	//router.LoadH TMLFiles("index.html")

	//router.Static("/front", "static")
	router.StaticFS("/static", http.Dir("static"))

	router.GET("/", chapter01.Hello)
	router.GET("/user", chapter02.User)
	router.GET("/user_struct", chapter02.UserInforStruct)
	router.GET("/arr", chapter02.ArrController)
	router.GET("/arrStruct", chapter02.ArrStruct)

	router.GET("/param1/:id", chapter02.Param1)
	router.GET("/param2/*id", chapter02.Param2)
	router.GET("/query", chapter02.GetQueryData)
	router.GET("/query_arr", chapter02.GetQueryArrData)
	router.GET("/query_map", chapter02.GetQueryMapData)

	router.GET("/to_user_add", chapter02.ToUserAdd)
	router.POST("/do_user_add", chapter02.DoUserAdd)

	router.GET("/to_user_add2", chapter02.ToUserAdd2)
	router.POST("/do_user_add2", chapter02.DoUserAdd2)

	router.GET("/to_user_add3", chapter02.ToUserAdd3)
	router.POST("/do_user_add3", chapter02.DoUserAdd3)

	router.GET("/to_user_add4", chapter02.ToUserAdd4)
	router.POST("/do_user_add4", chapter02.DoUserAdd4)

	router.GET("/test_to_upload1", chapter02.ToUpload1)
	router.POST("/test_do_upload1", chapter02.DoUpload1)

	router.GET("/test_to_upload2", chapter02.ToUpload2)
	router.POST("/test_do_upload2", chapter02.DoUpload2)

	router.GET("/test_to_upload3", chapter02.ToUpload3)
	router.POST("/test_do_upload3", chapter02.DoUpload3)

	router.GET("/test_to_upload4", chapter02.ToUpload4)
	router.POST("/test_do_upload4", chapter02.DoUpload4)

	router.GET("/output_json", chapter02.OutputJson)
	router.GET("/output_ascii_json", chapter02.OutputAsciiJson)
	router.GET("/output_jsonp", chapter02.OutputJsonp)
	router.GET("/output_pure_json", chapter02.OutputPureJson)
	router.GET("/output_secure_json", chapter02.OutputSecureJson)
	router.GET("/output_xml", chapter02.OutputXML)
	router.GET("/output_yaml", chapter02.OutputYAML)
	router.GET("/output_proto", chapter02.OutputProto)

	router.GET("/redirect_a", chapter02.RedirectA)
	router.GET("/redirect_b", chapter02.RedirectB)

	//模板语法
	router.GET("/test_tpl", chapter03.TestSyntaxTpl)

	//模板函数
	router.GET("/test_func", chapter03.FuncTpl)

	//router.Run(":8080")
	//http.ListenAndServe("localhost:8080", router)

	//自定义http配置
	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := s.ListenAndServe(); err != nil {
		panic("start err")
	}
}
