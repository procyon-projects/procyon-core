package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	Log = NewSimpleLogger()
)

func configureLogger() {
	appId := GetApplicationId()
	Log.SetExtensibleLoggerFormatter(NewSimpleExtensibleLoggerFormatter(appId.String(), appId.String()))
}

type LoggerProvider interface {
	GetLogger() Logger
}

type Logger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type SimpleLogger struct {
	log logrus.Logger
}

func NewSimpleLogger() *SimpleLogger {
	logger := &SimpleLogger{}
	logger.log.Out = os.Stdout
	logger.log.Formatter = NewLoggerFormatter()
	return logger
}

func (logger *SimpleLogger) SetExtensibleLoggerFormatter(formatter ExtensibleLoggerFormatter) {
	logger.log.Formatter.(*LoggerFormatter).ExtensibleLoggerFormatter = formatter
}

func (logger *SimpleLogger) Trace(args ...interface{}) {
	logger.log.Trace(args)
}

func (logger *SimpleLogger) Debug(args ...interface{}) {
	logger.log.Debug(args)
}

func (logger *SimpleLogger) Info(args ...interface{}) {
	logger.log.Info(args)
}

func (logger *SimpleLogger) Print(args ...interface{}) {
	logger.log.Print(args)
}

func (logger SimpleLogger) Warning(args ...interface{}) {
	logger.log.Warning(args)
}

func (logger *SimpleLogger) Error(args ...interface{}) {
	logger.log.Error(args)
}

func (logger *SimpleLogger) Fatal(args ...interface{}) {
	logger.log.Fatal(args)
}

func (logger *SimpleLogger) Panic(args ...interface{}) {
	logger.log.Panic(args)
}

type ExtensibleLoggerFormatter interface {
	GetApplicationId() string
	GetContextId() string
}

type SimpleExtensibleLoggerFormatter struct {
	appId     string
	contextId string
}

func NewSimpleExtensibleLoggerFormatter(appId string, contextId string) SimpleExtensibleLoggerFormatter {
	return SimpleExtensibleLoggerFormatter{
		appId,
		contextId,
	}
}

func (e SimpleExtensibleLoggerFormatter) GetApplicationId() string {
	return e.appId
}

func (e SimpleExtensibleLoggerFormatter) GetContextId() string {
	return e.contextId
}

type LoggerFormatter struct {
	logrus.TextFormatter
	ExtensibleLoggerFormatter
}

func NewLoggerFormatter() *LoggerFormatter {
	formatter := &LoggerFormatter{}
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	return formatter
}

func (f *LoggerFormatter) SetExtensibleLoggerFormatter(formatter ExtensibleLoggerFormatter) {
	f.ExtensibleLoggerFormatter = formatter
}

func (f *LoggerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 37 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	loggerContextId := ""
	if f.ExtensibleLoggerFormatter != nil {
		loggerContextId = f.GetContextId()
		separatorIndex := strings.Index(loggerContextId, "-")
		loggerContextId = loggerContextId[:separatorIndex]
	}
	return []byte(
		fmt.Sprintf("[%s] \x1b[%dm%-7s\x1b[0m %s : %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), loggerContextId, entry.Message)), nil
}
