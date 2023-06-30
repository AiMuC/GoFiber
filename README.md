# GoFiber

## 介绍 👋 Welcome

基于Golang Fiber框架二次封装的框架

框架使用文档 [gofiber](https://docs.gofiber.io/)

本项目写法参考[webman](https://www.workerman.net/doc/webman/) [thinkphp](https://www.thinkphp.cn/)

- 环境配置 .example.env 修改为 .env

## 技术栈

| 名称        | 说明         | Doc                                                     |
|-----------|------------|---------------------------------------------------------|
| Gorm      | 数据库ORM     | https://gorm.io/zh_CN/docs/                             |
| GoDotEnv  | Env配置文件    | https://github.com/joho/godotenv                        |
| validator | 验证器        | https://github.com/go-playground/validator/tree/master  |
| Zap       | 日志库        | https://pkg.go.dev/go.uber.org/zap                      |
| Redis     | go redis驱动 | https://redis.uptrace.dev/zh/                           |
| Templates | Diango模板引擎 | https://docs.djangoproject.com/en/dev/topics/templates/ |

## 使用说明

- 使用 go mod 并安装go依赖包 运行命令 `go mod tidy` 相当于PHP中的 `composer install`

- 编译 `go build -o gofiber.exe start.go`

- 运行 `go run .`

## 初始化环境

```shell
# 如果您使用的 Go 版本是 1.13 - 1.15 需要手动设置GO111MODULE=on, 开启方式如下命令, 如果你的 Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
# 步骤三、同步依赖
go mod tidy
# 步骤四、启动项目
go run .
```

# 目录结构

- 多应用目录结构请参考 [webman多应用](https://www.workerman.net/doc/webman/multiapp.html)

| 目录         | 说明                         |
|------------|----------------------------|
| global     | 全局变量定义 如：orm, 日志, cache等   |
| initialize | 全局变量初始化（组件初始化）如：redis，orm等 |
| support    | 框架封装函数                     |
| utils      | 常用工具封装函数                   |
| runtime    | 运行时目录                      |