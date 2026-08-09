# zhifubao_t 支付宝支付流程演示

使用 Go 标准库实现支付宝开放平台 RSA2 签名/验签与完整支付流程，不依赖第三方支付宝 SDK，便于理解支付对接原理。

## 覆盖的支付过程

1. 扫码支付下单 `alipay.trade.precreate`：生成带 RSA2 签名的预下单请求，返回二维码内容并渲染为二维码展示
2. 用户使用支付宝扫码付款，支付页轮询 `/pay/query` 自动确认结果
3. 异步通知 `/pay/notify`：支付宝服务器回调，验签 + 业务校验（app_id、订单号、金额），确认订单并返回 `success`
4. 主动查询 `alipay.trade.query`：按商户订单号查询支付宝侧交易状态
5. 退款 `alipay.trade.refund`：按订单号发起退款

## 目录结构

- `config.go`：配置加载，读取 `config.yaml` 的 `alipay` 节点并填充默认值
- `client.go`：客户端核心，密钥解析、RSA2 签名/验签、网关请求与响应验签
- `payment.go`：扫码预下单、查询、退款、异步通知验签等支付业务方法
- `order.go`：内存订单存储，模拟商户订单库（金额校验、状态流转）
- `server.go`：HTTP 演示服务，暴露 `/pay`、`/pay/notify`、`/pay/query`、`/pay/refund`
- `demo/`：可运行入口
- `alipay_test.go`：签名/验签、通知校验、配置加载等单元测试

## 运行方式

1. 在 `config.yaml` 中填入沙箱应用的 APPID、应用私钥（PKCS8）与支付宝公钥
2. 将 `notify_url` 改为支付宝可访问的公网地址（本地开发可用内网穿透工具）
3. 启动服务：

   ```shell
   go run ./practice/zhifubao_t/demo
   ```

4. 浏览器访问 `http://127.0.0.1:18081/pay` 发起下单，使用支付宝扫码完成支付

## 测试

```shell
go test ./practice/zhifubao_t/ -v
```

## 关键点

- RSA2 即 SHA256WithRSA，签名内容为参数按 key 升序拼接的 `k1=v1&k2=v2` 字符串（空值不参与）
- 请求签名包含 `sign_type`；异步通知验签时需剔除 `sign` 与 `sign_type`
- 网关 JSON 响应的签名是对业务响应字段（如 `alipay_trade_query_response`）原文的签名
- 异步通知必须校验 `app_id`、`out_trade_no`、`total_amount`，成功应答固定返回文本 `success`
- 扫码支付无同步跳转，页面轮询查询接口仅用于展示，订单状态以异步通知为准
