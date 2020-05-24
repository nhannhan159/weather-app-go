package common

import (
	"os"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//func NewBizLogger() *zap.SugaredLogger {
//	writerSyncer := NewBizLoggerConfig()
//	encoder := getEncoder()
//	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
//	logger := zap.New(core)
//	defer logger.Sync()
//	return logger.Sugar()
//}

func GinLoggerConfig() {

}

//func InitLogger() {
//	writerSyncer := getLogWriter()
//	encoder := getEncoder()
//	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
//	logger := zap.New(core)
//	sugarLogger = logger.Sugar()
//}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
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

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
