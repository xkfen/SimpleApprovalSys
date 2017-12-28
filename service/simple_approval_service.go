package service

import (
		"errors"
		"simpleApproval/model"
		)

func CreateApprovalOrder(approvalOrder *model.SimpleApprovalOrder) (err error){
	if err := approvalOrder.CheckIsValidSimpleApprovalOrder(); err != nil {
		return err
	}
}
