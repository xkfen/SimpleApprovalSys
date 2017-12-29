package main

import (
	"fmt"
	"gcoresys/common/mysql"
	"flag"
	"simpleApproval/db/config"
	"simpleApproval/model"
	"gcoresys/common"

)

func main() {
	common.DefineDbMigrateCommonFlag()
	env := common.DefineRunTimeCommonFlag()
	action := flag.Lookup("action").Value.String()
	switch action {
	case "create":
		doCreate(env)
	case "drop":
		doDrop(env)
	case "migrate":
		doMigrate(env)
	default:
		panic("未知操作")
	}
}

func doCreate(env string) {
	fmt.Println("do create ...")
	dbConfig := config.GetSimpleApprovalDbConfig(env)
	mysql.CreateDB(dbConfig)
}

func doDrop(env string) {
	fmt.Println("do drop ...")
	dbConfig := config.GetSimpleApprovalDbConfig(env)
	mysql.DropDB(dbConfig)
}

func doMigrate(env string) {
	fmt.Println("do migrate ...")
	config.GetSimpleApprovalDbConfig(env)
	db := config.GetDb()
	db.AutoMigrate(&model.SimpleApprovalOrder{}, &model.SimpleApprovalOrderFile{})
}
