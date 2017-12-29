package model

import (
	"gcoresys/common/mysql"
	"errors"
)


// 面签资料
type SimpleApprovalOrderFile struct {
	mysql.BaseModel
	// 订单id
	FileId string `gorm:"not null;index" json:"fileId"`
	// 文件类型
	FileType string `json:"fileType"`
	//文件名称
	FileName string `json:"fileName"`
	// 文件url
	FileUrl   string `json:"file_url"`
	// 文件描述
	Desc      string `gorm:"type:text" json:"desc"`
}

func (f *SimpleApprovalOrderFile) IsValidApprovalOrderFile() error {
	switch {
	case f.FileId == "":
		return errors.New("订单id不能为空")
	case f.FileType == "":
		return errors.New("文件类型不能为空")
	case f.FileUrl == "":
		return errors.New("文件不能为空")
	case f.FileName == "":
		return errors.New("文件名称不能为空")
	}
	return nil
}


func GetDefaultSimpleApprovalOrderFile() *SimpleApprovalOrderFile {
	return &SimpleApprovalOrderFile{
		FileId:   "AAAAA",
		FileName:  "测试文件",
		FileType:  "txt",
		FileUrl:   "/usr/local/.db/mysql.pas",
		Desc:      "文件描述",
	}
}
//func IsValidFileType(fileType string) error {
//	if fileType == Contract || fileType == IdCardA || fileType == IdCardB ||
//		fileType == HouseDeed || fileType == CreditReport || fileType == LoanBankCard ||
//		fileType == Sound || fileType == Policy || fileType == IdCardC {
//		return nil
//	}
//	return errors.New("文件类型错误")
//}