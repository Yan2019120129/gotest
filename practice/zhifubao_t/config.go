// Package zhifubao_t 演示支付宝开放平台支付完整流程：
// 扫码支付下单 -> 用户扫码付款 -> 异步通知验签 -> 订单查询 -> 退款。
// 整个流程使用 Go 标准库实现 RSA2 签名与验签，不依赖第三方支付宝 SDK。
package zhifubao_t

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// 支付接口方法名与通用常量
const (
	// DefaultGatewayURL 支付宝沙箱网关地址
	DefaultGatewayURL = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"
	// SignTypeRSA2 RSA2 签名算法（SHA256WithRSA）
	SignTypeRSA2 = "RSA2"
	// FormatJSON 网关响应格式
	FormatJSON = "JSON"
	// CharsetUTF8 请求字符集
	CharsetUTF8 = "UTF-8"
	// VersionAPI 开放接口版本号
	VersionAPI = "1.0"

	// MethodTradePrecreate 扫码支付预下单接口
	MethodTradePrecreate = "alipay.trade.precreate"
	// MethodTradeQuery 交易查询接口
	MethodTradeQuery = "alipay.trade.query"
	// MethodTradeRefund 交易退款接口
	MethodTradeRefund = "alipay.trade.refund"
)

// Config 支付宝开放平台配置，对应 config.yaml 文件 alipay 节点下的字段
type Config struct {
	AppID           string `yaml:"app_id"`            // 沙箱应用 APPID
	PrivateKey      string `yaml:"private_key"`       // 应用私钥（PKCS8 格式）
	AlipayPublicKey string `yaml:"alipay_public_key"` // 支付宝公钥
	GatewayURL      string `yaml:"gateway_url"`       // 网关地址
	NotifyURL       string `yaml:"notify_url"`        // 支付结果异步通知地址
	Charset         string `yaml:"charset"`           // 字符集
	SignType        string `yaml:"sign_type"`         // 签名类型
	Format          string `yaml:"format"`            // 响应格式
}

// fileConfig 配置文件根结构，所有配置项位于 alipay 节点下
type fileConfig struct {
	Alipay Config `yaml:"alipay"`
}

// LoadConfig 从 YAML 文件加载支付宝配置，缺失的可选字段会填充默认值
func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("读取配置文件失败: %w", err)
	}
	var fc fileConfig
	if err := yaml.Unmarshal(data, &fc); err != nil {
		return Config{}, fmt.Errorf("解析配置文件失败: %w", err)
	}
	cfg := fc.Alipay
	cfg.applyDefaults()
	return cfg, nil
}

// applyDefaults 为可选的通用参数填充默认值
func (c *Config) applyDefaults() {
	if c.GatewayURL == "" {
		c.GatewayURL = DefaultGatewayURL
	}
	if c.Charset == "" {
		c.Charset = CharsetUTF8
	}
	if c.SignType == "" {
		c.SignType = SignTypeRSA2
	}
	if c.Format == "" {
		c.Format = FormatJSON
	}
}
