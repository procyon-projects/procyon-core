package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

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

func NewSimpleLogger(appId string, contextId string) *SimpleLogger {
	log := &SimpleLogger{
		&logrus.Logger{
			Out:       os.Stdout,
			Formatter: NewLogFormatter(appId, contextId),
			Level:     logrus.InfoLevel,
		},
	}
	return log
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

type LogFormatter struct {
	logrus.TextFormatter
	appId     string
	contextId string
}

func NewLogFormatter(appId string, contextId string) *LogFormatter {
	formatter := &LogFormatter{
		appId:     appId,
		contextId: contextId,
	}
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	return formatter
}

func (f *LogFormatter) GetApplicationId() string {
	return f.appId
}

func (f *LogFormatter) GetContextId() string {
	return f.contextId
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

	logContextId := f.GetContextId()
	separatorIndex := strings.Index(logContextId, "-")
	logContextId = logContextId[:separatorIndex]

	return []byte(
		fmt.Sprintf("[%s] \x1b[%dm%-7s\x1b[0m %s : %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), logContextId, entry.Message)), nil
}
