package route

import (
	"context"
	"github.com/aimuc/gofiber/app"
	"github.com/aimuc/gofiber/global"
	"github.com/gofiber/fiber/v2"
)

func Routers(server *fiber.App) {
	api := server.Group("/api")                       //api 一级路由分组
	v1 := api.Group("/v1")                            //api/v1 二级路由分组
	v2 := api.Group("/v2")                            //api/v1 二级路由分组
	v1.Get("/hello/:name", func(c *fiber.Ctx) error { //示例路由 Url:http://localhost:8787/api/v1/hello/world
		app.ResponseSuccess(app.SUCCESS, "hello:"+c.Params("name"), nil, c)
		return nil
	})
	v1.Get("/error", func(c *fiber.Ctx) error { //示例路由 Url:http://localhost:8787/api/v1/error
		app.ResponseFail("这是一个响应错误示例", c) //自定义函数方式
		////(由于定义了全局异常捕获 以下两种方式抛出的异常也会被响应到客户端中)
		//panic("这是一个响应错误示例")             //抛出异常方式
		//return errors.New("这是一个响应错误示例") //返回错误方式
		return nil
	})
	server.Get("/redis", func(c *fiber.Ctx) error {
		ctx := context.Background()
		global.Redis.Set(ctx, "name", "gofiber", 0)
		res, _ := global.Redis.Get(ctx, "name").Result()
		return c.JSON(res)
	})
	Router.IndexRouter.IndexRouters(server) //导入Index路由在/目录下  	Url:http://localhost:8787/json
	Router.IndexRouter.IndexRouters(v1)     //导入到v1分组  			Url:http://localhost:8787/api/v1/json
	Router.IndexRouter.IndexRouters(v2)     //导入到v2分组			Url:http://localhost:8787/api/v2/json

}
