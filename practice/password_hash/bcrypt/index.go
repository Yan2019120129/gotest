package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// bcrypt 自适应密码哈希函数（Adaptive Password Hash）
//
// 案例来源：根目录 main.go，主题为 bcrypt 密码校验
//
// 这种密码比较方式的好处：
//  1. 单向不可逆：bcrypt 是单向散列算法，无法从哈希反推明文，
//     即使数据库泄露，攻击者拿到的也只是无法直接使用的哈希值。
//  2. 自动加盐：每次 GenerateFromPassword 都会生成随机盐，
//     相同密码产生的哈希也不同（如下方 hash1 与 hash2），可抵御彩虹表攻击。
//  3. 成本可控：cost 参数控制计算强度，cost 越大单次校验越耗时，
//     可显著拖慢离线暴力破解的速度（默认成本为 10）。
//  4. 常数时间比较：CompareHashAndPassword 内部使用常数时间比较，
//     避免因哈希不匹配时的提前返回而泄露信息，可抵御时序攻击。
//  5. 无需保存明文：数据库只存哈希，校验时用明文重新计算再比对，
//     任何环节都不需要还原明文密码。
//
// 预置的两个 bcrypt 哈希（同一密码 "Aa123098.." 的两份合法哈希，用于演示加盐效果）
var (
	hash1 = "$2a$10$4l1rFD3gAYCBU6IDKEnhPuRXZegZTtqMg8.5AlJLoBSScl4vyTJsy"
	hash2 = "$2a$10$YtQ0zYhMEcNJMuM.hBYpmOw7.L5GXVuyRM3n3uBVONqwyxA3nW4P6"
)

// PasswordManager 密码管理对象，封装 bcrypt 的哈希生成与校验
type PasswordManager struct {
	Cost int // 计算成本，越大越耗时、越安全
}

// NewPasswordManager 创建密码管理器，使用 bcrypt 默认成本
func NewPasswordManager() *PasswordManager {
	return &PasswordManager{Cost: bcrypt.DefaultCost}
}

// Hash 使用 bcrypt 生成密码哈希（内部自动加盐）
func (m *PasswordManager) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), m.Cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Verify 校验明文密码与已有哈希是否匹配
func (m *PasswordManager) Verify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func main() {
	m := NewPasswordManager()
	p := "Aa123098.."

	// 1. 校验预置哈希，验证同一密码的两份哈希都能通过
	for _, h := range []string{hash1, hash2} {
		if err := m.Verify(h, p); err != nil {
			panic(err)
		}
		fmt.Println("校验通过：", h)
	}

	// 2. 现场生成新哈希并校验（取消注释即可体验加盐效果）
	//if hash, err := m.Hash(p); err != nil {
	//	panic(err)
	//} else if err = m.Verify(hash, p); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("新生成的哈希：", hash)
	//}
}
