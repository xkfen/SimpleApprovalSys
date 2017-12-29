package simpleApprovalOrderService

import (
	"errors"
	"simpleApproval/model"
	"simpleApproval/db/config"
	"gcoresys/common/logger"
	"fmt"
	"gcoresys/common/util"
)

// 创建订单
func CreateSimpleApprovalOrder(simpleApprovalOrder *model.SimpleApprovalOrder) (err error) {
	//if err = simpleApprovalOrder.CheckIsValidSimpleApprovalOrder(); err != nil {
	//	return
	//}
	//if err, tmp := GetSimpleApprovalOrder(&model.SimpleApprovalOrder{JinJianId: simpleApprovalOrder.JinJianId}); err != nil {
	//	return
	//} else {
	//
	//}

	err, tmp := GetSimpleApprovalOrder(&model.SimpleApprovalOrder{JinJianId: simpleApprovalOrder.JinJianId})
	if err != nil {
		return
	}
	if tmp.ID > 0 {
		if tmp.JinJianUserName != simpleApprovalOrder.JinJianUserName {
			return errors.New("该条记录已被 " + tmp.JinJianUserName + " 编辑")
		}
		simpleApprovalOrder = tmp
	} else {
		if err = config.GetDb().Model(&model.SimpleApprovalOrder{}).Create(simpleApprovalOrder).Error; err != nil {
			logger.Error("创建失败", "info", err.Error())
			return errors.New("创建失败, 请联系管理员")
		}
	}
	return
}

// 查询订单
func GetSimpleApprovalOrder(simpleApprovalOrder *model.SimpleApprovalOrder) (error, *model.SimpleApprovalOrder) {
	result := &model.SimpleApprovalOrder{}
	err := config.GetDb().Where(simpleApprovalOrder).First(result).Error

	if err != nil && err.Error() != "record not found" {
		logger.Warn("查询失败", "info", err.Error())
		return errors.New("查询失败, 请联系管理员"), result
	}
	return nil, result
}

//  根据进件id查询订单
func QuerySimpleApprovalOrderById(jinJianId string) (err error){
	if jinJianId== "" {
		err = errors.New("id不能为空, 请检查")
		return
	}
	var ao model.SimpleApprovalOrder
	if err = config.GetDb().Model(ao).Where("jin_jian_id =?", jinJianId).First(&ao).Error; err != nil {
		return err
	}
	return
}

// 修改订单,有nil值就不会更新
func UpdateSimpleApprovalOrder(simpleApprovalOrder *model.SimpleApprovalOrder) (err error) {
	// 检查查询条件
	//if err = simpleApprovalOrder.CheckIsValidSimpleApprovalOrder(); err != nil {
	//	return err
	//}
	// 有nil值就不会更新
	if err = config.GetDb().Model(&model.SimpleApprovalOrder{}).Where("jin_jian_id=?", simpleApprovalOrder.JinJianId).Update(simpleApprovalOrder).Error; err != nil {
		return
	}
	var tmpOrder model.SimpleApprovalOrder
	if err = config.GetDb().Model(&model.SimpleApprovalOrder{}).Where("jin_jian_id=?", simpleApprovalOrder.JinJianId).First(&tmpOrder).Error; err != nil {
		return
	}
	fmt.Println(util.StringifyJson(tmpOrder))
	return
}


// 修改订单(强制把前端传的nil值也一起更新到数据库)
func UpdateOrderWithBlankValues(simpleApprovalOrder *model.SimpleApprovalOrder) (err error) {
	if err = config.GetDb().Model(&model.SimpleApprovalOrder{}).Where("jin_jian_id=?", simpleApprovalOrder.JinJianId).Update(map[string]interface{}{
		"jin_jian_user_name": simpleApprovalOrder.JinJianUserName,
		"status": simpleApprovalOrder.Status,
	}).Error; err != nil {
		return
	}
	return
}

