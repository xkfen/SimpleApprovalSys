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

}

func (s *testingSuite) TearDownTest() {
	config.ClearAllData()
}


// 创建订单
func (s *testingSuite) TestCreateOrder(){
	simpleApprovalOrder := model.SimpleApprovalOrder{
		JinJianId:         "123",
		JinJianUserName:     "123",
		Status:            model.ApprovalStatusRefuse,
	}
	s.NoError(CreateSimpleApprovalOrder(&simpleApprovalOrder))
	fmt.Println(simpleApprovalOrder.JinJianId)
	fmt.Println(simpleApprovalOrder.JinJianUserName)
	fmt.Println(simpleApprovalOrder.Status)
	fmt.Println(simpleApprovalOrder.ID)
}


// 根据进件id查询订单
func (s *testingSuite) TestQueryOrderByJinJIanId(){
	ao := model.GetDefaultSimpleApprovalOrder()
	ao.JinJianId = "J20170616007"
	s.NoError(GetSimpleApprovalOrder(ao))
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

// 测试强制保存nil值到数据库
func (s *testingSuite) TestUpdateOrderWithBlankValues(){
	order := model.GetDefaultSimpleApprovalOrder()
	order.Status = model.ApprovalStatusRepulse
	order.JinJianUserName = "katy"
	err := UpdateOrderWithBlankValues(order)
	s.Equal(nil, err)
}

// 文件上传
func (s *testingSuite) TestUploadOrderFile(){
	file := model.GetDefaultSimpleApprovalOrderFile()

	s.NoError(UploadOrderFile(file))
	fmt.Println(file.FileId)
	fmt.Println(file.FileName)
	fmt.Println(file.FileUrl)
	fmt.Println(file.FileType)
	fmt.Println(file.Desc)
}

// 测试根据文件url查找文件
func (s *testingSuite) TestQueryOrderFile()  {
	file := model.GetDefaultSimpleApprovalOrderFile()
	s.NoError(QueryOrderFile(file.FileUrl))
	fmt.Println(file.FileUrl)
	fmt.Println(file.FileName)
	fmt.Println(file.FileType)
	fmt.Println(file.FileId)
	fmt.Println(file.Desc)
}

// 测试下载文件：根据文件url下载
func (s *testingSuite) TestDownloadFile(){
	file := model.GetDefaultSimpleApprovalOrderFile()
	s.NoError(DownloadFile(file.FileUrl))
	fmt.Println(file.FileUrl)
	fmt.Println(file.FileName)
	fmt.Println(file.FileType)
	fmt.Println(file.FileId)
	fmt.Println(file.Desc)
}
