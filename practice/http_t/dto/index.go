package dto

// EhParams 易汇获取产品K线图参数
type EhParams struct {
	Id   string        `json:"id"`
	Cmd  string        `json:"cmd"`
	Args []interface{} `json:"args"`
}

// RspData okx返回数据
type RspData struct {
	Code int         //	返回代码
	Msg  string      //	消息
	Data interface{} //	内容
}
