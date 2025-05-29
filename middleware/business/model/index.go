package model

import "encoding/json"

// base_config 文件模型
type (
	// Config 定义全局配置结构
	Config struct {
		Bandwidth BandwidthConfig `yaml:"bandwidth"`  // 带宽调整基础配置
		AgentSche AgentSchedule   `yaml:"agent_sche"` // 服务端信息
	}

	// BandwidthConfig 带宽相关配置
	BandwidthConfig struct {
		CheckInterval         int    `yaml:"check_interval"`          // 检查间隔(秒)
		BandwidthInterval     int    `yaml:"bandwidth_interval"`      // 带宽调整间隔(分钟)
		BandwidthIntervalDesc string `yaml:"bandwidth_interval_desc"` // 带宽间隔描述
		NetworkFile           string `yaml:"network_file"`            // 网络配置文件路径
		NetworkFileDesc       string `yaml:"network_file_desc"`       // 网络文件描述
	}

	// AgentSchedule 代理调度配置
	AgentSchedule struct {
		URL string `yaml:"url"` // 调度接口地址
	}
)

// 程序需要获取文件模型
type (
	// BizConf 容器控制实例
	BizConf struct {
		InstanceCount uint16 `json:"instance_count"` // 容器启动数量
	}

	// BwTmp 临时带宽记录数据
	BwTmp struct {
		Percentage  uint8   `json:"percentage"`  // 占比只用填 【1:1，1/2】、【1:2,1/3】
		NetworkCard string  `json:"networkCard"` // 网卡名
		Bandwidth   float64 `json:"bandwidth"`   // 带宽大小
		Count       uint16  `json:"count"`       // 容器数量
		AppID       string  `json:"appID"`       // appID
		UpdateAt    string  `json:"updateAt"`    // 修改时间
	}
)

// 服务端接口模型
type (
	// BandHost 抖音带宽控制接口参数
	BandHost struct {
		Host      string  `json:"host"`      // 节点host
		BandWidth float64 `json:"bandwidth"` // 节点规定带宽
		BandWay   float64 `json:"band_way"`  // 0-调度贷款 1-业务评估带宽 2-规定带宽
		From      int     `json:"from"`      // 1-运控， 2-api，其他
	}

	// HostInfoReport 上报机房带宽信息
	HostInfoReport struct {
		HostName  string  `json:"hostname"`  // 主机名
		BandWidth float64 `json:"bandwidth"` // 主机带宽
		Appid     string  `json:"appid"`     // 业务appID
	}

	// HostInfoParams 获取机房带宽信息
	HostInfoParams struct {
		HostName string `json:"hostname"` // 主机名
		Appid    string `json:"appid"`    // 业务appid
	}

	// ResMessage 接口相应信息
	ResMessage struct {
		Code    int             `json:"code"`    // 相应码
		Message string          `json:"message"` // 响应信息
		Data    json.RawMessage `json:"data"`    // 相应数据
	}

	// BandwidthInfo 控制带宽信息
	BandwidthInfo struct {
		Bandwidth     float64 `json:"bandwidth"`      // 要调整到的带宽值
		BandwidthOrig float64 `json:"bandwidth_orig"` // 当前主机带宽值
		LastBandwidth float64 `json:"last_bandwidth"` // 上一次上报的带宽值
		Ret           int     `json:"ret"`            // 操作码 0不调整 1调整到（bandwidth）目标值 -1调整为0【忽略bandwidth、bandwidth_orig值，强制设置为0】
	}
)
