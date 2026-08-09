### 🔐 JWT — JSON Web Token 身份认证令牌

> **定义**：基于 RFC 7519 标准的 JSON 格式安全令牌，用于在客户端与服务端之间传递身份认证和授权信息；JWT 本身不是加密算法，而是通过数字签名保证数据完整性。

* **核心作用**：携带用户身份信息，实现无状态认证和接口授权

* **解决的问题**：

  * 解决分布式系统 Session 共享问题
  * 实现微服务之间用户身份传递
  * 降低服务端登录状态存储压力

* **基本原理**：
  `用户登录 → 服务端生成 Token → 客户端携带 JWT → 服务端验证 Signature → 解析 Claims`

  JWT结构：

  ```
  Header.Payload.Signature
  ```

  * Header：声明 Token 类型和签名算法 `{
"alg":"HS256",
"typ":"JWT"
}`
  * Payload：存储用户声明（Claims）`{
 "user_id":1001,
 "role":"admin",
 "exp":1720000000
}`
  * Signature：验证 Token 是否被篡改 `new_Signature(Header+Payload+公钥)==Signature`

  签名：服务端
  * alg：签名方式`对称加密`、`非对称加密`、`none不加密不用或慎用（会造成算法攻击，避免算法攻击需要校验指定类型的算法）`
  * 非对称加密服务端使用私钥加签，客户端使用公钥验签，
  * 对称加密只在服务端加签验签
  * 私钥（secret）暴露风险：伪造攻击、权限提升攻击

  ```
  new_Signature = HMAC/RSA/ECDSA(Header + Payload)
  x
  ```

* **API**：

  * Go：`github.com/golang-jwt/jwt/v5`
  * 创建：`jwt.NewWithClaims()`
  * 签名：`token.SignedString()`
  * 验证：`jwt.Parse()`
  * 常用算法：

    * `HS256`（HMAC 对称签名）
    * `RS256`（RSA 非对称签名）
    * `ES256`（ECDSA 非对称签名）

* **适用场景**：

  * 用户登录认证
  * 微服务 API 鉴权
  * 单点登录（SSO）
  * 第三方开放平台授权

✅ **优点**：

* 无状态认证，不依赖服务端 Session
* 适合分布式和微服务架构
* 跨语言标准，生态成熟

⚠️ **风险**：

* Token 泄露后可直接冒用身份
* Payload 默认 Base64 编码，不是加密，敏感数据不可存储
* Secret 泄露可能导致 Token 伪造
* 长生命周期 Token 存在重放攻击风险

❗ **限制**：

* 无法像 Session 一样主动注销
* Token 体积较大，请求开销高于 Session ID
* 密钥管理和算法配置要求较高
* 不适合存储大量用户状态数据

🔁 **替代**：

* Session + Redis（高控制性场景）
* Opaque Token（服务端可撤销 Token）
* PASETO（JWT 安全简化方案）

