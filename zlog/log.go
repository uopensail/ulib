package zlog

import (
	"fmt"
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
)

var LOG *zap.Logger
var SLOG *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("ulib log error :%v", err))
	}
	zap.ReplaceGlobals(logger)
	LOG = zap.L()
	SLOG = zap.S()
}

func InitLogger(appName string, debug bool, logDir string) {
	var logFilePath string
	var lv zapcore.Level
	if debug {
		logFilePath = fmt.Sprintf("%s/%s_debug.log", logDir, appName)
		lv = zap.DebugLevel
	} else {
		logFilePath = fmt.Sprintf("%s/%s_release.log", logDir, appName)
		lv = zap.InfoLevel
	}

	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   logFilePath, // 日志文件路径
		MaxSize:    100,         // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 20,          // 日志文件最多保存多少个备份
		MaxAge:     14,          // 文件最多保存多少天
		Compress:   true,        // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(lv)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var cores []zapcore.Core
	cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), // 编码器配置
		zapcore.AddSync(&hook), // 打印到控制台和文件
		atomicLevel))
	if debug {
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout), atomicLevel))
	}
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller())
	if debug {
		logger.WithOptions(zap.AddStacktrace(zap.WarnLevel))
	}
	zap.ReplaceGlobals(logger)
	LOG = zap.L()
	SLOG = zap.S()
}
