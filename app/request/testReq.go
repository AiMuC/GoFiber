package request

type TestReq struct {
	Id   int    `json:"id" validate:"required,max=1" alise:"测试ID"`
	Name string `json:"name" validate:"required" alise:"测试名称"`
	Test []int  `json:"test" validate:"required" alise:"测试列表"`
}
