package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapPropertySource(t *testing.T) {
	Log.Trace("Something very low level.")
	Log.Debug("Useful debugging information.")
	Log.Info("Something noteworthy happened!")
	Log.Warning("You should probably take a look at this.")
	Log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	Log.Fatal("Bye.")
	// Calls panic() after logging
	Log.Panic("I'm bailing.")
	testMap := map[string]interface{}{
		"test":  "hello",
		"test2": "world",
	}
	mapPropertySource := NewMapPropertySource("testMap", testMap)
	assert.Equal(t, "testMap", mapPropertySource.GetName())
	//assert.Equal(t, 2, len(mapPropertySource.GetSource().(map[string]interface{})))
	assert.Equal(t, "hello", mapPropertySource.GetProperty("test"))
	assert.Equal(t, 2, len(mapPropertySource.GetPropertyNames()))
	assert.Equal(t, true, mapPropertySource.ContainsProperty("test2"))

	var propertySource PropertySource = mapPropertySource
	assert.Equal(t, "testMap", propertySource.GetName())
	assert.Equal(t, 2, len(propertySource.GetSource().(map[string]interface{})))

	var enumerablePropertySource EnumerablePropertySource = mapPropertySource
	assert.Equal(t, 2, len(enumerablePropertySource.GetPropertyNames()))
}
