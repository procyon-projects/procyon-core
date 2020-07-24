// +build linux

package core

import "golang.org/x/sys/unix"

func init() {
	getThreadIdFunc = func() uint32 {
		threadId := uint32(unix.Gettid())
		return threadId
	}
}
