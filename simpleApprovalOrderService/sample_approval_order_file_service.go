package simpleApprovalOrderService

import (
	"simpleApproval/model"
	"errors"
	"gcoresys/common/logger"
	"simpleApproval/db/config"
	)

func CreateOrderFile(orderFile *model.SimpleApprovalOrderFile) error {
	if err := orderFile.IsValidApprovalOrderFile(); err != nil {
		return err
	}
	if err, tmp := GetSimpleApprovalOrderFile(&model.SimpleApprovalOrderFile{FileId: orderFile.FileId}); err != nil {
		return err
	} else {
		if tmp.ID > 0 {
			//if tmp.Username != orderFile.Username {
			//	return errors.New("该条记录已被 " + tmp.Username + " 编辑")
			//}
			orderFile = tmp
			return nil
		} else {
			if err = config.GetDb().Model(&model.SimpleApprovalOrderFile{}).Create(orderFile).Error; err != nil {
				logger.Error("创建失败", "info", err.Error())
				return errors.New("创建失败, 请联系管理员")
			}
			return nil
		}
	}
}


func GetSimpleApprovalOrderFile(order *model.SimpleApprovalOrderFile) (error, *model.SimpleApprovalOrderFile) {
	result := &model.SimpleApprovalOrderFile{}
	err := config.GetDb().Where(order).First(result).Error
	if err != nil && err.Error() != "record not found" {
		logger.Warn("查询失败", "info", err.Error())
		return errors.New("查询失败, 请联系管理员"), result
	}
	return nil, result
}


func UploadApprovalFileV1(file *model.SimpleApprovalOrderFile, opName string) (err error) {
	if err = file.IsValidApprovalOrderFile(); err != nil {
		return
	}
	//asyncNewApprovalLog(&model.ApprovalLog{ApprovalJinjianId: file.JinjianId, ApprovalName: opName,
	//	ApprovalStatus: opName + "上传文件"  + "，文件名：" + file.FileName,
	//	ApprovalType: "cs"})
	if err = config.GetDb().Create(file).Error; err != nil {
		return
	}
	return
}