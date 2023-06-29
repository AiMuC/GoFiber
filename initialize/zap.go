package initialize

import (
	"fmt"
	"github.com/aimuc/gofiber/utils"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.ErrorLevel
	})
	// panic 级别
	errorPanic := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.PanicLevel
	})
	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/%s/server_debug.log", utils.Env("ZAP.LOG_DIR", "runtime/log").(string), time.Now().Format("20060102")), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/server_info.log", utils.Env("ZAP.LOG_DIR", "runtime/log").(string), time.Now().Format("20060102")), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/server_warn.log", utils.Env("ZAP.LOG_DIR", "runtime/log").(string), time.Now().Format("20060102")), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/server_error.log", utils.Env("ZAP.LOG_DIR", "runtime/log").(string), time.Now().Format("20060102")), errorPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/server_panic.log", utils.Env("ZAP.LOG_DIR", "runtime/log").(string), time.Now().Format("20060102")), errorPanic),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	logger = logger.WithOptions(zap.AddCaller())
	zap.ReplaceGlobals(logger) // 替换全局记录器
	return logger
}

func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    2,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 10,   // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}
	if utils.Env("ZAP.LOG_IN_CONSOLE", true).(bool) { // 是否打印在控制台 Stdout
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := getWriteSyncer(fileName) // 使用lumberjack进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

func getEncoder() zapcore.Encoder {
	if utils.Env("ZAP.FORMAT", "console").(string) == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case utils.Env("ZAP.ENCODE_LEVEL", "CapitalLevelEncoder").(string) == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case utils.Env("ZAP.ENCODE_LEVEL", "CapitalLevelEncoder").(string) == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case utils.Env("ZAP.ENCODE_LEVEL", "CapitalLevelEncoder").(string) == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case utils.Env("ZAP.ENCODE_LEVEL", "CapitalLevelEncoder").(string) == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(utils.Env("ZAP.PREFIX", "[GoFiber]").(string) + "2006/01/02 - 15:04:05.000"))
}
