package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Db    *gorm.DB
	Redis *redis.Client
	Log   *zap.Logger
)
