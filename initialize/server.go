package initialize

import (
	"fmt"
	"github.com/aimuc/gofiber/app"
	"github.com/aimuc/gofiber/global"
	"github.com/aimuc/gofiber/route"
	"github.com/aimuc/gofiber/support"
	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
	"net/http"
	"time"
)

func RunSever() {
	initialization() //初始化操作(此处配置框架启动前的操作)
	initServer(support.Env("SERVER.PORT", ":8787").(string))
}

func initialization() {
	global.Db = GormMysql()     //初始化Mysql
	global.Redis = RedisDrive() //初始化Redis
	global.Log = Zap()          //初始化Zap
}

func initServer(addr string) {
	engine := django.New("./app/view", ".html")
	server := fiber.New(fiber.Config{
		Prefork:       true,      //开启多进程
		CaseSensitive: true,      //区分大小写
		StrictRouting: true,      //严格路由模式
		ServerHeader:  "GoFiber", //Response Server Name
		Views:         engine,    //使用DianGoTemplates作为模板引擎 Doc:https://docs.djangoproject.com/en/dev/topics/templates/
		ErrorHandler: func(c *fiber.Ctx, err error) error { //全局错误处理
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"code": app.ERROR,
				"msg":  err.Error(),
			})
		},
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	server.Use(recover.New(), fiberzap.New(fiberzap.Config{
		Logger: global.Log,
	})) //开启全局异常捕获,替换日志组件,
	server.Use(limiter.New(limiter.Config{
		Next:       nil,
		Max:        support.Env("SERVER.LIMITER.MAX", 999).(int),
		Expiration: time.Duration(support.Env("SERVER.LIMITER.EXP", 1).(int)) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).SendString("429 to many request!")
		},
	})) //业务限流(如无需使用直接注释整段即可)
	server.Static("/", "./public", fiber.Static{ //开启静态目录代理
		Compress:      true,             //是否开启压缩
		ByteRange:     true,             //是否启用字节范围请求。
		Browse:        true,             //是否启用目录浏览
		Index:         "index.html",     //默认的访问
		CacheDuration: 10 * time.Second, //缓存时间
		MaxAge:        3600,
	})
	route.Routers(server)      //加载路由文件
	err := server.Listen(addr) //启动服务
	if err != nil {
		fmt.Printf("GoFiber 启动失败" + err.Error()) //服务启动失败输出原因
	}
}
