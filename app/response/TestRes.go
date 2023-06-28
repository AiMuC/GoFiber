package response

type TestRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ListRes struct {
	List []*TestRes `json:"list"`
}
