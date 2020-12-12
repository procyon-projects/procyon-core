package core

import (
	"github.com/codnect/goo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockComponentProcessor struct {
	mock.Mock
}

func newMockComponentProcessor() mockComponentProcessor {
	return mockComponentProcessor{}
}

func (processor mockComponentProcessor) SupportsComponent(typ goo.Type) bool {
	results := processor.Called(typ)
	return results.Bool(0)
}

func (processor mockComponentProcessor) ProcessComponent(typ goo.Type) error {
	results := processor.Called(typ)
	return results.Error(0)
}

type testInterface interface {
	test()
}

type testComponent struct {
}

func newTestComponent() testComponent {
	return testComponent{}
}

func (testComponent) test() {

}

type testComponent2 struct {
}

func newTestComponentWithParam(str string) testComponent2 {
	return testComponent2{}
}

func (testComponent2) test() {

}

func init() {
	Register(newMockComponentProcessor)
	Register(newTestComponent)
	Register(newTestComponentWithParam)
}

func TestForEachComponentType(t *testing.T) {
	count := 0
	ForEachComponentType(func(componentName string, typ goo.Type) error {
		count++
		return nil
	})
	assert.Equal(t, 3, count)
}

func TestForEachComponentProcessor(t *testing.T) {
	mockComponentProcessorType := goo.GetType(newMockComponentProcessor)
	ForEachComponentProcessor(func(componentName string, typ goo.Type) error {
		assert.Equal(t, "github.com.procyon.projects.procyon.core.mockComponentProcessor", componentName)
		assert.True(t, mockComponentProcessorType.GetGoType() == typ.GetGoType())
		return nil
	})
}

func TestGetComponentTypes_WithNil(t *testing.T) {
	_, err := GetComponentTypes(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "type must not be null", err.Error())
}

func TestGetComponentTypes(t *testing.T) {
	types, err := GetComponentTypes(goo.GetType(testComponent{}))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(types))

	types, err = GetComponentTypes(goo.GetType((*testInterface)(nil)))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(types))
}

func TestGetComponentTypesWithParam(t *testing.T) {
	types, err := GetComponentTypesWithParam(goo.GetType(testComponent2{}), []goo.Type{goo.GetType("")})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(types))

	types, err = GetComponentTypesWithParam(goo.GetType(testComponent2{}), []goo.Type{goo.GetType("")})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(types))
}
