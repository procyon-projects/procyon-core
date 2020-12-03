package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getSimplePropertyResolver() *SimplePropertyResolver {
	propertySources := NewPropertySources()
	commandLinePropertySource := NewSimpleCommandLinePropertySource(getTestApplicationArguments())
	propertySources.Add(commandLinePropertySource)
	return NewSimplePropertyResolver(propertySources)
}

func TestSimplePropertyResolver_ContainsProperty(t *testing.T) {
	propertyResolver := getSimplePropertyResolver()
	assert.True(t, propertyResolver.ContainsProperty("procyon.application.name"))
	assert.False(t, propertyResolver.ContainsProperty("procyon.server.timeout"))
}

func TestSimplePropertyResolver_GetProperty(t *testing.T) {
	propertyResolver := getSimplePropertyResolver()
	assert.Equal(t, "\"Test Application\"", propertyResolver.GetProperty("procyon.application.name", ""))
	assert.Equal(t, "Default Timeout", propertyResolver.GetProperty("procyon.server.timeout", "Default Timeout"))
}
