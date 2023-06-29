package config

import (
	"fmt"
	"github.com/aimuc/gofiber/utils"
)

func DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.Env("DB.USERNAME", "root").(string),
		utils.Env("DB.PASSWORD", "root").(string),
		utils.Env("DB.HOSTNAME", "localhost").(string),
		utils.Env("DB.PORT", 3306).(int),
		utils.Env("DB.DATABASE", "Fiber").(string),
	)
}
