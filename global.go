package common

import(
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	TRANSACTION_SERVER_PORT = "8081"
	REPAYMENT_SERVER_PORT   = "8082"
	PAYMENT_SERVER_PORT     = "8083"
	SPECIAL_PAYMENT_SERVER_PORT    = "8084"
	OVERDUE_SERVER_PORT = "8085"
	ACCOUNTING_SERVER_PORT  = "8086"
	// 支付网关
	GPAYMENT_SERVER_PORT = "8071"
	// 财务/催收系统
	GFIN_SERVER_PORT = "8072"

	// 算出来的钱保留多少位小数（为了精确计算利息）
	KeepFigures = 9
)

// 所有docker中的dns
const(
	AccountingMicroDns = "accounting-app"
	CoreSysApigwMicroDns = "core-sys-apigw-app"
	OverdueMicroDns = "overdue-app"
	PaymentMicroDns = "payment-app"
	RepaymentMicroDns = "repayment-app"
	SpecialPaymentMicroDns = "special-payment-app"
	TransactionMicroDns = "transaction-app"
	GssoMicroDns = "gsso-app"
	ConsulMicroDns = "consul-master-app"
	NqiyuanMicroDns = "nqiyuan-app"
	GriskcontrolMicroDns = "griskcontrol-app"
	GriskcontrolApigwMicroDns = "riskc-apigw-app"
	GcontractMicroDns = "gcontract-app"
	GgcallcenterMicroDns = "gcallcenter-app"
	GapprovalMicroDns = "gapproval-app"
	GpaymentMicroDns = "gpayment-app"
	GscoreMicroDns = "gscore-app"
	GoploggerMicroDns = "goplogger-http"
	GprodmngMicroDns = "gprodmng-app"
)

// 运行环境
const(
	EnvTest = "test"
	EnvDev = "dev"
	EnvProd = "prod"
)

// 应用的配置
type AppConfig struct {
	RedisUrl string
	MysqlHost string
	MysqlPort string
	MysqlUname string
	// http服务的端口（在docker中都是3000）
	HttpServerPort string
	MongoHost string
	MongoUname string
}

// 定义公用的flag，必须放在自己定义flag之后，因为会在这里边parse
// env是每个应用启动时都需要用到的，因此这里返回出去
func DefineRunTimeCommonFlag() (string) {
	env := flag.String("env", "dev", "数据库配置：test，dev，prod")
	dockerEnv := flag.String("docker_env", "0", "程序运行环境配置：0非docker，1docker非生产，2docker生产")
	flag.Parse()
	// 如果是生产，那么配置必须是一样的
	if *dockerEnv == "2" && *env != "prod" {
		panic("docker_env是2时，env必须是prod，检查环境及配置是否正确")
	}
	return *env
}

// 定义数据库操作的flag
func DefineDbMigrateCommonFlag() {
	flag.String("action", "dev", "数据库配置：test，dev，prod")
}

// 获取use docker的配置，避免每次打包修改use docker配置
// 表明是否使用容器配置 0是本地开发   1 是 dev  2是 prod
func GetUseDocker() int {
	f := flag.Lookup("docker_env")
	if f == nil || f.Value.String() == "0" {
		// 非docker
		return 0
	} else if f.Value.String() == "2" {
		// 生产
		return 2
	} else {
		// 默认返回开发和测试的配置
		return 1
	}
}

// 打印当前的docker配置
func PrintCurDockerEnv() {
	switch GetUseDocker() {
	case 0:
		fmt.Println("当前UseDocker配置为0，为非docker环境")
	case 2:
		fmt.Println("当前UseDocker配置为2，为docker生产环境")
	default:
		fmt.Println("当前UseDocker配置为1，为docker非生产环境")
	}
}

// 获取http端口
func GetHttpServerPort(notDocker string) string {
	if GetUseDocker() != 0 {
		return GetAppConfig().HttpServerPort
	}
	return notDocker
}

var appConfig *AppConfig

func GetAppConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}
	useDocker := GetUseDocker()
	if useDocker == 1 {
		fmt.Println("采用的是devops容器的配置")
		appConfig = GetDevDockerConf()
	} else if useDocker == 2 {
		fmt.Println("采用的是prod生产的配置")
		appConfig = GetProdDockerConf()
	} else {
		fmt.Println("采用的是非docker环境的配置")
		appConfig = GetAppDefaultConf()
	}
	return appConfig
}
// 获取正常环境下的配置
func GetAppDefaultConf() *AppConfig {
	return &AppConfig{
		RedisUrl: "127.0.0.1:6379",
		MysqlHost: "localhost",
		MysqlPort: "3306",
		MysqlUname: "root",
		MongoHost: "localhost:27017",
		MongoUname: "root",
	}
}
// 获取容器中的配置
func GetDevDockerConf() *AppConfig {
	return &AppConfig{
		RedisUrl: "redis-master:6379",
		MysqlHost: "qy-mysql",
		//MysqlHost: "172.16.0.101",
		MysqlPort: "3306",
		MysqlUname: "qy",
		HttpServerPort: "3000",
		MongoHost: "",
		MongoUname: "",
	}
}

// 获取容器中的配置
func GetProdDockerConf() *AppConfig {
	uname := "qy"
	if data, err := ioutil.ReadFile("/usr/local/.db/mysql.uname"); err != nil {
		fmt.Println("读取mysql用户名文件出错:" + err.Error() + "。使用默认用户名。")
	}else{
		uname = strings.TrimSpace(string(data))
	}
	return &AppConfig{
		RedisUrl: "redis-master:6379",
		MysqlHost: "qy-mysql",
		//MysqlHost: "172.16.1.90",
		MysqlPort: "3306",
		MysqlUname: uname,
		HttpServerPort: "3000",
		MongoHost: "",
		MongoUname: "",
	}
}

// 获取consul token
func GetConsulToken() string {
	switch GetUseDocker() {
	case 0:
		return "root"
	default:
		return "qy_prod_root"
	}
}

// 获取consul address
func GetConsulAddress() string {
	if GetUseDocker() != 0 {
		return ConsulMicroDns + ":8500"
	} else {
		return "127.0.0.1:8500"
	}
}

// 获取prodmng的host
func GetGprodmngHost() string {
	if GetUseDocker() != 0 {
		return "http://" + GprodmngMicroDns + ":3000"
	} else {
		return "http://localhost:12000"
	}
}

func GetGriskGwHost() string {
	if GetUseDocker() != 0 {
		return "http://" + GriskcontrolApigwMicroDns + ":3000"
	} else {
		return "http://172.16.0.29:8062"
	}
}
