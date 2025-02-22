package client

// Subscribe 订阅数据
type Subscribe struct {
	Name string //	通道名称
	Data []byte //	订阅数据
}
