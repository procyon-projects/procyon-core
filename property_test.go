package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"runtime"
	"testing"
	"time"
)

func TestMapPropertySource(t *testing.T) {
	runtime.GOMAXPROCS(20)
	log.Print("")
	time.Sleep(10e9)
	NewSimpleLogger().Print("", "Hello")
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 1))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 1))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 2))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 2))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 3))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 3))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 4))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 4))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 5))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 5))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 6))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 6))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 7))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 7))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 8))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 8))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 9))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 9))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 10))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 10))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 11))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 11))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 12))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 12))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 12))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 12))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 13))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 13))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 14))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 14))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 15))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 15))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 16))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 16))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 17))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 17))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 18))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 18))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 19))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 19))
	}()
	go func() {
		NewSimpleLogger().Print("", fmt.Sprintf("Start %d", 20))
		for i := 0; i != 1; {
			continue
		}
		NewSimpleLogger().Print("", fmt.Sprintf("Finish %d", 20))
	}()
	for i := 0; i != 1; {
		continue
	}
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
