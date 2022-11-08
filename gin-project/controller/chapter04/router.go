package chapter04

import "github.com/gin-gonic/gin"

func Router(cp04 *gin.RouterGroup) {
	//数据绑定
	cp04.GET("/to_bind_form", ToBindForm)
	cp04.POST("/do_bind_form", DoBindForm)
	cp04.GET("/bind_query_string", ToBindJson)
	cp04.GET("/to_bind_json", ToBindJson)
	cp04.POST("/do_bind_json", DoBindJson)
	cp04.GET("/bind_uri/:name/:age/:addr", BindUri)

	//数据校验
	cp04.GET("/to_valid", ToValidData)
	cp04.POST("/do_valid", DoValidData)

}
