package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	Log = NewSimpleLog()
)

func configureLog() {
	appId := GetApplicationId()
	Log.SetExtensibleLogFormatter(NewSimpleExtensibleLogFormatter(appId.String(), appId.String()))
}

type LogProvider interface {
	GetLog() Logger
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

func NewSimpleLog() *SimpleLogger {
	Log := &SimpleLogger{}
	Log.log.Out = os.Stdout
	Log.log.Formatter = NewLogFormatter()
	return Log
}

func (Log *SimpleLogger) SetExtensibleLogFormatter(formatter ExtensibleLogFormatter) {
	Log.log.Formatter.(*LogFormatter).ExtensibleLogFormatter = formatter
}

func (Log *SimpleLogger) Trace(args ...interface{}) {
	Log.log.Trace(args)
}

func (Log *SimpleLogger) Debug(args ...interface{}) {
	Log.log.Debug(args)
}

func (Log *SimpleLogger) Info(args ...interface{}) {
	Log.log.Info(args)
}

func (Log *SimpleLogger) Print(args ...interface{}) {
	Log.log.Print(args)
}

func (Log SimpleLogger) Warning(args ...interface{}) {
	Log.log.Warning(args)
}

func (Log *SimpleLogger) Error(args ...interface{}) {
	Log.log.Error(args)
}

func (Log *SimpleLogger) Fatal(args ...interface{}) {
	Log.log.Fatal(args)
}

func (Log *SimpleLogger) Panic(args ...interface{}) {
	Log.log.Panic(args)
}

type ExtensibleLogFormatter interface {
	GetApplicationId() string
	GetContextId() string
}

type SimpleExtensibleLogFormatter struct {
	appId     string
	contextId string
}

func NewSimpleExtensibleLogFormatter(appId string, contextId string) SimpleExtensibleLogFormatter {
	return SimpleExtensibleLogFormatter{
		appId,
		contextId,
	}
}

func (e SimpleExtensibleLogFormatter) GetApplicationId() string {
	return e.appId
}

func (e SimpleExtensibleLogFormatter) GetContextId() string {
	return e.contextId
}

type LogFormatter struct {
	logrus.TextFormatter
	ExtensibleLogFormatter
}

func NewLogFormatter() *LogFormatter {
	formatter := &LogFormatter{}
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	return formatter
}

func (f *LogFormatter) SetExtensibleLogFormatter(formatter ExtensibleLogFormatter) {
	f.ExtensibleLogFormatter = formatter
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
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
	LogContextId := ""
	if f.ExtensibleLogFormatter != nil {
		LogContextId = f.GetContextId()
		separatorIndex := strings.Index(LogContextId, "-")
		LogContextId = LogContextId[:separatorIndex]
	}
	return []byte(
		fmt.Sprintf("[%s] \x1b[%dm%-7s\x1b[0m %s : %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), LogContextId, entry.Message)), nil
}
