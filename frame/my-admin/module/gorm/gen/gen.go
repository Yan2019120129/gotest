package gen

import (
	"gorm.io/gen"
)

var GenDb *gen.Generator

// init 初始化gen
func init() {
	GenDb = gen.NewGenerator(gen.Config{
		//  设置输出路径
		OutPath: configs.GetGen().OutPath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 选择生成模式
	})
}
