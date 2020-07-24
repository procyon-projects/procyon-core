// +build !linux

package core

import "golang.org/x/sys/windows"

func init() {
	GetThreadId = func() uint32 {
		return windows.GetCurrentThreadId()
	}
}
