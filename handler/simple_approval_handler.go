package handler

import (
	"gopkg.in/gin-gonic/gin.v1"
	"simpleApproval/model"
	"simpleApproval/simpleApprovalOrderService"
	"gcoresys/common/util"
)

// 创建订单
func CreateSimpleApprovalOrderHandler(context *gin.Context){
	var simpleApprovalOrder *model.SimpleApprovalOrder
	if context.BindJSON(&simpleApprovalOrder) == nil && simpleApprovalOrder != nil {
		if err := simpleApprovalOrderService.CreateSimpleApprovalOrder(simpleApprovalOrder);err == nil {
			util.RenderGinSuccessJson("创建成功", nil, context)
		}else {
			util.RenderGinErrorJson(err.Error(), nil, context)
		}
	}else {
		util.RenderGinErrorJson("参数解析错误", nil, context)
	}
}


func QuerySimpleApprovalOrderById(context *gin.Context){
	order := &model.SimpleApprovalOrder{}
	order.JinJianId = context.PostForm("jin_jian_id")
	if context.BindJSON(&order) == nil && order != nil {
		if err := simpleApprovalOrderService.QuerySimpleApprovalOrderById(order.JinJianId); err == nil{
			util.RenderGinSuccessJson("查找成功", nil, context)
		}else {
			util.RenderGinErrorJson(err.Error(), nil, context)
		}
	}else {
		util.RenderGinErrorJson("参数解析错误", nil, context)
	}
}

//// 根据进件id查询订单信息
//func QuerySimpleApprovalOrderById(context*gin.Context){
//	//var simpleApprovalOrder *model.SimpleApprovalOrder
//	id ,err := strconv.Atoi()
//	jinJianId := context.Params.ByName("jin_jian_id")
//	if context.BindJSON(jin_jian_id) == nil && jin_jian_id != "" {
//		err, interview := service.GetInterview(&model.Interview{OrderId: orderId})
//		if err := simpleApprovalOrderService.QuerySimpleApprovalOrderById(&model.SimpleApprovalOrder{JinJianId: jin_jian_id});err == nil {
//			util.RenderGinSuccessJson("查询成功", nil, context)
//		}else {
//			//util.RenderGinErrorJson(err, nil, context)
//			util.RenderGinErrorJson("error",nil, context)
//		}
//	}else {
//		util.RenderGinErrorJson("参数解析错误", nil, context)
//	}
//}

func GetSimpleApprovalOrderByJinJianId(context *gin.Context){
	var simpleApprovalOrder *model.SimpleApprovalOrder
	//err, info := simpleApprovalOrderService.GetSimpleApprovalOrder(&model.SimpleApprovalOrder{JinJianId:simpleApprovalOrder.JinJianId,
	//					JinJianUserName:simpleApprovalOrder.JinJianUserName, Status:simpleApprovalOrder.Status}})
	//if err ,info := simpleApprovalOrderService.GetSimpleApprovalOrder(&model.SimpleApprovalOrder{JinJianId:simpleApprovalOrder.JinJianId}); err == nil{
	//	return err, info
	//}
	if err, tmp := simpleApprovalOrderService.GetSimpleApprovalOrder(&model.SimpleApprovalOrder{JinJianId: simpleApprovalOrder.JinJianId}); err != nil {
		//if info := QuerySimpleApprovalOrderById(simpleApprovalOrder.JinJianId); info != nil {
		//return err, tmp
	}else{
		if tmp.ID> 0 {

		}
	}


}



// 更新订单
func UpdateSimpleApprovalOrder(context *gin.Context) {
	var simpleApprovalOrder *model.SimpleApprovalOrder
	if context.BindJSON(&simpleApprovalOrder) == nil && simpleApprovalOrder != nil {
		if err := simpleApprovalOrderService.UpdateSimpleApprovalOrder(simpleApprovalOrder);err == nil {
			util.RenderGinSuccessJson("修改成功", nil, context)
		}else {
			util.RenderGinErrorJson(err.Error(), nil, context)
		}
	}else {
		util.RenderGinErrorJson("参数解析错误", nil, context)
	}
}



