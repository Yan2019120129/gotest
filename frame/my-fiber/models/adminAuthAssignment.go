package models

// AdminAuthAssignmentsAttrs 属性
type AdminAuthAssignmentsAttrs struct {
	ItemName  string `gorm:"type:varchar(50)）not null;comment:"` //	名称
	UserId    int    `gorm:"type:int not null;comment:用户id"`     //	用户ID
	CreatedAt int    `gorm:"type:int not null;comment:创建时间"`     //	创建时间
}
