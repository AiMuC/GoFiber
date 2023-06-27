package controller

import (
	"github.com/aimuc/gofiber/app"
	"github.com/gofiber/fiber/v2"
)

type IndexController struct {
}

func (a *IndexController) Index(c *fiber.Ctx) error {
	app.ResponseSuccess(app.SUCCESS, "Hello GoFiber", nil, c)
	return nil
}
