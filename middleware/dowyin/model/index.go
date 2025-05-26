package model

type Result struct {
	NetworkCard string  // 网卡名
	Bandwidth   float64 // 带宽大小
	Action      uint    // 状态码
	AppID       string  // appID
}

type BandHost struct {
	Host      string  `json:"host"`      // 节点host
	BandWidth float64 `json:"bandwidth"` // 节点规定带宽
	BandWay   float64 `json:"band_way"`  // 0-调度贷款 1-业务评估带宽 2-规定带宽
	From      int     `json:"from"`      // 1-运控， 2-api，其他
}
