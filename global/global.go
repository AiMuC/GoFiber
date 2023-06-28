package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Db    *gorm.DB
	Redis *redis.Client
	Log   *zap.Logger
)
