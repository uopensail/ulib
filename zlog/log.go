package zlog

import (
	"fmt"
	"os"

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

// InitLogger initializes a customized logger with file rotation and configurable settings.
//
// @param appName: Name of the application used in log file naming
// @param debug: If true, enables debug level logging and console output
// @param logDir: Directory where log files will be stored
//
// The function sets up:
// - File logging with rotation using lumberjack
// - Console output in debug mode
// - Appropriate log levels based on debug flag
// - ISO8601 time encoding and caller information
func InitLogger(appName string, debug bool, logDir string) {
	// Create log directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create log directory: %v", err))
	}

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
		MaxSize:    100,         // Maximum size of each log file in megabytes
		MaxBackups: 20,          // Maximum number of backup log files
		MaxAge:     14,          // Maximum number of days to retain log files
		Compress:   true,        // Whether to compress backup log files
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
