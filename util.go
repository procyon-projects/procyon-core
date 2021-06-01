package core

import (
	"errors"
	"strconv"
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

	watch.isRunning = false
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

func BytesToStr(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func FlatMap(m map[string]interface{}) map[string]interface{} {
	flattenMap := map[string]interface{}{}

	for key, value := range m {
		switch child := value.(type) {
		case map[string]interface{}:
			nm := FlatMap(child)

			for nk, nv := range nm {
				flattenMap[key+"."+nk] = nv
			}
		case []interface{}:
			for i := 0; i < len(child); i++ {
				flattenMap[key+"."+strconv.Itoa(i)] = child[i]
			}
		default:
			flattenMap[key] = value
		}
	}

	return flattenMap
}
