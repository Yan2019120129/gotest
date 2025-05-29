package model

import "encoding/json"

// BandHost 抖音带宽控制接口参数
type BandHost struct {
	Host      string  `json:"host"`      // 节点host
	BandWidth float64 `json:"bandwidth"` // 节点规定带宽
	BandWay   float64 `json:"band_way"`  // 0-调度贷款 1-业务评估带宽 2-规定带宽
	From      int     `json:"from"`      // 1-运控， 2-api，其他
}

// BwTmp 临时带宽记录数据
type BwTmp struct {
	Percentage  uint8   `json:"percentage"`  // 占比只用填 【1:1，1/2】、【1:2,1/3】
	NetworkCard string  `json:"networkCard"` // 网卡名
	Bandwidth   float64 `json:"bandwidth"`   // 带宽大小
	Count       uint16  `json:"count"`       // 容器数量
	AppID       string  `json:"appID"`       // appID
	UpdateAt    string  `json:"updateAt"`    // 修改时间
}

// Config 定义全局配置结构
type Config struct {
	Bandwidth BandwidthConfig `yaml:"bandwidth"`
	AgentSche AgentSchedule   `yaml:"agent_sche"`
}

// BandwidthConfig 带宽相关配置
type BandwidthConfig struct {
	CheckInterval         int    `yaml:"check_interval"`          // 检查间隔(秒)
	BandwidthInterval     int    `yaml:"bandwidth_interval"`      // 带宽调整间隔(分钟)
	BandwidthIntervalDesc string `yaml:"bandwidth_interval_desc"` // 带宽间隔描述
	NetworkFile           string `yaml:"network_file"`            // 网络配置文件路径
	NetworkFileDesc       string `yaml:"network_file_desc"`       // 网络文件描述
}

// AgentSchedule 代理调度配置
type AgentSchedule struct {
	URL string `yaml:"url"` // 调度接口地址
}

// HostInfoReport 上报机房带宽信息
type HostInfoReport struct {
	HostName  string  `json:"hostname"`
	BandWidth float64 `json:"bandwidth"`
	Appid     string  `json:"appid"`
}

// HostInfoParams 获取机房带宽信息
type HostInfoParams struct {
	HostName string `json:"hostname"`
	Appid    string `json:"appid"`
}

// ResMessage 接口相应信息
type ResMessage struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// BandwidthInfo 控制带宽信息
type BandwidthInfo struct {
	Bandwidth     float64 `json:"bandwidth"`
	BandwidthOrig float64 `json:"bandwidth_orig"`
	LastBandwidth float64 `json:"last_bandwidth"`
	Ret           int     `json:"ret"`
}

// BizConf 容器控制实例
type BizConf struct {
	InstanceCount uint16 `json:"instance_count"`
}
