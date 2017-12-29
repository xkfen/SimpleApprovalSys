package model

import (
	"errors"
	"gcoresys/common/mysql"
)

const (
	ApprovalStatusPass    = "审核通过"
	ApprovalStatusRepulse = "审核打回"
	ApprovalStatusRefuse  = "审核拒绝"
)

type SimpleApprovalOrder struct {
	mysql.BaseModel
	JinJianId       string `gorm:"unique" json:"jin_jian_id"`
	JinJianUserName string `json:"jin_jian_user_name"`
	Status          string `json:"status"`
}

func (simpleApprovalOrder *SimpleApprovalOrder) CheckIsValidSimpleApprovalOrder() error {
	if simpleApprovalOrder.JinJianId == "" {
		return errors.New("id不能为空")
	}

	if simpleApprovalOrder.JinJianUserName == "" {
		return errors.New("user name不能为空")
	}

	if simpleApprovalOrder.Status == "" {
		return errors.New("status不能为空")
	}
	return nil
}

func GetDefaultSimpleApprovalOrder() *SimpleApprovalOrder {
	return &SimpleApprovalOrder{
		JinJianId:       "J20170616007",
		JinJianUserName: "test11111",
		Status:          ApprovalStatusPass,
	}
}

func (simpleApprovalOrder *SimpleApprovalOrder)updateValidateSimpleApprovalOrder() error{
	if simpleApprovalOrder.JinJianId == "" {
		return errors.New("id不能为空")
	}

	if simpleApprovalOrder.JinJianUserName == "" {
		return errors.New("user name不能为空")
	}

	if simpleApprovalOrder.Status == "" {
		return errors.New("status不能为空")
	}
	return nil

}