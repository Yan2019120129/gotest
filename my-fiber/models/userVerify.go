package models

const (
	UserVerifyTypeIdCard     = 1
	UserVerifyStatusPending  = 10
	UserVerifyStatusComplete = 20
)

// UserVerify 数据库模型属性
type UserVerify struct {
	Id        int    `json:"id"`         //主键
	AdminId   int    `json:"admin_id"`   //管理员ID
	UserId    int    `json:"user_id"`    //用户ID
	Type      int    `json:"type"`       //类型 1身份证 2护照
	RealName  string `json:"real_name"`  //真实姓名
	IdNumber  string `json:"id_number"`  //证件号码
	IdPhoto1  string `json:"id_photo1"`  //证件照1
	IdPhoto2  string `json:"id_photo2"`  //证件照2
	IdPhoto3  string `json:"id_photo3"`  //证件照3
	Address   string `json:"address"`    //地址
	Data      string `json:"data"`       //数据
	Status    int    `json:"status"`     //状态 -1拒绝｜10审核｜20通过
	CreatedAt int    `json:"created_at"` //创建时间
	UpdatedAt int64  `json:"updated_at"` //更新时间
}
