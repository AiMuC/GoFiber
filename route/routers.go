package route

import (
	"github.com/aimuc/gofiber/app"
	"github.com/aimuc/gofiber/app/controller"
	"github.com/gofiber/fiber/v2"
)

func Routers(server *fiber.App) {
	server.Get("/", (&controller.IndexController{}).Index)
	api := server.Group("/api")                       //路由分组
	v1 := api.Group("/v1")                            //v1 Api
	v1.Get("/hello/:name", func(c *fiber.Ctx) error { //示例路由
		app.ResponseSuccess(app.SUCCESS, "hello:"+c.Params("name"), nil, c)
		return nil
	})
}
