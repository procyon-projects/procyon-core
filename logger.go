package core

import (
	"fmt"
	"github.com/google/uuid"
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
	Clone(contextId uuid.UUID) Logger
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

func (l *SimpleLogger) Trace(contextId string, args ...interface{}) {
	l.log.WithFields(logrus.Fields{
		"context_id": contextId,
	}).Trace(args...)
}

func (l *SimpleLogger) Debug(contextId string, args ...interface{}) {
	l.log.WithFields(logrus.Fields{
		"context_id": contextId,
	}).Debug(args...)
}

func (l *SimpleLogger) Info(contextId string, args ...interface{}) {
	l.log.WithFields(logrus.Fields{
		"context_id": contextId,
	}).Info(args...)
}

func (l *SimpleLogger) Print(contextId string, args ...interface{}) {
	l.log.WithFields(logrus.Fields{
		"context_id": contextId,
	}).Print(args...)
}

func (l SimpleLogger) Warning(contextId string, args ...interface{}) {
	l.log.WithFields(logrus.Fields{
		"context_id": contextId,
	}).Warning(args...)
}

func (l *SimpleLogger) Error(contextId string, args ...interface{}) {
	l.log.WithFields(logrus.Fields{
		"context_id": contextId,
	}).Error(args...)
}

func (l *SimpleLogger) Fatal(contextId string, args ...interface{}) {
	l.log.WithFields(logrus.Fields{
		"context_id": contextId,
	}).Fatal(args...)
}

func (l *SimpleLogger) Panic(contextId string, args ...interface{}) {
	l.log.WithFields(logrus.Fields{
		"context_id": contextId,
	}).Panic(args...)
}

type LogFormatter struct {
	logrus.TextFormatter
}

func NewLogFormatter() *LogFormatter {
	formatter := &LogFormatter{}
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	return formatter
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

	logContextId := entry.Data["context_id"].(string)
	separatorIndex := strings.Index(logContextId, "-")
	logContextId = logContextId[:separatorIndex]

	return []byte(
		fmt.Sprintf("[%s] \x1b[%dm%-7s\x1b[0m %s : %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), logContextId, entry.Message)), nil
}

func newProxyLogger() interface{} {
	return &ProxyLogger{}
}

func getProxyLoggerFromPool() *ProxyLogger {
	return GetFromPool(proxyLoggerType).(*ProxyLogger)
}

type ProxyLogger struct {
	logger    *SimpleLogger
	contextId string
}

func NewProxyLogger(logger *SimpleLogger, contextId uuid.UUID) *ProxyLogger {
	return &ProxyLogger{
		logger,
		contextId.String(),
	}
}

func (l *ProxyLogger) Clone(contextId uuid.UUID) Logger {
	cloneLogger := getProxyLoggerFromPool()
	cloneLogger.contextId = contextId.String()
	cloneLogger.logger = l.logger
	return cloneLogger
}

func (l *ProxyLogger) Trace(args ...interface{}) {
	l.logger.Trace(l.contextId, args...)
}

func (l *ProxyLogger) Debug(args ...interface{}) {
	l.logger.Debug(l.contextId, args...)
}

func (l *ProxyLogger) Info(args ...interface{}) {
	l.logger.Info(l.contextId, args...)
}

func (l *ProxyLogger) Print(args ...interface{}) {
	l.logger.Print(l.contextId, args...)
}

func (l ProxyLogger) Warning(args ...interface{}) {
	l.logger.Warning(l.contextId, args...)
}

func (l *ProxyLogger) Error(args ...interface{}) {
	l.logger.Error(l.contextId, args...)
}

func (l *ProxyLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(l.contextId, args...)
}

func (l *ProxyLogger) Panic(args ...interface{}) {
	l.logger.Panic(l.contextId, args...)
}
