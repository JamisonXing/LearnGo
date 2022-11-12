package chapter02

import (
	"gin-project/controller/chapter05"
	"github.com/gin-gonic/gin"
)

func Router(cp02 *gin.RouterGroup) {
	cp02.GET("/user", User)
	cp02.GET("/user_struct", chapter05.Middleware03(), UserInforStruct) //局部中间件
	cp02.GET("/arr", ArrController)
	cp02.GET("/arrStruct", ArrStruct)

	cp02.GET("/param1/:id", Param1)
	cp02.GET("/param2/*id", Param2)
	cp02.GET("/query", GetQueryData)
	cp02.GET("/query_arr", GetQueryArrData)
	cp02.GET("/query_map", GetQueryMapData)

	cp02.GET("/to_user_add", ToUserAdd)
	cp02.POST("/do_user_add", DoUserAdd)

	cp02.GET("/to_user_add2", ToUserAdd2)
	cp02.POST("/do_user_add2", DoUserAdd2)

	cp02.GET("/to_user_add3", ToUserAdd3)
	cp02.POST("/do_user_add3", DoUserAdd3)

	cp02.GET("/to_user_add4", ToUserAdd4)
	cp02.POST("/do_user_add4", DoUserAdd4)

	cp02.GET("/test_to_upload1", ToUpload1)
	cp02.POST("/test_do_upload1", DoUpload1)

	cp02.GET("/test_to_upload2", ToUpload2)
	cp02.POST("/test_do_upload2", DoUpload2)

	cp02.GET("/test_to_upload3", ToUpload3)
	cp02.POST("/test_do_upload3", DoUpload3)

	cp02.GET("/test_to_upload4", ToUpload4)
	cp02.POST("/test_do_upload4", DoUpload4)

	cp02.GET("/output_json", OutputJson)
	cp02.GET("/output_ascii_json", OutputAsciiJson)
	cp02.GET("/output_jsonp", OutputJsonp)
	cp02.GET("/output_pure_json", OutputPureJson)
	cp02.GET("/output_secure_json", OutputSecureJson)
	cp02.GET("/output_xml", OutputXML)
	cp02.GET("/output_yaml", OutputYAML)
	cp02.GET("/output_proto", OutputProto)

	cp02.GET("/redirect_a", RedirectA)
	cp02.GET("/redirect_b", RedirectB)
}
