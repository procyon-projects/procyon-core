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

func configureLog() {
	appId := GetApplicationId()
	Log.SetExtensibleLogFormatter(NewSimpleExtensibleLogFormatter(appId.String(), appId.String()))
}

type LoggerProvider interface {
	GetLogger() Logger
}

type LogLevel uint32

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type Logger interface {
	SetLevel()
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
	log *logrus.Logger
}

func NewSimpleLogger() *SimpleLogger {
	log := &SimpleLogger{
		&logrus.Logger{
			Out:       os.Stdout,
			Formatter: NewLogFormatter(),
			Level:     logrus.InfoLevel,
		},
	}
	return log
}

func (l *SimpleLogger) SetExtensibleLogFormatter(formatter ExtensibleLogFormatter) {
	l.log.Formatter.(*LogFormatter).ExtensibleLogFormatter = formatter
}

func (l *SimpleLogger) Trace(args ...interface{}) {
	l.log.Trace(args...)
}

func (l *SimpleLogger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *SimpleLogger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *SimpleLogger) Print(args ...interface{}) {
	l.log.Print(args...)
}

func (l SimpleLogger) Warning(args ...interface{}) {
	l.log.Warning(args...)
}

func (l *SimpleLogger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *SimpleLogger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l *SimpleLogger) Panic(args ...interface{}) {
	l.log.Panic(args...)
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
