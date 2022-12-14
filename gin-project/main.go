package main

import (
	"gin-project/controller/chapter03"
	"gin-project/controller/chapter04"
	allRouter "gin-project/router"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"html/template"
	"net/http"
	"time"
)

func main() {
	router := gin.Default() //有中间件Logger(), Recovery()
	//router := gin.New()      //没有中间件
	//router.Use(gin.Logger()) //使用中间件

	//全局中间件
	/*	router.Use(chapter05.Middleware01)   //方式一
		router.Use(chapter05.Middleware02()) //方式二
		router.Use(chapter05.Middleware03()) //方式二*/
	//router.Use(chapter05.Middleware02())

	//自定义模板函数 第二步 setMap
	router.SetFuncMap(template.FuncMap{
		"add":    chapter03.AddNum,
		"subStr": chapter03.SubStr,
		"safe":   chapter03.Safe,
	})

	/*	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello gin")
	})*/

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("lenValid", chapter04.LenValid)
	}

	router.LoadHTMLGlob("template/**/*")
	//router.LoadH TMLFiles("index.html")

	//router.Static("/front", "static")
	router.StaticFS("/static", http.Dir("static"))
	/*
		//路由太多了，使用路由组
		v1 := router.Group("/v1")
		{
			v1.GET("/", chapter01.Hello)
		}
		v2 := router.Group("/v2")
		{
			v2.GET("/user_struct", chapter02.UserInforStruct)
			v2.GET("/arr", chapter02.ArrController)
		}*/

	//路由分组
	allRouter.Router(router)

	//在路由之后使用中间件，要注意中间件的执行顺序
	//router.Use(gin.Recovery())

	//router.GET("/", chapter01.Hello)
	//router.GET("/user", chapter02.User)
	//router.GET("/user_struct", chapter02.UserInforStruct)
	//router.GET("/arr", chapter02.ArrController)
	//router.GET("/arrStruct", chapter02.ArrStruct)
	//
	//router.GET("/param1/:id", chapter02.Param1)
	//router.GET("/param2/*id", chapter02.Param2)
	//router.GET("/query", chapter02.GetQueryData)
	//router.GET("/query_arr", chapter02.GetQueryArrData)
	//router.GET("/query_map", chapter02.GetQueryMapData)
	//
	//router.GET("/to_user_add", chapter02.ToUserAdd)
	//router.POST("/do_user_add", chapter02.DoUserAdd)
	//
	//router.GET("/to_user_add2", chapter02.ToUserAdd2)
	//router.POST("/do_user_add2", chapter02.DoUserAdd2)
	//
	//router.GET("/to_user_add3", chapter02.ToUserAdd3)
	//router.POST("/do_user_add3", chapter02.DoUserAdd3)
	//
	//router.GET("/to_user_add4", chapter02.ToUserAdd4)
	//router.POST("/do_user_add4", chapter02.DoUserAdd4)
	//
	//router.GET("/test_to_upload1", chapter02.ToUpload1)
	//router.POST("/test_do_upload1", chapter02.DoUpload1)
	//
	//router.GET("/test_to_upload2", chapter02.ToUpload2)
	//router.POST("/test_do_upload2", chapter02.DoUpload2)
	//
	//router.GET("/test_to_upload3", chapter02.ToUpload3)
	//router.POST("/test_do_upload3", chapter02.DoUpload3)
	//
	//router.GET("/test_to_upload4", chapter02.ToUpload4)
	//router.POST("/test_do_upload4", chapter02.DoUpload4)
	//
	//router.GET("/output_json", chapter02.OutputJson)
	//router.GET("/output_ascii_json", chapter02.OutputAsciiJson)
	//router.GET("/output_jsonp", chapter02.OutputJsonp)
	//router.GET("/output_pure_json", chapter02.OutputPureJson)
	//router.GET("/output_secure_json", chapter02.OutputSecureJson)
	//router.GET("/output_xml", chapter02.OutputXML)
	//router.GET("/output_yaml", chapter02.OutputYAML)
	//router.GET("/output_proto", chapter02.OutputProto)
	//
	//router.GET("/redirect_a", chapter02.RedirectA)
	//router.GET("/redirect_b", chapter02.RedirectB)

	//模板语法
	//router.GET("/test_tpl", chapter03.TestSyntaxTpl)

	//模板函数
	//router.GET("/test_func", chapter03.FuncTpl)

	/*//数据绑定
	router.GET("/to_bind_form", chapter04.ToBindForm)
	router.POST("/do_bind_form", chapter04.DoBindForm)
	router.GET("/bind_query_string", chapter04.ToBindJson)
	router.GET("/to_bind_json", chapter04.ToBindJson)
	router.POST("/do_bind_json", chapter04.DoBindJson)
	router.GET("/bind_uri/:name/:age/:addr", chapter04.BindUri)

	//数据校验
	router.GET("/to_valid", chapter04.ToValidData)
	router.POST("/do_valid", chapter04.DoValidData)

	*/
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
