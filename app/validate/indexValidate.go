package validate

import (
	"errors"
	"fmt"
	"github.com/aimuc/gofiber/utils"
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

// AutoValidate 使用自定义错误规则
func (a *IndexValidate) AutoValidate(req any) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			v, ok := IndexValidateMessage[e.Field()][e.ActualTag()]
			if !ok {
				return fmt.Errorf("输入规则为:%s %s %s 请检查您的提交参数后重试", e.Field(), e.ActualTag(), e.Param())
			}
			return errors.New(v)
		}
	}
	return nil
}

// AutoTransValidate 根据alise中的名称使用中文翻译自动输出
func (a *IndexValidate) AutoTransValidate(req any) error {
	validate := validator.New()
	trans := utils.ValidatorTrainInit(validate)
	if err := validate.Struct(req); err != nil {
		for _, validateErr := range err.(validator.ValidationErrors) {
			return errors.New(validateErr.Translate(trans))
		}
	}
	return nil
}
