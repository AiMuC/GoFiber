package route

import (
	"github.com/aimuc/gofiber/app"
	"github.com/gofiber/fiber/v2"
)

func Routers(server *fiber.App) {
	api := server.Group("/api")                       //路由分组
	v1 := api.Group("/v1")                            //v1 Api
	v1.Get("/hello/:name", func(c *fiber.Ctx) error { //示例路由 Url:http://localhost:8787/api/v1/hello/world
		app.ResponseSuccess(app.SUCCESS, "hello:"+c.Params("name"), nil, c)
		return nil
	})
	//引入IndexRoutes
	Router.IndexRouter.IndexRouters(server) //导入Index路由
}
