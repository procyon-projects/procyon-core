package core

import (
	"errors"
	"github.com/codnect/goo"
	"time"
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

func GetMapKeys(mapObj interface{}) []string {
	argMapKeys := goo.GetType(mapObj).GetGoValue().MapKeys()
	mapKeys := make([]string, len(argMapKeys))
	for i := 0; i < len(argMapKeys); i++ {
		mapKeys[i] = argMapKeys[i].String()
	}
	return mapKeys
}

func HasFunctionSameParametersWithGivenParameters(componentType goo.Type, parameterTypes []goo.Type) bool {
	fun := componentType.(goo.Function)
	functionParameterCount := fun.GetFunctionParameterCount()
	if len(parameterTypes) != functionParameterCount {
		return false
	}
	inputParameterTypes := fun.GetFunctionReturnTypes()
	for index, inputParameterType := range inputParameterTypes {
		if !inputParameterType.Equals(parameterTypes[index]) {
			return false
		}
	}
	return true
}
