package simpleApprovalOrderService

import(
	"simpleApproval/model"
	"fmt"
)

func (s *testingSuite) TestGetSimpleApprovalOrderList() {
	ao1 := model.GetDefaultSimpleApprovalOrder()
	s.NoError(CreateSimpleApprovalOrder(ao1))
	ao2 := model.GetDefaultSimpleApprovalOrder()
	ao2.JinJianId = ao2.JinJianId + "2"
	ao2.JinJianUserName = ao2.JinJianUserName + "2"
	ao2.Status = model.ApprovalStatusPass
	s.NoError(CreateSimpleApprovalOrder(ao2))
	ao3 := model.GetDefaultSimpleApprovalOrder()
	ao3.JinJianId = ao3.JinJianId + "3"
	ao3.JinJianUserName = ao3.JinJianUserName + "3"
	ao3.Status = model.ApprovalStatusRefuse
	s.NoError(CreateSimpleApprovalOrder(ao3))

	fmt.Println("ao3-status",ao3.Status)
	fmt.Println("ao3-jinjianid",ao3.JinJianId)
	fmt.Println("ao3-jinjianusername",ao3.JinJianUserName)
	//aoList, tp := GetAllSimpleApprovalOrderList(ao3);
	//s.Equal(1, tp)
	//s.Equal(true, len(aoList) > 0)
	//s.Equal(3, len(aoList))

}