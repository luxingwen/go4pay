package logger

import (
	"fmt"
	"go4pay/pkg/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	filename := fmt.Sprintf("%s/%s", config.GetLogConfig().Dir, config.GetLogConfig().Filename)
	logFile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	Logger = zap.New(zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(zap.InfoLevel),
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(logFile),
			zap.NewAtomicLevelAt(zap.InfoLevel),
		),
	))
	defer Logger.Sync()
}

func Info(fields ...interface{}) {
	Infof("", fields...)
}

func Infof(message string, fields ...interface{}) {
	Logger.Sugar().Infof(message, fields)
}

func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}

func Errorf(messages string, fields ...interface{}) {
	Logger.Sugar().Errorf(messages, fields)
}
