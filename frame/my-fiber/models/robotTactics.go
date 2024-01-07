package models

const (
	RobotTacticsStatusDelete   = -2 //	删除
	RobotTacticsStatusActivate = 10 //	运行状态
	RobotTacticsStatusComplete = 20 //	完成
)

// RobotTactics 数据库模型属性
type RobotTactics struct {
	Id        int64   `json:"id"`         //主键
	AdminId   int64   `json:"admin_id"`   //管理员ID
	UserId    int64   `json:"user_id"`    //用户ID
	ProductId int64   `json:"product_id"` //产品ID
	Type      int64   `json:"type"`       //类型 1期货机器人
	Money     float64 `json:"money"`      //投资金额
	Nums      int64   `json:"nums"`       //当前期数
	Fee       float64 `json:"fee"`        //手续费金额
	Earnings  float64 `json:"earnings"`   //收益金额
	Data      string  `json:"data"`       //数据
	Status    int64   `json:"status"`     //状态 -2删除 -1暂停 10启用 20完成
	ExpiredAt int64   `json:"expired_at"` //过期时间
	UpdatedAt int64   `json:"updated_at"` //更新时间
	CreatedAt int64   `json:"created_at"` //创建时间
}
