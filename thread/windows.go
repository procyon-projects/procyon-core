// +build !linux

package core

import "golang.org/x/sys/windows"

func init() {
	getThreadIdFunc = func() uint32 {
		return windows.GetCurrentThreadId()
	}
}
