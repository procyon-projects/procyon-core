package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
}

func TestGetExistingInstanceFromPool(t *testing.T) {
	RegisterPool(GetType((*Test)(nil)), func() interface{} {
		return &Test{}
	})
	instance1 := GetFromPool(GetType((*Test)(nil)))
	PutToPool(instance1)
	instance2 := GetFromPool(GetType((*Test)(nil)))
	assert.ObjectsAreEqual(instance1, instance2)
}
