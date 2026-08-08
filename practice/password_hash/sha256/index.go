package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// SHA-256 通用哈希函数，不适合直接存密码
//
// 解析 SHA-256：
// 1. 固定长度输出：无论输入多长，输出恒为 256 位 = 32 字节 = 64 个十六进制字符。
// 2. 单向性：只能由输入算出输出，无法从输出反推输入，常用于完整性校验与数字签名。
// 3. 雪崩效应：输入哪怕只改动一个比特，输出的 256 位中约一半会发生变化。
// 4. 无密钥：SHA-256 是散列算法而不是加密算法，不依赖密钥，任何人都能计算。
// 5. 适用场景：文件校验、消息摘要、Merkle 树、HMAC 底层等。
// 6. 注意：SHA-256 速度快且无盐，不适合直接存储用户密码，密码应使用 bcrypt 等慢速加盐算法。

func main() {
	// 待散列的原始数据
	data := []byte("hello world")

	// 方式一：一次性接口，适合小数据，直接得到 [32]byte 数组
	sum := sha256.Sum256(data)
	fmt.Printf("方式一  hex：%x\n", sum)

	// 方式二：流式接口，适合大数据或分块写入（如文件校验）
	h := sha256.New()
	h.Write(data)
	digest := h.Sum(nil)
	fmt.Printf("方式二  hex：%s\n", hex.EncodeToString(digest))
	fmt.Printf("方式二  base64：%s\n", base64.StdEncoding.EncodeToString(digest))

	// 解析输出：64 个十六进制字符 = 32 字节 = 256 位
	fmt.Printf("摘要长度：%d 字节（%d 位），hex 字符数：%d\n", len(digest), len(digest)*8, len(hex.EncodeToString(digest)))

	// 演示雪崩效应：仅改动一个字符，摘要完全不同
	slightly := sha256.Sum256([]byte("hello worle"))
	fmt.Printf("改动后的摘要：%x\n", slightly)
	fmt.Printf("修改前后摘要是否相同：%t\n", sum == slightly)
}
