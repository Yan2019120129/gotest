package dto

// UserData 返回的用户数据
type UserData struct {
	AdminId     int     `json:"adminId"`
	ParentId    int     `json:"parentId"`
	UserName    string  `json:"userName"`
	NickName    string  `json:"nickName"`
	Email       string  `json:"email"`
	Telephone   string  `json:"telephone"`
	Avatar      string  `json:"avatar"`
	Sex         int     `json:"sex"`
	Birthday    int     `json:"birthday"`
	Password    string  `json:"password"`
	SecurityKey string  `json:"securityKey"`
	Money       float64 `json:"money"`
	Type        int     `json:"type"`
	Status      int     `json:"status"`
	Data        string  `json:"data"`
}

// IndexData 产品列表返回参数
type IndexData struct {
	Id         int     `json:"id"`
	CategoryId int     `json:"category_id"` // 分类ID
	Images     string  `json:"images"`      // 图片数组
	Name       string  `json:"name"`        // 标题
	Money      float64 `json:"money"`       // 金额
	Type       int     `json:"type"`        // 类型 1默认
	Sales      int     `json:"sales"`       // 销售量
	Nums       int     `json:"nums"`        // 限购 -1无限
	Used       int     `json:"used"`        // 已使用
	Total      int     `json:"total"`       // 总数
	IsCollect  bool    `json:"isCollect"`   // 是否收藏
	Data       string  `json:"data"`        // 数据
}
