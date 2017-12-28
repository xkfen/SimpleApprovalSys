package simpleApprovalOrderService

import (
	"errors"
	"simpleApproval/model"
	"simpleApproval/db/config"
	"gcoresys/common/logger"
	"gcoresys/common/util"
	)

//创建订单
func CreateSimpleApprovalOrder(simpleApprovalOrder *model.SimpleApprovalOrder) (err error){
	if err := simpleApprovalOrder.CheckIsValidSimpleApprovalOrder(); err != nil {
		return err
	}
	if err, tmp := GetSimpleApprovalOrder(&model.SimpleApprovalOrder{JinJianId: simpleApprovalOrder.JinJianId}); err != nil {
		return err
	}else {
		if tmp.ID > 0 {
			if tmp.JinJianUserName != simpleApprovalOrder.JinJianUserName {
				return errors.New("该条记录已被 " + tmp.JinJianUserName + " 编辑")
			}
			simpleApprovalOrder = tmp
			return nil
		} else {
			if err = config.GetDb().Model(&model.SimpleApprovalOrder{}).Create(simpleApprovalOrder).Error; err != nil {
				logger.Error("创建失败", "info", err.Error())
				return errors.New("创建失败, 请联系管理员")
			}
			return nil
		}
	}
	return nil
}

//查询订单
func GetSimpleApprovalOrder(simpleApprovalOrder *model.SimpleApprovalOrder) (error, *model.SimpleApprovalOrder) {
	result := &model.SimpleApprovalOrder{}
	err := config.GetDb().Where(simpleApprovalOrder).First(result).Error

	if err != nil && err.Error() != "record not found" {
		logger.Warn("查询失败", "info", err.Error())
		return errors.New("查询失败, 请联系管理员"), result
	}
	return nil, result
}

//  获取初审列表 (包含搜索)
//func GetAllSimpleApprovalOrderList(typeKey string, status string, username string, name string,
//	sort string, condition string, page int) ([]model.SimpleApprovalOrder, int) {
//	//results := []*SimpleApprovalOrder{}
//	//err := config.GetDb().Where(simpleApprovalOrder).Find()
//	return transactGetApprovalList("cs", typeKey, status, username, name, sort, condition, page)
//}

//修改订单
func UpdateSimpleApprovalOrder(simpleApprovalOrder *model.SimpleApprovalOrder) (err error){
	if err = config.GetDb().Model(&model.SimpleApprovalOrder{}).Update(simpleApprovalOrder).Error; err != nil {
		logger.Error("修改失败", "info", err.Error())
		return errors.New("修改信息失败，请联系管理员")
	}
	return
}


//  transact func 获取审批列表和搜索
func transactGetApprovalList(approvalType string, typeKey string, status string, username string, name string,
	sort string, condition string, page int) (aoList []model.SimpleApprovalOrder, totalPage int) {
	offset := util.GetOffset(page, perPage)
	// 判定前端页面选择的是 全部、我的或者是历史 isSelectAll
	//username, name = isSelectAll(username, name, typeKey)
	sqlBase := config.GetDb().Model(aoList)
	switch approvalType {
	//case "cs":
	//	sqlBase = sqlBase.Where("(jin_jian_id = ? AND jin_jian_user_name = ? AND status in (?))",
	//		username, name, firstTrailStatusList(typeKey, status))
	//case "zs":
	//	sqlBase = sqlBase.Where("re_trail_id = ? AND re_trail_name = ? AND re_trail_status in (?)",
	//		username, name, reTrailStatusList(typeKey, status))
	//case "kf":
	//	sqlBase = sqlBase.Where("custom_service_id = ? AND custom_service_name = ? AND custom_service_status in (?)",
	//		username, name, customServiceStatusList(typeKey, status))
	}
	if condition != "" {
		condition = "%" + condition + "%"
		sqlBase = sqlBase.Where(
			"((jinjian_user_name LIKE ?) OR (show_id LIKE ?) OR (agency_name LIKE ?) "+
				"OR (agency_employee LIKE ?) OR (user_id_num LIKE ?))",
			condition, condition, condition, condition, condition)
	}
	if typeKey == "history" {
		if sort == "" {
			sqlBase = sqlBase.Order("created_at desc")
		}
	}
	if sort != "" {
		sqlBase = sqlBase.Order(sort)
	}
	// 获取总页数
	totalPage, _ = util.GetTotalPagesAndCount(sqlBase, &aoList, perPage)
	// 获取初审列表 -- 分页后的数据
	sqlBase.Offset(offset).Limit(perPage).Find(&aoList)
	return
}