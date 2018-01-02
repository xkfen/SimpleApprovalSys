package simpleApprovalOrderService

import (
	"simpleApproval/model"
	"errors"
	"gcoresys/common/logger"
	"simpleApproval/db/config"
	)

func UploadOrderFile(orderFile *model.SimpleApprovalOrderFile) error {
	if err := orderFile.IsValidApprovalOrderFile(); err != nil {
		return err
	}
	if err, tmp := QuerySimpleApprovalOrderFile(&model.SimpleApprovalOrderFile{FileId: orderFile.FileId}); err != nil {
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

// 查询文件
func QuerySimpleApprovalOrderFile(order *model.SimpleApprovalOrderFile) (error, *model.SimpleApprovalOrderFile) {
	result := &model.SimpleApprovalOrderFile{}
	err := config.GetDb().Model(&model.SimpleApprovalOrderFile{}).Where(order).First(result).Error
	if err != nil && err.Error() != "record not found" {
		logger.Warn("查询失败", "info", err.Error())
		return errors.New("查询失败, 请联系管理员"), result
	}
	return nil, result
}


// 查询文件
func QueryOrderFile(fileUrl string) (error, *model.SimpleApprovalOrderFile) {
	result := &model.SimpleApprovalOrderFile{}
	err := config.GetDb().Model(&model.SimpleApprovalOrderFile{}).Where("file_url = ?", fileUrl).First(result).Error
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
	if err = config.GetDb().Model(&model.SimpleApprovalOrderFile{}).Create(file).Error; err != nil {
		return
	}
	return
}

// 测试文件下载
func DownloadFile(fileUrl string) (err error){
	if fileUrl == "" {
		logger.Info("文件路径不能为空", "info", err.Error())
		return errors.New("查询失败, 请联系管理员")
	}

	if tempErr := config.GetDb().Model(&model.SimpleApprovalOrderFile{}).Where("file_url",fileUrl).Error;tempErr != nil{
		logger.Info("文件不存在", "info", err.Error())
		return errors.New("文件不存在, 请检查文件路径")
	}
	return

}