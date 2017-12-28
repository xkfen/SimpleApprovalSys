package handler

import (
	//"github.com/gin-gonic/gin"
	"gopkg.in/gin-gonic/gin.v1"
	"simpleApproval/model"
	"simpleApproval/simpleApprovalOrderService"
	"gcoresys/common/util"
)

//创建订单
func CreateSimpleApprovalOrderHandler(context *gin.Context){
	var simpleApprovalOrder *model.SimpleApprovalOrder
	if context.BindJSON(&simpleApprovalOrder) == nil && simpleApprovalOrder != nil {
		if err := simpleApprovalOrderService.CreateSimpleApprovalOrder(simpleApprovalOrder);err == nil {
			//util.RenderGinSuccessJson("创建成功", nil ,context)
			util.RenderGinSuccessJson("创建成功", nil, context)
		}else {
			util.RenderGinErrorJson(err.Error(), nil, context)
		}
	}else {
		util.RenderGinErrorJson("参数解析错误", nil, context)
	}
}
