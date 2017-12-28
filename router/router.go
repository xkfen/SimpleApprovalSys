package router

import (
	"simpleApproval/handler"
	"gopkg.in/gin-gonic/gin.v1"
)

func StartHttpRouter() *gin.Engine{

	//获得路由实例
	r := gin.Default()
	//正式用
	g := r.Group("/api/v1")
	// 路由配置
	{
		// 登录
		g.POST("/createSimpleApprovalOrder", handler.CreateSimpleApprovalOrderHandler)
	}
	return r
}