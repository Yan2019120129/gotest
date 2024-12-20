package model

import (
	"go.uber.org/zap"
	"my-frame/module/gorm/orm"
	"my-frame/module/logs"
)

const (
	logMsg = "model"
)

type table struct {
	name  string // 表名
	model any    // 模型
	data  any    // 初始化数据
}

// 初始化数据 initData
func (t *table) initData(data any) *table {
	t.model = data
	return t
}

var TableManage = &Manage{}

// Manage 表管理器
type Manage struct {
	tables []*table
}

// addTable 添加表信息
func (m *Manage) addTable(model any, name string) *table {
	t := &table{model: model, name: name}
	m.tables = append(m.tables, t)
	return t
}

// Init 初始化
func (m *Manage) Init() {
	for _, t := range m.tables {
		err := orm.DB.AutoMigrate(t.model)
		if err != nil {
			logs.Logger.Error(logMsg, zap.Error(err))
		}
	}
}
