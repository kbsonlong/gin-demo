package core

import (
	"github.com/kbsonlong/gin-demo/pkg/global"
	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var level zapcore.Level

func Zap() (logger *zap.Logger) {

	writeSyncer := getLogWriter(
		global.CONFIG.Zap.Filename,
		global.CONFIG.Zap.MaxSize,
		global.CONFIG.Zap.MaxBackups,
		global.CONFIG.Zap.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	l.UnmarshalText([]byte(global.CONFIG.Zap.Level))
	core := zapcore.NewCore(encoder, writeSyncer, l)

	logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return logger
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	if global.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
