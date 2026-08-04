package models

import "time"

// User is the account used by the JWT login endpoint. Role must be admin,
// operate or user; Casbin policies grant the role its HTTP permissions.
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:64;uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"size:100;not null" json:"-"`
	Role      string    `gorm:"size:20;not null;index" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string { return "casbin_demo_users" }

// Test is the CRUD resource protected by Casbin in this example.
type Test struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Content   string    `gorm:"type:text" json:"content"`
	Owner     string    `gorm:"size:64;not null" json:"owner"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Test) TableName() string { return "casbin_demo_tests" }

// CasbinRule mirrors casbin_rule, the policy table read by gorm-adapter.
// Policies are deliberately seeded only by db.sql, never during application
// startup, so deployment policy changes remain data-driven.
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Ptype string `gorm:"column:ptype;size:100;not null;uniqueIndex:idx_casbin_rule" json:"ptype"`
	V0    string `gorm:"column:v0;size:100;not null;default:'';uniqueIndex:idx_casbin_rule" json:"v0"`
	V1    string `gorm:"column:v1;size:100;not null;default:'';uniqueIndex:idx_casbin_rule" json:"v1"`
	V2    string `gorm:"column:v2;size:100;not null;default:'';uniqueIndex:idx_casbin_rule" json:"v2"`
	V3    string `gorm:"column:v3;size:100;not null;default:'';uniqueIndex:idx_casbin_rule" json:"v3"`
	V4    string `gorm:"column:v4;size:100;not null;default:'';uniqueIndex:idx_casbin_rule" json:"v4"`
	V5    string `gorm:"column:v5;size:100;not null;default:'';uniqueIndex:idx_casbin_rule" json:"v5"`
}

func (CasbinRule) TableName() string { return "casbin_rule" }
