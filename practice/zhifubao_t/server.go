package zhifubao_t

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"net/url"

	"github.com/skip2/go-qrcode"
)

// Server 支付宝支付演示 HTTP 服务，覆盖扫码下单、异步通知、查询与退款
type Server struct {
	client *Client     // 支付宝客户端
	orders *OrderStore // 订单存储
	config Config      // 应用配置
}

// NewServer 创建支付演示服务
func NewServer(client *Client, config Config) *Server {
	return &Server{
		client: client,
		orders: NewOrderStore(),
		config: config,
	}
}

// Handler 返回注册了全部支付相关路由的 HTTP Handler
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/pay", s.handlePay)
	mux.HandleFunc("/pay/notify", s.handleNotify)
	mux.HandleFunc("/pay/query", s.handleQuery)
	mux.HandleFunc("/pay/refund", s.handleRefund)
	return mux
}

// handleIndex 首页，展示各接口的用法
func (s *Server) handleIndex(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, `<html><body>
<h3>支付宝扫码支付流程演示</h3>
<ul>
<li>扫码下单：<a href="/pay?amount=0.01&amp;subject=测试商品">/pay?amount=0.01&amp;subject=测试商品</a></li>
<li>异步通知：/pay/notify（支付宝服务器调用）</li>
<li>查询：/pay/query?out_trade_no=订单号</li>
<li>退款：POST /pay/refund</li>
</ul>
</body></html>`)
}

// handlePay 发起扫码支付：创建订单、调用预下单接口获取二维码并渲染支付页，
// 页面轮询查询接口，检测到支付成功后自动展示结果。
func (s *Server) handlePay(w http.ResponseWriter, r *http.Request) {
	amount := r.URL.Query().Get("amount")
	if amount == "" {
		amount = "0.01"
	}
	subject := r.URL.Query().Get("subject")
	if subject == "" {
		subject = "测试商品"
	}

	order := s.orders.Create(subject, amount)
	result, err := s.client.Precreate(PrecreateParams{
		OutTradeNo:  order.OutTradeNo,
		TotalAmount: order.TotalAmount,
		Subject:     order.Subject,
	})
	if err != nil {
		http.Error(w, "下单失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
	png, err := qrcode.Encode(result.QrCode, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "生成二维码失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
	qrData := base64.StdEncoding.EncodeToString(png)
	outTradeNo := order.OutTradeNo
	fmt.Fprintf(w, `<html><body>
<h3>订单创建成功，请使用支付宝扫码付款</h3>
<p>商户订单号：%s</p>
<p>金额：%s 元</p>
<p>标题：%s</p>
<p><img src="data:image/png;base64,%s" alt="支付宝收款二维码"></p>
<div id="result"></div>
<script>
var outTradeNo = %q;
var timer = setInterval(function() {
  fetch('/pay/query?out_trade_no=' + encodeURIComponent(outTradeNo))
    .then(function(resp) { return resp.json(); })
    .then(function(data) {
      if (data.trade_status === 'TRADE_SUCCESS') {
        clearInterval(timer);
        document.getElementById('result').innerHTML =
          '<h3>支付成功</h3>' +
          '<p>订单号：' + data.out_trade_no + '</p>' +
          '<p>支付宝交易号：' + data.trade_no + '</p>' +
          '<p><a href="/">返回首页</a></p>';
      }
    }).catch(function() {});
}, 3000);
</script>
</body></html>`,
		order.OutTradeNo, order.TotalAmount, html.EscapeString(order.Subject),
		qrData, outTradeNo)
}

// handleNotify 处理支付宝异步通知：验签并完成业务校验后返回 success，否则返回 fail。
// 支付宝收到 fail 后会按策略重试通知。
func (s *Server) handleNotify(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "fail")
		return
	}
	params, err := s.client.VerifyNotify(r.PostForm)
	if err != nil {
		fmt.Fprint(w, "fail")
		return
	}
	if err := s.checkNotifyBusiness(params); err != nil {
		fmt.Fprint(w, "fail")
		return
	}
	fmt.Fprint(w, "success")
}

// checkNotifyBusiness 校验异步通知的业务字段：app_id、订单号、金额，并更新订单状态
func (s *Server) checkNotifyBusiness(params url.Values) error {
	if params.Get("app_id") != s.config.AppID {
		return errors.New("异步通知 app_id 与配置不一致")
	}
	outTradeNo := params.Get("out_trade_no")
	order, ok := s.orders.Get(outTradeNo)
	if !ok {
		return fmt.Errorf("异步通知订单不存在: %s", outTradeNo)
	}
	if params.Get("total_amount") != order.TotalAmount {
		return fmt.Errorf("异步通知金额 %s 与订单金额 %s 不一致",
			params.Get("total_amount"), order.TotalAmount)
	}
	switch params.Get("trade_status") {
	case "TRADE_SUCCESS", "TRADE_FINISHED":
		return s.orders.MarkPaid(outTradeNo)
	default:
		// 其余状态（如等待付款）按支付宝要求返回 success，避免重复通知
		return nil
	}
}

// handleQuery 主动查询订单支付状态并返回 JSON
func (s *Server) handleQuery(w http.ResponseWriter, r *http.Request) {
	outTradeNo := r.URL.Query().Get("out_trade_no")
	if outTradeNo == "" {
		http.Error(w, "缺少参数 out_trade_no", http.StatusBadRequest)
		return
	}
	status, err := s.client.QueryTrade(TradeQueryParams{OutTradeNo: outTradeNo})
	if err != nil {
		http.Error(w, "查询失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, status)
}

// refundRequest 退款请求体
type refundRequest struct {
	OutTradeNo   string `json:"out_trade_no"`   // 商户订单号
	RefundAmount string `json:"refund_amount"`  // 退款金额
	OutRequestNo string `json:"out_request_no"` // 退款请求号，可空
}

// handleRefund 发起退款并将结果以 JSON 返回
func (s *Server) handleRefund(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "请使用 POST 请求", http.StatusMethodNotAllowed)
		return
	}
	var req refundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "请求体解析失败: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.OutTradeNo == "" || req.RefundAmount == "" {
		http.Error(w, "缺少参数 out_trade_no 或 refund_amount", http.StatusBadRequest)
		return
	}
	result, err := s.client.Refund(RefundParams{
		OutTradeNo:   req.OutTradeNo,
		RefundAmount: req.RefundAmount,
		OutRequestNo: req.OutRequestNo,
	})
	if err != nil {
		http.Error(w, "退款失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_ = s.orders.MarkRefunded(req.OutTradeNo)
	writeJSON(w, result)
}

// renderPage 渲染一个简单的 HTML 结果页
func (s *Server) renderPage(w http.ResponseWriter, title, body string) {
	fmt.Fprintf(w, `<html><body>
<h3>%s</h3>
<p>%s</p>
<p><a href="/">返回首页</a></p>
</body></html>`, html.EscapeString(title), body)
}

// writeJSON 以 JSON 格式输出响应
func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(data)
}
