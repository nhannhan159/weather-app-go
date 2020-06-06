package common

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewGinLogger(config GlobalConfig) *zap.SugaredLogger {
	return newLogger(NewGinLoggerConfig(config))
}

func NewBizLogger(config GlobalConfig) *zap.SugaredLogger {
	return newLogger(NewBizLoggerConfig(config))
}

func NewGinLoggerConfig(config GlobalConfig) *lumberjack.Logger {
	fileName := fmt.Sprintf("%s/logs/%s/gin.log", config.BaseDir, config.AppName)
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 100,
		MaxAge:     30,
		Compress:   false,
	}
}

func NewBizLoggerConfig(config GlobalConfig) *lumberjack.Logger {
	fileName := fmt.Sprintf("%s/logs/%s/biz.log", config.BaseDir, config.AppName)
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 100,
		MaxAge:     30,
		Compress:   false,
	}
}

func newLogger(writer io.Writer) *zap.SugaredLogger {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer)),
		zap.InfoLevel,
	)
	logger := zap.New(core, zap.AddCaller(), zap.Development())
	defer logger.Sync()
	return logger.Sugar()
}
