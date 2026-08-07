package main

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2id（最低安全级别）
// Memory-内存：19MB
// time-迭代：2
// Threads-并行度：1

// Argon2id（企业安全级别）
// Memory-内存：64MB
// time-迭代次数：3
// Threads-并行度：4
// 标准：RFC 9106（国家级别密码规范）

// Argon2id（金融政务：高安全需求）
// Memory-内存：256MB
// time-迭代：4
// Threads-并行度：8

// Argon2id 内存困难型密码哈希函数（Memory-Hard Password Hash）
//
// 特点：2015 年密码哈希竞赛冠军，需要大量内存参与计算，
// 可同时抵御 GPU 暴力破解与 ASIC 专用硬件攻击，是目前推荐的密码存储方案。
//
// 这种密码比较方式的好处：
//  1. 单向不可逆：Argon2id 是单向散列算法，无法从哈希反推明文，
//     即使数据库泄露，攻击者拿到的也只是无法直接使用的哈希值。
//  2. 自动加盐：每次 Hash 都会用 crypto/rand 生成随机盐，
//     相同密码产生的哈希也不同（下方演示两次生成结果不同），可抵御彩虹表攻击。
//  3. 内存困难：计算需要 64 MiB 内存参与，GPU/ASIC 并行破解成本远高于
//     bcrypt、PBKDF2，这是 Argon2id 最大的优势。
//  4. 参数自描述：哈希字符串中自带 v/m/t/p 参数，校验时解析后重算，
//     未来调高参数后旧哈希仍可继续校验，方便平滑升级。
//  5. 常数时间比较：校验时使用 crypto/subtle.ConstantTimeCompare，
//     避免提前返回泄露信息，可抵御时序攻击。

// Argon2idParams Argon2id 计算参数
type Argon2idParams struct {
	Time    uint32 // 迭代次数
	Memory  uint32 // 内存消耗，单位 KiB
	Threads uint8  // 并行度
	KeyLen  uint32 // 派生密钥长度（字节）
	SaltLen uint32 // 盐长度（字节）
}

// PasswordManager 密码管理对象，封装 Argon2id 的哈希生成与校验
type PasswordManager struct {
	Params Argon2idParams
}

// DefaultArgon2idParams 返回 RFC 9106 推荐参数：time=1, memory=64MiB, threads=4
func DefaultArgon2idParams() Argon2idParams {
	return Argon2idParams{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
		SaltLen: 16,
	}
}

// NewPasswordManager 创建使用 Argon2id 默认参数的密码管理器
func NewPasswordManager() *PasswordManager {
	return &PasswordManager{Params: DefaultArgon2idParams()}
}

// Hash 使用 Argon2id 生成密码哈希：随机盐 + 派生密钥，并按标准格式编码
func (m *PasswordManager) Hash(password string) (string, error) {
	// 1. 使用 crypto/rand 生成随机盐，保证相同密码每次哈希结果不同
	salt := make([]byte, m.Params.SaltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("生成随机盐失败: %w", err)
	}

	// 2. 调用 IDKey 派生密钥，盐参与计算，防止预计算攻击
	key := argon2.IDKey([]byte(password), salt, m.Params.Time, m.Params.Memory, m.Params.Threads, m.Params.KeyLen)

	// 3. 按 $argon2id$v=19$m=65536,t=1,p=4$<base64盐>$<base64哈希> 格式编码
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		m.Params.Memory, m.Params.Time, m.Params.Threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key),
	), nil
}

// Verify 校验明文密码与已存储哈希是否匹配：解析参数与盐后重算，再做常数时间比较
func (m *PasswordManager) Verify(hash, password string) (bool, error) {
	// 期望格式：["", "argon2id", "v=19", "m=65536,t=1,p=4", "盐", "哈希"]
	parts := strings.Split(hash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return false, errors.New("非法的 argon2id 哈希格式")
	}
	if !strings.HasPrefix(parts[2], "v=") {
		return false, errors.New("哈希中缺少版本号 v")
	}

	// 解析 m=..,t=..,p=.. 参数，后续按相同参数重算才能得到一致结果
	var memory, time uint32
	var threads uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads); err != nil {
		return false, fmt.Errorf("解析 Argon2id 参数失败: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("解码盐失败: %w", err)
	}
	expected, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("解码哈希失败: %w", err)
	}

	// 用相同参数和盐重新派生，再与存储值做常数时间比较
	actual := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(expected)))
	return subtle.ConstantTimeCompare(actual, expected) == 1, nil
}

func main() {
	m := NewPasswordManager()
	password := "Aa123098.."

	// 1. 生成新哈希并校验：正确密码应匹配
	hash, err := m.Hash(password)
	if err != nil {
		panic(err)
	}
	ok, err := m.Verify(hash, password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("新生成的哈希：%s\n", hash)
	fmt.Printf("正确密码校验结果：%v\n", ok)

	// 2. 错误密码校验：应不匹配
	ok, err = m.Verify(hash, "wrong-password")
	if err != nil {
		panic(err)
	}
	fmt.Printf("错误密码校验结果：%v\n", ok)

	// 3. 同一密码再生成一次：盐不同，哈希也不同（加盐效果演示）
	hash2, err := m.Hash(password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("再次生成的哈希：%s\n", hash2)
	fmt.Printf("两次哈希是否相同：%v\n", hash == hash2)
}
