# GoFiber

## 介绍

基于Golang Fiber框架二次封装的框架

框架使用文档 [Doc](https://docs.gofiber.io/)

本项目写法参考[webman](https://www.workerman.net/doc/webman/) [thinkphp](https://www.thinkphp.cn/)

使用Diango Templates作为模板引擎 [Doc](https://docs.djangoproject.com/en/dev/topics/templates/)

Orm框架使用Gorm [Doc](https://gorm.io/zh_CN/docs/)

Redis文档参考 [Doc](https://redis.uptrace.dev/zh/)

validator验证器文档参考[Doc](https://github.com/go-playground/validator/tree/master)

env配置 删除根目录下 .example 例:.env 即可

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