package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Logger struct {
	LogFile    *logrus.Logger
	LogConsole *logrus.Logger
	service    string
	version    string
}

const basePath = "logs"

func initialize(service string, version string, pathFile string, levelStr string) *Logger {
	var level logrus.Level

	switch levelStr {
	case "trace":
		{
			level = logrus.TraceLevel
		}
	case "debug":
		{
			level = logrus.DebugLevel
		}
	case "info":
		{
			level = logrus.InfoLevel
		}
	case "warning":
		{
			level = logrus.WarnLevel
		}
	case "error":
		{
			level = logrus.ErrorLevel
		}
	case "fatal":
		{
			level = logrus.FatalLevel
		}
	case "panic":
		{
			level = logrus.PanicLevel
		}
	default:
		{
			level = logrus.InfoLevel
		}
	}

	logFile := logrus.New()
	logFile.SetFormatter(&logrus.JSONFormatter{})
	file, _ := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logFile.SetOutput(file)
	logFile.SetLevel(level)

	logConsole := logrus.New()
	logConsole.SetFormatter(&logrus.JSONFormatter{})
	logConsole.SetOutput(os.Stdout)
	logConsole.SetLevel(level)

	return &Logger{
		LogFile:    logFile,
		LogConsole: logConsole,
		service:    service,
		version:    version,
	}
}

var lgr *Logger

func NewLogger(config *viper.Viper) *Logger {
	if lgr == nil {
		today := time.Now().Format("2006-01-02")
		pathFile := fmt.Sprintf("%s-%s.log", basePath, today)
		lgr = initialize(config.GetString("app.name"), config.GetString("app.version"), pathFile, config.GetString("app.logLevel"))
	}
	return lgr
}

func getCallerFilePath() string {
	_, file, line, _ := runtime.Caller(3)
	return string(fmt.Sprint(file, ":", line))
}
