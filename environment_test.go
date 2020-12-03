package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStandardEnvironment_ContainsProperty(t *testing.T) {
	standardEnvironment := NewStandardEnvironment()
	standardEnvironment.GetPropertySources().Add(NewSimpleCommandLinePropertySource(getTestApplicationArguments()))
	assert.True(t, standardEnvironment.ContainsProperty("procyon.application.name"))
	assert.False(t, standardEnvironment.ContainsProperty("procyon.server.timeout"))
}

func TestStandardEnvironment_GetProperty(t *testing.T) {
	standardEnvironment := NewStandardEnvironment()
	standardEnvironment.GetPropertySources().Add(NewSimpleCommandLinePropertySource(getTestApplicationArguments()))
	assert.NotNil(t, standardEnvironment.GetProperty("procyon.application.name", "default value"))
	assert.Equal(t, "default value", standardEnvironment.GetProperty("procyon.server.timeout", "default value"))
}

func TestStandardEnvironment_GetPropertySources(t *testing.T) {
	standardEnvironment := NewStandardEnvironment()
	standardEnvironment.GetPropertySources().Add(NewSimpleCommandLinePropertySource(getTestApplicationArguments()))
	assert.Equal(t, 1, standardEnvironment.GetPropertySources().GetSize())
}

func TestStandardEnvironment_GetSystemEnvironment(t *testing.T) {
	standardEnvironment := NewStandardEnvironment()
	assert.NotNil(t, standardEnvironment.GetSystemEnvironment())
}

func TestStandardEnvironment_GetTypeConverterService(t *testing.T) {
	standardEnvironment := NewStandardEnvironment()
	assert.NotNil(t, standardEnvironment.GetTypeConverterService())
}
