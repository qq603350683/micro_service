package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	// 日志文件名称 C:\ProgramData\Elastic\Beats\filebeat
	filename := "micro.log"
	syncWriter := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   filename,
			MaxSize:    512,  // MB
			// MaxAge:     0,
			MaxBackups: 1,   // 最大备份数
			LocalTime:  true,
			Compress:   true,  // 是否启用压缩
		})

	// 编码
	encoder := zap.NewProductionEncoderConfig()

	// 时间格式
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		// 编码器
		zapcore.NewJSONEncoder(encoder),
		syncWriter,
		zap.NewAtomicLevelAt(zap.DebugLevel))

	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1))

	logger = log.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}