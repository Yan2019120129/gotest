package models

const (
	LocalesRedisName                = "_locales"
	LangDictionaryTypeApiTip        = 1  //	接口提示
	LangDictionaryTypeDataTranslate = 2  //	数据翻译
	LangDictionaryTypeHomeTranslate = 10 //	前台翻译

)

// LangDictionary 数据库模型属性
type LangDictionary struct {
	Id        int    `json:"id"`         //主键
	AdminId   int    `json:"admin_id"`   //管理员ID
	Type      int    `json:"type"`       //类型 1接口提示 2前台翻译 3数据翻译
	Alias     string `json:"alias"`      //别名
	Name      string `json:"name"`       //名称
	Field     string `json:"field"`      //键
	Value     string `json:"value"`      //值
	Data      string `json:"data"`       //数据
	CreatedAt int    `json:"created_at"` //创建时间
}
