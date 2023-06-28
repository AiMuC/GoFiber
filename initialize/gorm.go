package initialize

import (
	"fmt"
	"github.com/aimuc/gofiber/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func GormMysql() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DSN(),
		DefaultStringSize:         191,  // 字符串的默认长度
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		SkipInitializeWithVersion: true, // // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   os.Getenv("DB.PREFIX"), // 表名前缀，`User` 的表名应该是 `prefix_users`
			SingularTable: true,                   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `prefix_user`
		},
		SkipDefaultTransaction: true, // 跳过默认事务
		Logger: logger.New( // 日志
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic(fmt.Errorf("链接数据库出现错误:%s", err))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConn 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConn 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。
	return db
}
