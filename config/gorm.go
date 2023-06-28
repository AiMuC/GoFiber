package config

import (
	"fmt"
	"os"
)

func DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB.USERNAME"),
		os.Getenv("DB.PASSWORD"),
		os.Getenv("DB.HOSTNAME"),
		os.Getenv("DB.PORT"),
		os.Getenv("DB.DATABASE"),
	)
}
