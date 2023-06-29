package config

import (
	"fmt"
	"github.com/aimuc/gofiber/support"
)

func DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		support.Env("DB.USERNAME", "root").(string),
		support.Env("DB.PASSWORD", "root").(string),
		support.Env("DB.HOSTNAME", "localhost").(string),
		support.Env("DB.PORT", 3306).(int),
		support.Env("DB.DATABASE", "Fiber").(string),
	)
}
