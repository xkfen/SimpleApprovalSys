package router

import (
	"time"
	"github.com/gavv/httpexpect"
	"simpleApproval/model"
	"github.com/stretchr/testify/suite"
	"testing"
	"gcoresys/common/logger"
	"simpleApproval/db/config"

	"encoding/json"
)

type iHttpReqSuite struct {
	suite.Suite
}

type DataJson struct {
	JinJianId       uint   `json:"jin_jian_id"`
	JinJianUserName string `json:"jin_jian_user_name"`
	Status          string `json:"status"`
}

type RespJson struct {
	Data    DataJson `json:"data"`
	Info    string   `json:"info"`
	Success bool     `json:"success"`
}

const (
	baseURL = "http://localhost:7005/api/v1"
)

func (s *iHttpReqSuite) SetupTest() {
}

func (s *iHttpReqSuite) TearDownTest() {
	config.ClearAllData()
}

func TestRun(t *testing.T) {
	logger.InitLogger(logger.LvlDebug, nil)
	config.GetSimpleApprovalDbConfig("test")
	config.GetDb().LogMode(true)
	go StartHttpRouter()
	suite.Run(t, new(iHttpReqSuite))
}

func (s *iHttpReqSuite) TestCreateSimpleApprovalOrder001() {
	time.Sleep(1000 * time.Millisecond)
	resp := httpexpect.New(s.T(), baseURL).
		POST("/createSimpleApprovalOrder").
		WithJSON(model.SimpleApprovalOrder{
		JinJianId:       "12344",
		JinJianUserName: "123",
		Status:          "123",
	}).Expect()
	resp.Status(200).JSON()
	logger.Info("res", "data", resp.Body().Raw())
	var respJson RespJson
	err := json.Unmarshal([]byte(resp.Body().Raw()), &respJson)
	s.Equal(nil, err)
	s.Equal(true, respJson.Success)
}
