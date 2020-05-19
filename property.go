package core

import (
	"errors"
	"log"
)

type PropertySource interface {
	GetName() string
	GetSource() interface{}
	GetProperty(name string) interface{}
	ContainsProperty(name string) bool
}

type AbstractPropertySource struct {
	PropertySource
	name   string
	source interface{}
}

func NewAbstractPropertySourceWithSource(name string, source interface{}) AbstractPropertySource {
	propertySource := AbstractPropertySource{
		name:   name,
		source: source,
	}
	return propertySource
}

func (source AbstractPropertySource) GetName() string {
	return source.name
}

func (source AbstractPropertySource) GetSource() interface{} {
	return source.source
}

func (source AbstractPropertySource) GetProperty(name string) interface{} {
	panic("Implement me!. This is an abstract method. AbstractPropertySource.GetProperty(string)")
}

func (source AbstractPropertySource) ContainsProperty(name string) bool {
	panic("Implement me!. This is an abstract method. AbstractPropertySource.ContainsProperty(string)")
}

type EnumerablePropertySource interface {
	GetPropertyNames() []string
}

type AbstractEnumerablePropertySource struct {
	EnumerablePropertySource
	AbstractPropertySource
}

func NewAbstractEnumerablePropertySourceWithSource(name string, source interface{}) AbstractEnumerablePropertySource {
	propertySource := AbstractEnumerablePropertySource{
		AbstractPropertySource: NewAbstractPropertySourceWithSource(name, source),
	}
	return propertySource
}

func (source AbstractEnumerablePropertySource) GetPropertyNames() []string {
	panic("Implement me!. This is an abstract method. AbstractEnumerablePropertySource.GetPropertyNames()")
}

func (source AbstractEnumerablePropertySource) ContainsProperty(name string) bool {
	for _, propertyName := range source.EnumerablePropertySource.GetPropertyNames() {
		if propertyName == name {
			return true
		}
	}
	return false
}

type MapPropertySource struct {
	AbstractEnumerablePropertySource
}

func NewMapPropertySource(name string, source map[string]interface{}) MapPropertySource {
	mapPropertySource := MapPropertySource{
		NewAbstractEnumerablePropertySourceWithSource(name, source),
	}
	mapPropertySource.PropertySource = mapPropertySource
	mapPropertySource.EnumerablePropertySource = mapPropertySource
	return mapPropertySource
}

func (source MapPropertySource) GetProperty(name string) interface{} {
	propertyMap := source.GetSource().(map[string]interface{})
	return propertyMap[name]
}

func (source MapPropertySource) ContainsProperty(name string) bool {
	propertyMap := source.GetSource().(map[string]interface{})
	return propertyMap[name] != nil
}

func (source MapPropertySource) GetPropertyNames() []string {
	return GetMapKeys(source.GetSource())
}

type CompositePropertySource struct {
	AbstractEnumerablePropertySource
	sources []PropertySource
}

func NewCompositePropertySource(name string) CompositePropertySource {
	compositePropertySource := CompositePropertySource{
		AbstractEnumerablePropertySource: NewAbstractEnumerablePropertySourceWithSource(name, nil),
		sources:                          make([]PropertySource, 0),
	}
	compositePropertySource.PropertySource = compositePropertySource
	compositePropertySource.EnumerablePropertySource = compositePropertySource
	return compositePropertySource
}

func (source CompositePropertySource) GetProperty(name string) interface{} {
	for _, propertySource := range source.sources {
		property := propertySource.GetProperty(name)
		if property != nil {
			return property
		}
	}
	return nil
}

func (source CompositePropertySource) ContainsProperty(name string) bool {
	for _, propertySource := range source.sources {
		if propertySource.ContainsProperty(name) {
			return true
		}
	}
	return false
}

func (source CompositePropertySource) GetPropertyNames() []string {
	names := make([]string, 0)
	for _, propertySource := range source.sources {
		if source, ok := propertySource.(EnumerablePropertySource); ok {
			names = append(names, source.GetPropertyNames()...)
		} else {
			log.Fatal("Property source does not support except EnumerablePropertySource")
		}
	}
	return names
}

func (source CompositePropertySource) AddPropertySource(propertySource PropertySource) {
	source.sources = append(source.sources, propertySource)
}

func (source CompositePropertySource) AddFirstPropertySource(propertySource PropertySource) {
	newPropertySources := make([]PropertySource, 0)
	newPropertySources[0] = propertySource
	source.sources = append(newPropertySources, source.sources[0:]...)
}

func (source CompositePropertySource) GetPropertySources() []PropertySource {
	return source.sources
}

type PropertySources struct {
	sources []PropertySource
}

func NewPropertySources() PropertySources {
	return PropertySources{
		sources: make([]PropertySource, 0),
	}
}

func (o PropertySources) Get(name string) (PropertySource, error) {
	for _, source := range o.sources {
		if source.GetName() == name {
			return source, nil
		}
	}
	return nil, errors.New("Property not found : " + name)
}

func (o PropertySources) Add(propertySource PropertySource) {
	o.RemoveIfPresent(propertySource)
	o.sources = append(o.sources, propertySource)
}

func (o PropertySources) Remove(name string) PropertySource {
	source, index := o.findPropertySourceByName(name)
	if index != -1 {
		o.sources = append(o.sources[:index], o.sources[index+1:]...)
	}
	return source
}

func (o PropertySources) Replace(name string, propertySource PropertySource) {
	_, index := o.findPropertySourceByName(name)
	if index != -1 {
		o.sources[index] = propertySource
	}
}

func (o PropertySources) RemoveIfPresent(propertySource PropertySource) {
	if propertySource == nil {
		return
	}
	_, index := o.findPropertySourceByName(propertySource.GetName())
	if index != -1 {
		o.sources = append(o.sources[:index], o.sources[index+1:]...)
	}
}

func (o PropertySources) findPropertySourceByName(name string) (PropertySource, int) {
	for index, source := range o.sources {
		if source.GetName() == name {
			return source, index
		}
	}
	return nil, -1
}

func (o PropertySources) GetSize() int {
	return len(o.sources)
}
