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

type EnumerablePropertySource interface {
	PropertySource
	GetPropertyNames() []string
}

type MapPropertySource struct {
	name   string
	source map[string]interface{}
}

func NewMapPropertySource(name string, source map[string]interface{}) MapPropertySource {
	mapPropertySource := MapPropertySource{
		name:   name,
		source: source,
	}
	return mapPropertySource
}

func (mapSource MapPropertySource) GetName() string {
	return mapSource.name
}

func (mapSource MapPropertySource) GetSource() interface{} {
	return mapSource.source
}

func (mapSource MapPropertySource) GetProperty(name string) interface{} {
	return mapSource.source[name]
}

func (mapSource MapPropertySource) ContainsProperty(name string) bool {
	return mapSource.source[name] != nil
}

func (mapSource MapPropertySource) GetPropertyNames() []string {
	return GetMapKeys(mapSource.source)
}

type CompositePropertySource struct {
	name    string
	sources []PropertySource
}

func NewCompositePropertySource(name string) CompositePropertySource {
	compositePropertySource := CompositePropertySource{
		name:    name,
		sources: make([]PropertySource, 0),
	}
	return compositePropertySource
}

func (compositeSource CompositePropertySource) GetName() string {
	return compositeSource.name
}

func (compositeSource CompositePropertySource) GetSource() interface{} {
	return compositeSource.sources
}

func (compositeSource CompositePropertySource) GetProperty(name string) interface{} {
	for _, propertySource := range compositeSource.sources {
		property := propertySource.GetProperty(name)
		if property != nil {
			return property
		}
	}
	return nil
}

func (compositeSource CompositePropertySource) ContainsProperty(name string) bool {
	for _, propertySource := range compositeSource.sources {
		if propertySource.ContainsProperty(name) {
			return true
		}
	}
	return false
}

func (compositeSource CompositePropertySource) GetPropertyNames() []string {
	names := make([]string, 0)
	for _, propertySource := range compositeSource.sources {
		if source, ok := propertySource.(EnumerablePropertySource); ok {
			names = append(names, source.GetPropertyNames()...)
		} else {
			log.Fatal("Property source does not support except EnumerablePropertySource")
		}
	}
	return names
}

func (compositeSource CompositePropertySource) AddPropertySource(propertySource PropertySource) {
	compositeSource.sources = append(compositeSource.sources, propertySource)
}

func (compositeSource CompositePropertySource) AddFirstPropertySource(propertySource PropertySource) {
	newPropertySources := make([]PropertySource, 0)
	newPropertySources[0] = propertySource
	compositeSource.sources = append(newPropertySources, compositeSource.sources[0:]...)
}

func (compositeSource CompositePropertySource) GetPropertySources() []PropertySource {
	return compositeSource.sources
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
