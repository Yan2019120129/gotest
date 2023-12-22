package models

// AdminAuthChildAttrs 属性
type AdminAuthChildAttrs struct {
	Parent string `gorm:"type:varchar(50) not null;comment:父级"`  //	父级
	Child  string `json:"type:varchar(50) not null;comment:子集"`  //	子级
	Type   int    `json:"type:int unsigned not null;comment:类型"` //	类型
}
