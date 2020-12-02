package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBytesToStr(t *testing.T) {
	assert.Equal(t, "Test", BytesToStr([]byte("Test")))
}

func TestTaskWatch(t *testing.T) {
	taskWatch := NewTaskWatch()
	assert.Equal(t, "[empty_task]", taskWatch.taskName)
	assert.Equal(t, time.Duration(0), taskWatch.totalTime)
	assert.Equal(t, int64(0), taskWatch.GetTotalTime())

	// start
	err := taskWatch.Start()
	assert.Nil(t, err)
	assert.True(t, taskWatch.IsRunning())

	err = taskWatch.Start()
	assert.NotNil(t, err)

	time.Sleep(1000)

	// stop
	err = taskWatch.Stop()
	assert.Nil(t, err)
	assert.False(t, taskWatch.IsRunning())

	err = taskWatch.Stop()
	assert.NotNil(t, err)

	assert.NotEqual(t, int64(0), taskWatch.GetTotalTime())
}
