package simpleApprovalOrderService

import (
	"simpleApproval/db/config"
	"gcoresys/common/logger"
	"github.com/stretchr/testify/suite"
	"testing"
	"fmt"
	"time"
	"simpleApproval/model"
)

type testingSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	logger.InitLogger(logger.LvlDebug, nil)
	config.GetSimpleApprovalDbConfig("test")
	config.GetDb().LogMode(false)
	suite.Run(t, new(testingSuite))
}

func (s *testingSuite) SetupTest() {
	//config.ClearAllData()
	simpleApprovalOrder := model.SimpleApprovalOrder{
		JinJianId:         "123",
		JinJianUserName:     "123",
		Status:            model.ApprovalStatusPass,
	}
	err := CreateSimpleApprovalOrder(&simpleApprovalOrder)
	s.Equal(nil, err)
}

func (s *testingSuite) TearDownTest() {
	config.ClearAllData()
}

func (s *testingSuite) NewSimpleApprovalOrders(count int) {
	for i := 0; i < count; i++ {
		ao := model.GetDefaultSimpleApprovalOrder()
		ao.JinJianId = fmt.Sprintf("JinjianIdSLZ%v%v", i, time.Now().UnixNano() )
		ao.JinJianUserName = fmt.Sprintf("JinJianUserNameSLZ%v%v", i, time.Now().UnixNano())
		ao.Status = fmt.Sprintf("StatusSLZ%v%v", i, time.Now().UnixNano())
		time.Sleep(7 * time.Millisecond)
		s.NoError(CreateSimpleApprovalOrder(ao))
	}

}