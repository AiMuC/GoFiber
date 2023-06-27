package app

import (
	"github.com/gofiber/fiber/v2"
)

const (
	SUCCESS = 1
	ERROR   = 0
	WARNING = -1
)

func ResponseSuccess(code int, msg string, data any, c *fiber.Ctx) {
	err := c.JSON(fiber.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	if err != nil {
		panic("响应失败:" + err.Error())
	}
}
