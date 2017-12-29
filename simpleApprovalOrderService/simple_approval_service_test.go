package simpleApprovalOrderService

import (
	"simpleApproval/db/config"
	"gcoresys/common/logger"
	"github.com/stretchr/testify/suite"
	"testing"
	//"fmt"
	//"time"
	"simpleApproval/model"
	"fmt"
)

type testingSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	logger.InitLogger(logger.LvlDebug, nil)
	config.GetSimpleApprovalDbConfig("dev")
	config.GetDb().LogMode(false)
	suite.Run(t, new(testingSuite))
}

func (s *testingSuite) SetupTest() {
	////config.ClearAllData()
	//simpleApprovalOrder := model.SimpleApprovalOrder{
	//	JinJianId:         "123",
	//	JinJianUserName:     "123",
	//	Status:            model.ApprovalStatusPass,
	//}
	//err := CreateSimpleApprovalOrder(&simpleApprovalOrder)
	//s.Equal(nil, err)
}

func (s *testingSuite) TearDownTest() {
	config.ClearAllData()
}


// 创建订单
func (s *testingSuite) TestCreateOrder(){
	ao := model.GetDefaultSimpleApprovalOrder();
	ao.JinJianId = "kay1"
	ao.JinJianUserName = "112321344"
	ao.Status = model.ApprovalStatusRefuse

	s.NoError(CreateSimpleApprovalOrder(ao))
	fmt.Println(ao.JinJianId)
	fmt.Println(ao.JinJianUserName)
	fmt.Println(ao.Status)
	fmt.Println(ao.ID)
}


// 根据进件id查询订单
func (s *testingSuite) TestQueryOrderByJinJIanId(){
	ao := model.GetDefaultSimpleApprovalOrder()
	ao.JinJianId = "J20170616007"
	//JinJianId := "J20170616007"

	s.NoError(GetSimpleApprovalOrder(ao))
	//ao := QuerySimpleApprovalOrderById(JinJianId)
	//s.NoError(QuerySimpleApprovalOrderById(JinJianId))

	fmt.Println(ao.JinJianId)
	fmt.Println(ao.JinJianUserName)
	fmt.Println(ao.Status)
}

// 修改订单
func (s *testingSuite) TestUpdateOrder(){
	order := model.SimpleApprovalOrder{
		Status:  model.ApprovalStatusRefuse,
		JinJianId: "123",
		JinJianUserName:  "usyweudhefr",
	}

	err := UpdateSimpleApprovalOrder(&order)
	s.Equal(nil, err)
}

//// 测试强制保存nil值到数据库
func (s *testingSuite) TestUpdateOrderWithBlankValues(){
	order := model.GetDefaultSimpleApprovalOrder()
	order.Status = model.ApprovalStatusRepulse
	order.JinJianUserName = "katy"
	err := UpdateOrderWithBlankValues(order)
	s.Equal(nil, err)
}

// 文件上传
func (s *testingSuite) TestCreateOrderFile(){
	file := model.GetDefaultSimpleApprovalOrderFile()

	s.NoError(CreateOrderFile(file))
	fmt.Println(file.FileId)
	fmt.Println(file.FileName)
	fmt.Println(file.FileUrl)
	fmt.Println(file.FileType)
	fmt.Println(file.Desc)
}
//func (s *testingSuite) NewSimpleApprovalOrders(count int) {
//	for i := 0; i < count; i++ {
//		ao := model.GetDefaultSimpleApprovalOrder()
//		ao.JinJianId = fmt.Sprintf("JinjianIdSLZ%v%v", i, time.Now().UnixNano() )
//		ao.JinJianUserName = fmt.Sprintf("JinJianUserNameSLZ%v%v", i, time.Now().UnixNano())
//		ao.Status = fmt.Sprintf("StatusSLZ%v%v", i, time.Now().UnixNano())
//		time.Sleep(7 * time.Millisecond)
//		s.NoError(CreateSimpleApprovalOrder(ao))
//	}
//
//}

//func (s *testingSuite) QuerySimpleOrderByJinJianId(count int){
//	for i := 0; i < count; i++ {
//		ao := model.GetDefaultSimpleApprovalOrder()
//		ao.JinJianId = "12344"
//		//ao.JinJianId = fmt.Sprintf("JinjianIdSLZ%v%v", i, time.Now().UnixNano())
//		//ao.JinJianUserName = fmt.Sprintf("JinJianUserNameSLZ%v%v", i, time.Now().UnixNano())
//		//ao.Status = fmt.Sprintf("StatusSLZ%v%v", i, time.Now().UnixNano())
//		time.Sleep(7 * time.Millisecond)
//		s.NoError(GetSimpleApprovalOrder(ao))
//		fmt.Println(ao.JinJianId)
//		fmt.Println(ao.JinJianUserName)
//		fmt.Println(ao.Status)
//
//	}
//}

func TestGetSimpleApprovalOrder(t *testing.T) {
	err := CreateSimpleApprovalOrder(&model.SimpleApprovalOrder{})
	fmt.Println(err.Error())

}