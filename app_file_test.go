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

	filePaths := []string{"resources/procyon.test.yaml", "resources/procyon.db.yaml"}

	propertyMap, err := parser.Parse(filePaths)
	assert.NoError(t, err)
	assert.NotNil(t, propertyMap)

	assert.Contains(t, propertyMap, "procyon.application.name")
	assert.Equal(t, "Procyon Test Application", propertyMap["procyon.application.name"])

	assert.Contains(t, propertyMap, "logging.level")
	assert.Equal(t, "INFO", propertyMap["logging.level"])

	assert.Contains(t, propertyMap, "server.port")
	assert.Equal(t, 8095, propertyMap["server.port"])

	// check the properties in procyon.db.yaml
}
