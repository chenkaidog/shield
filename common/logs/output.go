package logs

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	logFileName = "./log/%s.log"
	// LogFileMaxSize 每个日志文件最大 MB
	logFileMaxSize = 512
	// LogFileMaxBackups 保留日志文件个数
	logFileMaxBackups = 10
	// LogFileMaxAge 保留日志最大天数
	logFileMaxAge   = 14
	defaultLogLevel = LevelDebug

	envLogOutputFileName      = "logs_output_file_name"
	envLogsSetOutputLocalFile = "logs_set_local_file"
	envLogsLevel              = "logs_level"
)

func newOutput() io.Writer {
	if os.Getenv(envLogsSetOutputLocalFile) == "true" {
		return io.MultiWriter(
			os.Stdout,
			&lumberjack.Logger{
				Filename:   fmt.Sprintf(logFileName, os.Getenv(envLogOutputFileName)),
				MaxSize:    logFileMaxSize,
				MaxAge:     logFileMaxAge,
				MaxBackups: logFileMaxBackups,
				LocalTime:  true,
				Compress:   false,
			})
	}

	return os.Stdout
}

func getLogLevel() Level {
	defaultLevel := defaultLogLevel
	switch os.Getenv(envLogsLevel) {
	case "trace":
		defaultLevel = LevelTrace
	case "debug":
		defaultLevel = LevelDebug
	case "info":
		defaultLevel = LevelInfo
	case "warn":
		defaultLevel = LevelWarn
	case "error":
		defaultLevel = LevelError
	case "fatal":
		defaultLevel = LevelFatal
	}

	return defaultLevel
}
