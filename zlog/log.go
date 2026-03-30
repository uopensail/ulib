package zlog

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
)

// LOG is the global zap.Logger instance
var LOG *zap.Logger

// SLOG is the global zap.SugaredLogger instance
var SLOG *zap.SugaredLogger

// init initializes the logger with default production configuration.
// It sets up stdout as the output destination and replaces the global loggers.
// Panics if the logger cannot be initialized.
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

// InitLogger initializes a customized logger with file rotation and configurable
// settings. appName is used in the log filename; logDir is the directory where
// log files are written. When debug is true, the level is set to Debug and
// output is also mirrored to stdout with stack traces on Warn and above.
func InitLogger(appName string, debug bool, logDir string) {
	// Create log directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create log directory: %v", err))
	}

	suffix := "release"
	lv := zap.InfoLevel
	if debug {
		suffix = "debug"
		lv = zap.DebugLevel
	}
	logFilePath := filepath.Join(logDir, appName+"_"+suffix+".log")

	hook := lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    100, // MiB
		MaxBackups: 20,
		MaxAge:     14, // days
		Compress:   true,
	}

	// Set log level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(lv)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var cores []zapcore.Core

	// File core with console encoder for better readability
	fileWriteSyncer := zapcore.AddSync(&hook)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		fileWriteSyncer,
		atomicLevel,
	)
	cores = append(cores, fileCore)

	// Add stdout core in debug mode
	if debug {
		stdoutCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			atomicLevel,
		)
		cores = append(cores, stdoutCore)
	}

	core := zapcore.NewTee(cores...)

	// Create logger with caller information
	logger := zap.New(core, zap.AddCaller())

	// Add stacktrace for warn level and above in debug mode
	if debug {
		logger = logger.WithOptions(zap.AddStacktrace(zap.WarnLevel))
	}

	// Replace global loggers
	zap.ReplaceGlobals(logger)
	LOG = zap.L()
	SLOG = zap.S()
}
