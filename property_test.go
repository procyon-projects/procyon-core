package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPropertySources_Add(t *testing.T) {
	propertySources := NewPropertySources()
	argPropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Add(argPropertySource)
	assert.Equal(t, argPropertySource, propertySources.GetPropertyResources()[0])
	assert.Equal(t, 1, propertySources.GetSize())
}

func TestPropertySources_Get(t *testing.T) {
	propertySources := NewPropertySources()
	argPropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Add(argPropertySource)
	propertySource, ok := propertySources.Get(ProcyonApplicationCommandLinePropertySource)
	assert.True(t, ok)
	assert.Equal(t, argPropertySource, propertySource)
}

func TestPropertySources_GetSize(t *testing.T) {
	propertySources := NewPropertySources()
	argPropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Add(argPropertySource)
	assert.Equal(t, 1, propertySources.GetSize())
}

func TestPropertySources_Remove(t *testing.T) {
	propertySources := NewPropertySources()
	argPropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Add(argPropertySource)
	propertySources.Remove(ProcyonApplicationCommandLinePropertySource)
	assert.Equal(t, 0, propertySources.GetSize())
}

func TestPropertySources_Replace(t *testing.T) {
	propertySources := NewPropertySources()
	argPropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Add(argPropertySource)
	anotherArgPropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Replace(ProcyonApplicationCommandLinePropertySource, anotherArgPropertySource)
	propertySource, _ := propertySources.Get(ProcyonApplicationCommandLinePropertySource)
	assert.Equal(t, anotherArgPropertySource, propertySource)
}

func TestPropertySources_RemoveIfPresent(t *testing.T) {
	propertySources := NewPropertySources()
	argPropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Add(argPropertySource)
	propertySources.RemoveIfPresent(argPropertySource)
	assert.Equal(t, 0, propertySources.GetSize())
}

func TestPropertySources_GetPropertyResources(t *testing.T) {
	propertySources := NewPropertySources()
	argPropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Add(argPropertySource)
	assert.Equal(t, 1, len(propertySources.GetPropertyResources()))
}
