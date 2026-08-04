package main

// PBKDF2 基于迭代次数的密码派生函数（Password-Based Key Derivation Function 2）
//
// 特点：通过重复迭代散列函数（如 HMAC-SHA256）增加计算成本，
// 标准成熟、适用面广，但本身不消耗内存，GPU 并行破解成本较低。
//
// 实现内容（待填写）：
//  1. 使用 crypto/rand 生成随机盐；
//  2. 调用 golang.org/x/crypto/pbkdf2.Key 派生密钥，
//     迭代次数建议不低于 600000，keyLen=32，散列函数选 SHA-256；
//  3. 按 $pbkdf2-sha256$i=600000$<base64盐>$<base64哈希> 格式编码存储；
//  4. 校验时解析迭代次数与盐，重新计算后做常数时间比较。

func main() {
	// TODO: 待实现 PBKDF2 的密钥派生与校验逻辑
}
