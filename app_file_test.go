package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppFilePropertySource(t *testing.T) {
	propertySource := NewAppFilePropertySource("dev, db")
	value := propertySource.GetProperty("procyon.application.name")
	if value == "" {

	}
}

func TestAppFileParser_Parse(t *testing.T) {
	parser := NewAppFileParser()

	filePaths := []string{"test-resources/procyon.test.yaml", "test-resources/procyon.db.yaml"}

	propertyMap, err := parser.Parse(filePaths)
	assert.NoError(t, err)
	assert.NotNil(t, propertyMap)

	assert.Contains(t, propertyMap, "procyon.application.name")
	assert.Equal(t, "Procyon Test Application", propertyMap["procyon.application.name"])

	assert.Contains(t, propertyMap, "logging.level")
	assert.Equal(t, "INFO", propertyMap["logging.level"])

	assert.Contains(t, propertyMap, "server.port")
	assert.Equal(t, 8095, propertyMap["server.port"])

	assert.Contains(t, propertyMap, "procyon.datasource.url")
	assert.Equal(t, "test-url", propertyMap["procyon.datasource.url"])

	assert.Contains(t, propertyMap, "procyon.datasource.username")
	assert.Equal(t, "test-username", propertyMap["procyon.datasource.username"])

	assert.Contains(t, propertyMap, "procyon.datasource.password")
	assert.Equal(t, "test-password", propertyMap["procyon.datasource.password"])
}

func TestAppFileParser_ContainsProperty(t *testing.T) {
	propertySource := NewAppFilePropertySource("dev, db")
	assert.True(t, propertySource.ContainsProperty("procyon.application.name"))
	assert.False(t, propertySource.ContainsProperty("procyon.server.timeout"))
}

func TestAppFileParser_GetProperty(t *testing.T) {
	propertySource := NewAppFilePropertySource("dev")
	assert.NotNil(t, propertySource.GetProperty("procyon.application.name"))
	assert.Equal(t, "Procyon Dev Application", propertySource.GetProperty("procyon.application.name"))
}

func TestAppFileParser_GetPropertyNames(t *testing.T) {
	propertySource := NewAppFilePropertySource("dev")
	assert.NotNil(t, propertySource.GetPropertyNames())
}

func TestAppFileParser_GetName(t *testing.T) {
	propertySource := NewAppFilePropertySource("dev")
	assert.Equal(t, ProcyonAppFilePropertySource, propertySource.GetName())
}

func TestAppFileParser_GetSource(t *testing.T) {
	propertySource := NewAppFilePropertySource("dev")
	assert.NotNil(t, propertySource.GetSource())
}
