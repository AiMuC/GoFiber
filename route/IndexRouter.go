package route

import (
	"github.com/aimuc/gofiber/app/controller"
	"github.com/aimuc/gofiber/app/middleware"
	"github.com/gofiber/fiber/v2"
)

type IndexRouter struct{}

func (r *IndexRouter) IndexRouters(server *fiber.App) {
	server.Get("/index", middleware.Test, controller.Controller.IndexController.Index) //HelloWorld (包含Test中间件)
	server.Get("/json", controller.Controller.IndexController.Json)                    //输出Json
	server.Get("/view", controller.Controller.IndexController.View)                    //输出模板
	server.Get("/sql", controller.Controller.IndexController.Sql)                      //数据库响应Demo
	server.Get("/MyRes", controller.Controller.IndexController.MyRes)                  //请求Demo
	server.Post("/MyReq", controller.Controller.IndexController.MyReq)                 //响应Demo
}
