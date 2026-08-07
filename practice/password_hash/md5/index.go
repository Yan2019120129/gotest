package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// MD5 通用哈希函数（Message-Digest Algorithm 5），不能用于密码存储
//
// 特点：输出固定 128 位（16 字节、32 个十六进制字符），
// 早已被证明存在碰撞攻击，可伪造相同摘要，严禁用于密码存储与安全校验。
// 单向不可逆：只能由原文算出摘要，无法从摘要反推原文；
// 现实中所谓“破解”只能靠穷举/字典/彩虹表比对，能否成功取决于原文强度。
//
// 实现内容：
//  1. 调用 crypto/md5.Sum 计算 16 字节摘要；
//  2. 用 hex 编码输出 32 个十六进制字符；
//  3. 在注释中说明：仅可用于非安全场景（如校验非关键数据的完整性），
//     密码存储必须改用 bcrypt / Argon2id 等慢速加盐算法。

func main() {
	// 待散列的原始数据
	data := []byte("hello world")

	// 方式一：一次性接口，适合小数据，直接得到 [16]byte 数组
	sum := md5.Sum(data)
	fmt.Printf("方式一  hex：%x\n", sum)

	// 方式二：流式接口，适合大数据或分块写入（如文件校验）
	h := md5.New()
	h.Write(data)
	digest := h.Sum(nil)
	fmt.Printf("方式二  hex：%s\n", hex.EncodeToString(digest))
	fmt.Printf("方式二  base64：%s\n", base64.StdEncoding.EncodeToString(digest))

	// 解析输出：32 个十六进制字符 = 16 字节 = 128 位
	fmt.Printf("摘要长度：%d 字节（%d 位），hex 字符数：%d\n", len(digest), len(digest)*8, len(hex.EncodeToString(digest)))

	// 演示雪崩效应：仅改动一个字符，摘要完全不同
	slightly := md5.Sum([]byte("hello worle"))
	fmt.Printf("改动后的摘要：%x\n", slightly)

	// 安全说明：MD5 仅可用于非安全场景（如校验非关键数据的完整性），
	// 存在碰撞攻击，严禁用于密码存储；密码存储必须改用 bcrypt / Argon2id 等慢速加盐算法。
	fmt.Println("安全提示：MD5 存在碰撞攻击，严禁用于密码存储与安全校验，密码存储应使用 bcrypt / Argon2id。")

	// MD5 可逆性演示：MD5 不可逆，不存在从摘要还原原文的逆算法，
	// 唯一可行的“还原”方式是字典/穷举比对：对候选原文逐一计算摘要再比较。
	target := md5.Sum([]byte("Aa123098..")) // 假设泄露的摘要来自该密码
	fmt.Printf("待破解的目标摘要：%x\n", target)

	// 弱密码字典：真实攻击会使用更大字典与彩虹表，这里仅演示原理
	weakDict := []string{"123456", "password", "Aa123098..", "qwerty"}
	found := ""
	for _, c := range weakDict {
		if md5.Sum([]byte(c)) == target {
			found = c
			break
		}
	}
	if found != "" {
		fmt.Printf("字典命中：原文被“还原”为 %q（本质是穷举比对，MD5 本身不可逆）\n", found)
	} else {
		fmt.Println("字典未命中：摘要无法直接反推原文，只能继续穷举或扩大字典")
	}

	// 强密码演示：若原文不在任何字典中，摘要无法在合理时间内被还原
	strong := md5.Sum([]byte("xK9#mQ2!pLv5@zT8"))
	fmt.Printf("强密码摘要：%x（字典外，无法从摘要反推原文，只能靠暴力穷举）\n", strong)
}
