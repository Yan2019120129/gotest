package appDeal

// Page 页面
type Page struct {
	url  string
	Body string
}

// View 视图
type View struct {
}

// Body 主体
type Body struct {
}

// Form 表单
type Form struct {
	Field string
}

// Table 表格
type Table struct {
	searchField []string
	tools       any
	operate     string
}
