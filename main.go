package main

import (
	"flag"
	"gcoresys/common"
	"simpleApproval/router"
	"gcoresys/common/logger"
	"gopkg.in/gin-gonic/gin.v1"
	"gcoresys/common/util"
	"simpleApproval/db/config"
)

func main() {
	var (
		//ginMode = flag.String("ginMode", "debug", "DebugMode or ReleaseMode")
		port = flag.String("port", ":7005", "HTTP listen port")
	)
	env := common.DefineRunTimeCommonFlag()
	logger.Info("运行环境", "env", env)
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	}

	tmpPort := *port
	if common.GetUseDocker() != 0 {
		tmpPort = ":" + common.GetAppConfig().HttpServerPort
		logger.InitLogger(logger.LvlInfo, "qy_approval_apigw.log")
	} else {
		logger.InitLogger(logger.LvlDebug, nil)
	}

	util.WishNoBug()
	config.GetSimpleApprovalDbConfig("dev")
	config.GetDb()
	r := router.GetHttpRouter() //获得路由实例
	logger.Info("====api gateway 启动  端口 ： " + tmpPort + " ====")

	r.Run(tmpPort)
}
