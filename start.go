package main

import (
	"github.com/aimuc/gofiber/initialize"
	_ "github.com/joho/godotenv/autoload" //加载env配置
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	initialize.RunSever()
}
