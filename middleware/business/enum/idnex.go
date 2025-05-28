package enum

//// 模拟测试环境
//const (
//	Env = "dev"
//
//	PathRoot = "/home/yan/Documents/file/gofile/gotest/middleware/business" // zxagent 跟路径
//
//	PathXYAPP = PathRoot
//
//	// PathRootAgent agent跟目录文件
//	PathRootAgent = PathRoot
//
//	// PathTmp 临时文件
//	PathTmp = PathRoot
//)

// 正式环境
const (
	Env = "pro"
	//PathRoot zxagent 根路径
	PathRoot = "/xyapp/system/miner.plugin-zxagent.ipk"

	// PathRootAgent agent跟目录文件
	PathRootAgent = "/xyapp/system/miner.plugin-agent.ipk"

	//PathXYAPP 配置文件路径
	PathXYAPP = "/etc/xyapp"

	// PathTmp 临时文件
	PathTmp = "/etc"
)

// 配置文件路径常量
const (
	// PathLogFile zxagent 日志路径
	PathLogFile = PathRoot + "/logs/app.log"

	// PathBwFile zxagent 数据传输路径
	PathBwFile = PathRoot + "/bw"

	// PathBwTmpFile zxagent 存储临时数据路径
	PathBwTmpFile = PathRoot + "/bw_tmp.json"

	// PathBaseFile zxagent 存储临时数据路径
	PathBaseFile = PathRoot + "/base.yaml"

	// PathRecruitResultFile 存储不同类型业务的appid，如抖音，混跑，g3
	PathRecruitResultFile = PathXYAPP + "/recruitResult.json"

	PathStopScriptsFile  = PathRootAgent + "/stop.sh"
	PathStartScriptsFile = PathRootAgent + "/start.sh"

	PathBusinessMacFile = PathXYAPP + "/businessMac.json"
	PathNetworkV2File   = PathXYAPP + "/network_v2.json"

	// PathBizConfFile 容器控制文件
	PathBizConfFile = PathXYAPP + "/biz_conf.json"

	// PathZXBIZFile 容器控制文件
	PathZXBIZFile = PathTmp + "/zx_biz"
)

// 业务标识
const (
	BusinessTypeDounYIN = "52d531d3ea193a292485d06517b4b5fd" // 抖音 appid 标识
	BusinessTypeMixRun  = "be37b71de68ba3339cc196b6ef802706" // 混跑 appid 标识
)

// 默认值
const (

	// DefaultNic 网卡默认值
	DefaultNic = "eth0"
)
