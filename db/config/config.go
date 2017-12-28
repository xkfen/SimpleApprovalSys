package config

import (
	"gcoresys/common/mysql"
	"github.com/jinzhu/gorm"
	"gcoresys/common/mysql/connection"
)

// 记录当前运行时的数据库配置
var curDbConfig *mysql.DbConfig

// 获取mysql数据库配置，此处还可以加到配置文件去获取配置的逻辑
func GetSimpleApprovalDbConfig(env string) *mysql.DbConfig {
	curDbConfig = mysql.NewDbConfig()
	switch env {
	case "prod":
		curDbConfig.DbName = "qy_simple_approval_prod"
	case "test":
		curDbConfig.DbName = "qy_simple_approval_test"
	default:
		curDbConfig.DbName = "qy_simple_approval_dev"
	}
	return curDbConfig
}

// 获取当前服务的数据库连接
func GetDb() *gorm.DB {
	if curDbConfig == nil {
		panic("请先初始化数据库配置")
	}
	return connection.GetDb(curDbConfig)
}

// 清空当前数据库的所有数据
func ClearAllData() {
	connection.ClearAllData(curDbConfig)
}