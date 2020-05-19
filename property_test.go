package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type BasePropertySource struct {
	AbstractPropertySource
}

func NewBasePropertySource(name string, source interface{}) BasePropertySource {
	basePropertySource := BasePropertySource{
		NewAbstractPropertySourceWithSource(name, source),
	}
	basePropertySource.PropertySource = basePropertySource
	return basePropertySource
}

func (source BasePropertySource) GetProperty(name string) interface{} {
	return "test"
}

func (source BasePropertySource) ContainsProperty(name string) bool {
	return true
}

func TestAbstractPropertySource(t *testing.T) {
	source := NewBasePropertySource("test", "hello")
	assert.Equal(t, "test", source.GetName())
	assert.Equal(t, "hello", source.GetSource().(string))
	assert.Equal(t, "test", source.GetProperty("test").(string))
	assert.Equal(t, true, source.ContainsProperty("test"))
}

func TestMapPropertySource(t *testing.T) {
	testMap := map[string]interface{}{
		"test":  "hello",
		"test2": "world",
	}
	mapPropertySource := NewMapPropertySource("testMap", testMap)
	assert.Equal(t, "testMap", mapPropertySource.GetName())
	assert.Equal(t, 2, len(mapPropertySource.GetSource().(map[string]interface{})))
	assert.Equal(t, "hello", mapPropertySource.GetProperty("test"))
	assert.Equal(t, 2, len(mapPropertySource.GetPropertyNames()))
	assert.Equal(t, true, mapPropertySource.ContainsProperty("test2"))

	var propertySource PropertySource = mapPropertySource
	assert.Equal(t, "testMap", propertySource.GetName())
	assert.Equal(t, 2, len(propertySource.GetSource().(map[string]interface{})))

	var enumerablePropertySource EnumerablePropertySource = mapPropertySource
	assert.Equal(t, 2, len(enumerablePropertySource.GetPropertyNames()))
}
