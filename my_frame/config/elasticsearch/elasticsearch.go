package Esearch

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/delete"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/get"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"gotest/my_frame/models"
	"log"
	"net"
	"net/http"
	"time"
)

var url = []string{"http://47.101.70.217:1014"}

var ES Esearch

// Init 初始化elasticsearch
func Init(cfg *models.Elasticsearch) {
	// 配置bing创建类型连接
	var err error
	ES.Clint, err = elasticsearch.NewTypedClient(
		elasticsearch.Config{
			Addresses: cfg.IpAddress,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   cfg.MaxIdleConnsPerHost,
				ResponseHeaderTimeout: time.Duration(cfg.ResponseHeaderTimeout) * time.Second,
				DialContext:           (&net.Dialer{Timeout: time.Duration(cfg.DialerTimeout) * time.Second}).DialContext,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		})
	if err != nil {
		panic(err)
	}
}

// Esearch 操作ES的结构体
type Esearch struct {
	Clint *elasticsearch.TypedClient
}

// GetEs 获取配置对象
// Addresses Elasticsearch 集群的地址，可以是一个或多个节点的地址
// Transport  定义 HTTP 连接的传输层配置
// MaxIdleConnsPerHost 每个主机的最大空闲连接数
// ResponseHeaderTimeout 接收响应头的超时时间
// DialContext 建立连接的上下文和超时设置
// TLSClientConfig TLS 配置，用于加密通信
// MinVersion  设置最小支持的 TLS 版本
func GetEs() *Esearch {
	var err error
	ES.Clint, err = elasticsearch.NewTypedClient(
		elasticsearch.Config{
			Addresses: url,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: 10 * time.Second,
				DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		})
	if err != nil {
		panic(err)
	}
	return &ES
}

// Insert 添加数据
// index.Response,error
//
//	type Response struct {
//		ForcedRefresh *bool 				// 是否强制刷新
//		Id_           string 				// 返回插入的文档 ID
//		Index_        string 				// 插入的索引名称
//		PrimaryTerm_  int64 				// 主要术语
//		Result        result.Result 		// 操作结果
//		SeqNo_        int64 				// 序列号
//		Shards_       types.ShardStatistics // 分片统计信息
//		Version_      int64 				// 文档版本号
//	}
func (es *Esearch) Insert(index, id string, data interface{}) (*index.Response, error) {
	// 添加数据
	return es.Clint.Index(index).Id(id).Request(data).Do(context.TODO())
}

// Update 修改数据
// update.Response, error
//
//	type Response struct {
//	   	ForcedRefresh *bool                  // 是否强制刷新
//		Get           *types.InlineGet       // 获取的文档信息
//		Id_           string                 // 更新的文档 ID
//		Index_        string                 // 更新的文档所属的索引名称
//		PrimaryTerm_  int64                  // 主要术语
//		Result        result.Result          // 操作结果
//		SeqNo_        int64                  // 序列号
//		Shards_       types.ShardStatistics  // 分片统计信息
//		Version_      int64                  // 更新后的文档版本号
//	}
func (es *Esearch) Update(index, id string, data interface{}) (*update.Response, error) {
	// 将需要修改的数据转换[]byte
	updateData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 发送请求
	return es.Clint.Update(index, id).Request(&update.Request{Doc: updateData}).Do(context.TODO())
}

// Delete 删除数据
// delete.Response, error
//
//	type Response struct {
//		ForcedRefresh *bool                 	 // 是否强制刷新
//		Id_           string                	 // 被删除文档的 ID
//		Index_        string                	 // 被删除文档所属的索引名称
//		PrimaryTerm_  int64                 	 // 主要术语
//		Result        result.Result         	 // 操作结果
//		SeqNo_        int64                 	 // 序列号
//		Shards_       types.ShardStatistics 	 // 分片统计信息
//		Version_      int64                 	 // 被删除文档的版本号
//	}
func (es *Esearch) Delete(index, id string) (*delete.Response, error) {
	return es.Clint.Delete(index, id).Do(context.TODO())
}

// Find 查询数据
// get.Response, error
//
//	type Response struct {
//		Fields       map[string]json.RawMessage // 字段映射
//		Found        bool                        // 是否找到文档
//		Id_          string                      // 查询的文档 ID
//		Index_       string                      // 查询的文档所属的索引名称
//		PrimaryTerm_ *int64                      // 主要术语
//		Routing_     *string                     // 路由信息
//		SeqNo_       *int64                      // 序列号
//		Source_      json.RawMessage             // 查询结果的原始 JSON 数据
//		Version_     *int64                      // 查询的文档版本号
//	}
func (es *Esearch) Find(index, id string) (*get.Response, error) {
	return es.Clint.Get(index, id).Do(context.TODO())
}

// FindAll 分页查找数据
// search.Response, error
// // Response 包含分页查询操作的响应信息
//
//	type Response struct {
//		Aggregations    map[string]types.Aggregate        // 聚合操作的结果
//		Clusters_       *types.ClusterStatistics          // 集群统计信息
//		Fields          map[string]json.RawMessage        // 字段映射
//		Hits            types.HitsMetadata                // 命中的文档信息
//		MaxScore        *types.Float64                    // 最大分数
//		NumReducePhases *int64                            // 减少的阶段数
//		PitId           *string                           // 持久化搜索标识符
//		Profile         *types.Profile                    // 操作的性能分析信息
//		ScrollId_       *string                           // 滚动标识符
//		Shards_         types.ShardStatistics             // 分片统计信息
//		Suggest         map[string][]types.Suggest        // 建议操作的结果
//		TerminatedEarly *bool                             // 是否提前终止
//		TimedOut        bool                              // 操作是否超时
//		Took            int64                             // 操作所花费的时间
//	}
func (es *Esearch) FindAll(index string, size int) (*search.Response, error) {
	return es.Clint.Search().Index(index).Request(
		&search.Request{
			Query: &types.Query{
				MatchAll: &types.MatchAllQuery{},
			},
		}).Scroll("5s").Sort("account_number").Size(size).Do(context.TODO())
}

// InsertMany 批量插入操作
//
//	type Response struct {
//		Errors     bool                                                  false 插入成功
//		IngestTook *int64
//		Items      []map[operationtype.OperationType]types.ResponseItem
//		Took       int64                                                 执行时间
//	}
func (es *Esearch) InsertMany(index string, data []*interface{}) (*bulk.Response, error) {
	log.Println("开始配置！！！")
	bulkRequest := es.Clint.Bulk()
	for _, v := range data {
		err := bulkRequest.CreateOp(types.CreateOperation{Index_: &index}, v)
		if err != nil {
			panic(err)
		}
	}
	log.Println("开始插入！！！")
	return bulkRequest.Do(context.Background())
}
