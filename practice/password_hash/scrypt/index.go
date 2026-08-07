package main

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

// scrypt 内存困难型密码派生函数（Memory-Hard Password-Based Key Derivation Function）
//
// 特点：由 Colin Percival 设计，通过消耗大量内存增加破解成本，
// 适合在资源受限场景下替代 PBKDF2，曾用于加密文件系统等场景。
//
// 实现内容：
//  1. 使用 crypto/rand 生成随机盐；
//  2. 调用 golang.org/x/crypto/scrypt.Key 派生密钥，
//     参数 N=32768, r=8, p=1，keyLen=32；
//  3. 按 $scrypt$N=32768,r=8,p=1$<base64盐>$<base64哈希> 格式编码存储；
//  4. 校验时解析参数与盐，重新计算后做常数时间比较。

// ScryptParams scrypt 计算参数
type ScryptParams struct {
	N       int // CPU/内存成本参数，必须是 2 的幂
	R       int // 块大小参数
	P       int // 并行度参数
	KeyLen  int // 派生密钥长度（字节）
	SaltLen int // 盐长度（字节）
}

// PasswordManager 密码管理对象，封装 scrypt 的密钥派生与校验
type PasswordManager struct {
	Params ScryptParams
}

// DefaultScryptParams 返回推荐的 scrypt 参数：N=32768, r=8, p=1
func DefaultScryptParams() ScryptParams {
	return ScryptParams{
		N:       32768,
		R:       8,
		P:       1,
		KeyLen:  32,
		SaltLen: 16,
	}
}

// NewPasswordManager 创建使用默认参数的密码管理器
func NewPasswordManager() *PasswordManager {
	return &PasswordManager{Params: DefaultScryptParams()}
}

// Hash 使用 scrypt 生成密码哈希：随机盐 + 派生密钥，并按标准格式编码
func (m *PasswordManager) Hash(password string) (string, error) {
	// 1. 使用 crypto/rand 生成随机盐，保证相同密码每次哈希结果不同
	salt := make([]byte, m.Params.SaltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("生成随机盐失败: %w", err)
	}

	// 2. 调用 scrypt.Key 派生密钥，盐参与计算，防止预计算攻击
	key, err := scrypt.Key([]byte(password), salt, m.Params.N, m.Params.R, m.Params.P, m.Params.KeyLen)
	if err != nil {
		return "", fmt.Errorf("派生密钥失败: %w", err)
	}

	// 3. 按 $scrypt$N=32768,r=8,p=1$<base64盐>$<base64哈希> 格式编码
	return fmt.Sprintf("$scrypt$N=%d,r=%d,p=%d$%s$%s",
		m.Params.N, m.Params.R, m.Params.P,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key),
	), nil
}

// Verify 校验明文密码与已存储哈希是否匹配：解析参数与盐后重算，再做常数时间比较
func (m *PasswordManager) Verify(hash, password string) (bool, error) {
	// 期望格式：["", "scrypt", "N=32768,r=8,p=1", "盐", "哈希"]
	parts := strings.Split(hash, "$")
	if len(parts) != 5 || parts[1] != "scrypt" {
		return false, errors.New("非法的 scrypt 哈希格式")
	}

	// 解析 N=..,r=..,p=.. 参数，后续按相同参数重算才能得到一致结果
	var n, r, p int
	if _, err := fmt.Sscanf(parts[2], "N=%d,r=%d,p=%d", &n, &r, &p); err != nil {
		return false, fmt.Errorf("解析 scrypt 参数失败: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[3])
	if err != nil {
		return false, fmt.Errorf("解码盐失败: %w", err)
	}
	expected, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("解码哈希失败: %w", err)
	}

	// 用相同参数和盐重新派生，再与存储值做常数时间比较
	actual, err := scrypt.Key([]byte(password), salt, n, r, p, len(expected))
	if err != nil {
		return false, fmt.Errorf("重新派生密钥失败: %w", err)
	}
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
