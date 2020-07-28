package core

import "sync"

var (
	loggerPool sync.Pool
)

func initProxyLoggerPool() {
	loggerPool = sync.Pool{
		New: newProxyLogger,
	}
}
