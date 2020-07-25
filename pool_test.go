package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
}

func TestGetExistingInstanceFromPool(t *testing.T) {
	PoolManager.Register(GetType((*Test)(nil)), func() interface{} {
		return &Test{}
	})
	instance1, _ := PoolManager.Get(GetType((*Test)(nil)))
	PoolManager.Put(instance1)
	instance2, _ := PoolManager.Get(GetType((*Test)(nil)))
	assert.ObjectsAreEqual(instance1, instance2)
}
