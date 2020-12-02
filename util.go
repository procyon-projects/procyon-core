package core

import (
	"errors"
	"github.com/codnect/goo"
	"time"
	"unsafe"
)

type TaskWatch struct {
	taskName  string
	startTime time.Time
	totalTime time.Duration
	isRunning bool
}

func NewTaskWatch() *TaskWatch {
	return &TaskWatch{
		taskName: "[empty_task]",
	}
}

func NewTaskWatchWithName(taskName string) *TaskWatch {
	return &TaskWatch{
		taskName: taskName,
	}
}

func (watch *TaskWatch) Start() error {
	if watch.isRunning {
		return errors.New("TaskWatch is already running")
	}
	watch.startTime = time.Now()
	watch.isRunning = true
	return nil
}

func (watch *TaskWatch) Stop() error {
	if !watch.isRunning {
		return errors.New("TaskWatch is not running")
	}
	watch.isRunning = true
	watch.totalTime = time.Since(watch.startTime)
	watch.taskName = ""
	return nil
}

func (watch *TaskWatch) IsRunning() bool {
	return watch.isRunning
}

func (watch *TaskWatch) GetTotalTime() int64 {
	return watch.totalTime.Nanoseconds()
}

func HasFunctionSameParametersWithGivenParameters(componentType goo.Type, parameterTypes []goo.Type) bool {
	fun := componentType.ToFunctionType()
	functionParameterCount := fun.GetFunctionParameterCount()
	if parameterTypes == nil && functionParameterCount == 0 {
		return true
	} else if len(parameterTypes) != functionParameterCount || parameterTypes == nil && functionParameterCount != 0 {
		return false
	}
	inputParameterTypes := fun.GetFunctionParameterTypes()
	for index, inputParameterType := range inputParameterTypes {
		if !inputParameterType.Equals(parameterTypes[index]) {
			return false
		}
	}
	return true
}

func BytesToStr(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
