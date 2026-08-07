package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// PBKDF2 基于迭代次数的密码派生函数（Password-Based Key Derivation Function 2）
//
// 特点：通过重复迭代散列函数（如 HMAC-SHA256）增加计算成本，
// 标准成熟、适用面广，但本身不消耗内存，GPU 并行破解成本较低。
//
// 实现内容：
//  1. 使用 crypto/rand 生成随机盐；
//  2. 调用 golang.org/x/crypto/pbkdf2.Key 派生密钥，
//     迭代次数建议不低于 600000，keyLen=32，散列函数选 SHA-256；
//  3. 按 $pbkdf2-sha256$i=600000$<base64盐>$<base64哈希> 格式编码存储；
//  4. 校验时解析迭代次数与盐，重新计算后做常数时间比较。

// Pbkdf2Params PBKDF2 计算参数
type Pbkdf2Params struct {
	Iterations int // 迭代次数，越大单次计算越慢、越安全
	KeyLen     int // 派生密钥长度（字节）
	SaltLen    int // 盐长度（字节）
}

// PasswordManager 密码管理对象，封装 PBKDF2 的密钥派生与校验
type PasswordManager struct {
	Params Pbkdf2Params
}

// DefaultPbkdf2Params 返回推荐的 PBKDF2 参数：迭代 600000 次，keyLen=32
func DefaultPbkdf2Params() Pbkdf2Params {
	return Pbkdf2Params{
		Iterations: 600000,
		KeyLen:     32,
		SaltLen:    16,
	}
}

// NewPasswordManager 创建使用默认参数的密码管理器
func NewPasswordManager() *PasswordManager {
	return &PasswordManager{Params: DefaultPbkdf2Params()}
}

// Hash 使用 PBKDF2 生成密码哈希：随机盐 + 派生密钥，并按标准格式编码
func (m *PasswordManager) Hash(password string) (string, error) {
	// 1. 使用 crypto/rand 生成随机盐，保证相同密码每次哈希结果不同
	salt := make([]byte, m.Params.SaltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("生成随机盐失败: %w", err)
	}

	// 2. 调用 pbkdf2.Key 派生密钥，HMAC-SHA256 迭代计算，盐参与防预计算
	key := pbkdf2.Key([]byte(password), salt, m.Params.Iterations, m.Params.KeyLen, sha256.New)

	// 3. 按 $pbkdf2-sha256$i=600000$<base64盐>$<base64哈希> 格式编码
	return fmt.Sprintf("$pbkdf2-sha256$i=%d$%s$%s",
		m.Params.Iterations,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key),
	), nil
}

// Verify 校验明文密码与已存储哈希是否匹配：解析参数与盐后重算，再做常数时间比较
func (m *PasswordManager) Verify(hash, password string) (bool, error) {
	// 期望格式：["", "pbkdf2-sha256", "i=600000", "盐", "哈希"]
	parts := strings.Split(hash, "$")
	if len(parts) != 5 || parts[1] != "pbkdf2-sha256" {
		return false, errors.New("非法的 pbkdf2 哈希格式")
	}

	// 解析 i=.. 迭代次数，后续按相同参数重算才能得到一致结果
	var iterations int
	if _, err := fmt.Sscanf(parts[2], "i=%d", &iterations); err != nil {
		return false, fmt.Errorf("解析 pbkdf2 参数失败: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[3])
	if err != nil {
		return false, fmt.Errorf("解码盐失败: %w", err)
	}
	expected, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("解码哈希失败: %w", err)
	}

	// 用相同迭代次数和盐重新派生，再与存储值做常数时间比较
	actual := pbkdf2.Key([]byte(password), salt, iterations, len(expected), sha256.New)
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
