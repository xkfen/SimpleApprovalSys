package router

import (
	"simpleApproval/handler"
	"gopkg.in/gin-gonic/gin.v1"
)

func GetHttpRouter() *gin.Engine{

	//获得路由实例
	r := gin.Default()
	//正式用
	g := r.Group("/api/v1")
	// 路由配置
	{
		// 创建订单
		g.POST("/createSimpleApprovalOrder", handler.CreateSimpleApprovalOrderHandler)
		// 修改订单
		g.POST("/updateSimpleApprovalOrder", handler.UpdateSimpleApprovalOrder)

		// 根据进件id查询订单信息
		g.POST("/querySimpleApprovalOrderById", handler.QuerySimpleApprovalOrderById)

	//	g.POST("/getSimpleApprovalOrderByJinJianId",hander.GetSimpleApprovalOrderByJinJianId)
	}
	return r
}