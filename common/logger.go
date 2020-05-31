package common

import (
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewGinLogger() *zap.SugaredLogger {
	return newLogger(NewGinLoggerConfig())
}

func NewBizLogger() *zap.SugaredLogger {
	return newLogger(NewBizLoggerConfig())
}

func NewGinLoggerConfig() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "logs/gin.log",
		MaxSize:    10,
		MaxBackups: 100,
		MaxAge:     30,
		Compress:   false,
	}
}

func NewBizLoggerConfig() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "logs/biz.log",
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
