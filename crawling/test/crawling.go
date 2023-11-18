package test

type crawlingParam struct {
	Id   int64  `json:"id" validate:"required,gt=0"`
	Urls string `json:"urls" validate:"required"`
}

type TestStruct struct {
	name string
	age  int64
}

// NewTestStruct 爬取产品
func NewTestStruct() *TestStruct {
	return &TestStruct{
		age:  18,
		name: "yan",
	}
}

func (*TestStruct) GetName(name string) string {

	return name
}
