package initialize

import (
	"github.com/aimuc/gofiber/app"
	"github.com/aimuc/gofiber/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"os"
)

func RunSever() {
	initGlobal()
	initServer(os.Getenv("SERVER.PORT"))
}

func initGlobal() {
}

func initServer(addr string) {
	server := fiber.New(fiber.Config{
		Prefork:       true,      //开启多进程
		CaseSensitive: true,      //区分大小写
		StrictRouting: true,      //严格路由模式
		ServerHeader:  "GoFiber", //Response Server Name
		ErrorHandler: func(c *fiber.Ctx, err error) error { //全局错误处理
			return c.Status(app.ERROR).JSON(fiber.Map{
				"code": app.ERROR,
				"msg":  err.Error(),
			})
		},
	})
	server.Use(recover.New()) //开启全局异常捕获
	route.Routers(server)
	log.Fatal(server.Listen(addr))
}
