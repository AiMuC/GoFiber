package controller

import (
	"fmt"
	"github.com/aimuc/gofiber/app"
	"github.com/aimuc/gofiber/app/model"
	"github.com/aimuc/gofiber/app/request"
	"github.com/aimuc/gofiber/app/response"
	"github.com/aimuc/gofiber/app/validate"
	"github.com/aimuc/gofiber/global"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type IndexController struct {
}

func (a *IndexController) Index(c *fiber.Ctx) error {
	fmt.Println(c.Locals("Test")) //输出中间件设置的变量
	return c.Status(200).SendString("Hello, World!")
}

func (a *IndexController) Json(c *fiber.Ctx) error {
	app.ResponseSuccess(app.SUCCESS, "Hello GoFiber", nil, c)
	return nil
}

func (a *IndexController) View(c *fiber.Ctx) error {
	return c.Render("index/index", fiber.Map{
		"Body": "Hello, World!",
	})
}

func (a *IndexController) MyReq(c *fiber.Ctx) error {
	var req request.TestReq
	if err := c.BodyParser(&req); err != nil {
		panic("参数绑定失败,请检查输入参数是否正确")
	}
	if err := (&validate.IndexValidate{}).AutoTransValidate(&req); err != nil { //验证
		panic(err)
	}
	return c.JSON(req)
}

func (a *IndexController) MyRes(c *fiber.Ctx) error {
	res := &response.ListRes{
		List: []*response.TestRes{},
	}
	for i := 0; i < 5; i++ {
		tmp := &response.TestRes{
			Id:   i,
			Name: "测试" + strconv.Itoa(i),
		}
		res.List = append(res.List, tmp)
	}
	app.ResponseSuccess(app.SUCCESS, "列表获取成功", res, c)
	return nil
}
func (a *IndexController) Sql(c *fiber.Ctx) error {
	var users []model.Users
	global.Db.Model(model.Users{}).Find(&users)
	return c.JSON(users)
}
