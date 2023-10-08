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

	hook := lumberjack.Logger{
		Filename:   logFilePath, // Log file path
		MaxSize:    100,         // The maximum size of each log file saved in M
		MaxBackups: 20,          // Maximum number of backups to log files
		MaxAge:     14,          // The maximum number of days the file can be saved
		Compress:   true,        // whether to compress
	}

	// set log level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(lv)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var cores []zapcore.Core
	cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), // encoder
		zapcore.AddSync(&hook), // stdout or file
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
