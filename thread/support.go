package core

var getThreadIdFunc func() uint32

func GetThreadId() uint32 {
	return getThreadIdFunc()
}
