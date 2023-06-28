package app

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
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
func ResponseFail(msg string, c *fiber.Ctx) {
	err := c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"code": ERROR,
		"msg":  msg,
	})
	if err != nil {
		panic("响应失败:" + err.Error())
	}
}
