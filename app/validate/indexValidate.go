package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type IndexValidate struct {
}

var IndexValidateMessage = map[string]map[string]string{
	"Id": {
		"required": "ID为必填项！",
		"max":      "ID最大值为1",
	},
	"Name": {
		"required": "文件夹名称为必填项！",
	},
}

func (a *IndexValidate) AutoValidate(req any) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			v, ok := IndexValidateMessage[e.Field()][e.ActualTag()]
			if !ok {
				panic(fmt.Errorf("输入规则为:%s %s %s 请检查您的提交参数后重试", e.Field(), e.ActualTag(), e.Param()))
			}
			panic(v)
		}
	}
	return nil
}
