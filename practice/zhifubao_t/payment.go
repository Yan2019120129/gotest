package zhifubao_t

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// checkBizResult 检查支付宝业务响应是否成功（code=10000 表示成功）
func checkBizResult(code, msg, subMsg string) error {
	if code == "10000" {
		return nil
	}
	return fmt.Errorf("支付宝业务处理失败: code=%s msg=%s sub_msg=%s", code, msg, subMsg)
}

// PrecreateParams 扫码支付预下单参数
type PrecreateParams struct {
	OutTradeNo  string // 商户订单号，需保证唯一
	TotalAmount string // 订单金额，单位元，最多两位小数
	Subject     string // 订单标题
}

// PrecreateResult 扫码支付预下单结果
type PrecreateResult struct {
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
	QrCode     string `json:"qr_code"`      // 二维码内容，用户使用支付宝扫码即可付款
}

// precreateResponse 扫码支付预下单接口响应
type precreateResponse struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	SubMsg string `json:"sub_msg"`
	PrecreateResult
}

// Precreate 扫码支付预下单：alipay.trade.precreate 返回二维码内容，
// 商户将其渲染为二维码展示给用户，用户使用支付宝扫码完成付款。
func (c *Client) Precreate(params PrecreateParams) (*PrecreateResult, error) {
	raw, err := c.doRequest(MethodTradePrecreate, map[string]any{
		"out_trade_no": params.OutTradeNo,
		"total_amount": params.TotalAmount,
		"subject":      params.Subject,
	})
	if err != nil {
		return nil, err
	}
	var resp precreateResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("解析预下单响应失败: %w", err)
	}
	if err := checkBizResult(resp.Code, resp.Msg, resp.SubMsg); err != nil {
		return nil, err
	}
	return &resp.PrecreateResult, nil
}

// TradeQueryParams 交易查询参数
type TradeQueryParams struct {
	OutTradeNo string // 商户订单号
}

// TradeStatus 交易查询返回的交易状态信息
type TradeStatus struct {
	TradeNo      string `json:"trade_no"`       // 支付宝交易号
	OutTradeNo   string `json:"out_trade_no"`   // 商户订单号
	TradeStatus  string `json:"trade_status"`   // 交易状态
	TotalAmount  string `json:"total_amount"`   // 订单金额
	BuyerLogonID string `json:"buyer_logon_id"` // 买家支付宝账号
	SendPayDate  string `json:"send_pay_date"`  // 支付时间
	RefundFee    string `json:"refund_fee"`     // 累计退款金额
}

// tradeQueryResponse 交易查询接口响应
type tradeQueryResponse struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	SubMsg string `json:"sub_msg"`
	TradeStatus
}

// QueryTrade 主动查询订单在支付宝侧的支付状态
func (c *Client) QueryTrade(params TradeQueryParams) (*TradeStatus, error) {
	raw, err := c.doRequest(MethodTradeQuery, map[string]any{
		"out_trade_no": params.OutTradeNo,
	})
	if err != nil {
		return nil, err
	}
	var resp tradeQueryResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("解析查询响应失败: %w", err)
	}
	if err := checkBizResult(resp.Code, resp.Msg, resp.SubMsg); err != nil {
		return nil, err
	}
	return &resp.TradeStatus, nil
}

// RefundParams 退款参数
type RefundParams struct {
	OutTradeNo   string // 商户订单号
	RefundAmount string // 退款金额，单位元，不能超过订单金额
	OutRequestNo string // 退款请求号，同一笔订单多次退款需区分
}

// RefundResult 退款结果
type RefundResult struct {
	TradeNo      string `json:"trade_no"`       // 支付宝交易号
	OutTradeNo   string `json:"out_trade_no"`   // 商户订单号
	RefundFee    string `json:"refund_fee"`     // 实际退款金额
	BuyerLogonID string `json:"buyer_logon_id"` // 买家支付宝账号
}

// refundResponse 退款接口响应
type refundResponse struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	SubMsg string `json:"sub_msg"`
	RefundResult
}

// Refund 发起退款申请
func (c *Client) Refund(params RefundParams) (*RefundResult, error) {
	if params.OutRequestNo == "" {
		params.OutRequestNo = params.OutTradeNo + "-refund"
	}
	raw, err := c.doRequest(MethodTradeRefund, map[string]any{
		"out_trade_no":   params.OutTradeNo,
		"refund_amount":  params.RefundAmount,
		"out_request_no": params.OutRequestNo,
	})
	if err != nil {
		return nil, err
	}
	var resp refundResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("解析退款响应失败: %w", err)
	}
	if err := checkBizResult(resp.Code, resp.Msg, resp.SubMsg); err != nil {
		return nil, err
	}
	return &resp.RefundResult, nil
}

// verifyCallback 校验支付宝异步通知参数签名，
// 验签通过后返回去掉 sign/sign_type 的业务参数集合
func (c *Client) verifyCallback(params url.Values) (url.Values, error) {
	clean := url.Values{}
	for key, values := range params {
		clean[key] = append([]string(nil), values...)
	}
	sign := clean.Get("sign")
	clean.Del("sign")
	clean.Del("sign_type")

	if err := c.verifySign(clean, sign); err != nil {
		return nil, err
	}
	return clean, nil
}

// VerifyNotify 校验支付宝异步通知参数，验签通过后返回业务参数集合
func (c *Client) VerifyNotify(params url.Values) (url.Values, error) {
	return c.verifyCallback(params)
}
