package core

type PropertySource interface {
	GetName() string
	GetSource() interface{}
	GetProperty(name string) interface{}
	ContainsProperty(name string) bool
	GetPropertyNames() []string
}

type PropertySources struct {
	sources []PropertySource
}

func NewPropertySources() *PropertySources {
	return &PropertySources{
		sources: make([]PropertySource, 0),
	}
}

func (o *PropertySources) Get(name string) (PropertySource, bool) {
	for _, source := range o.sources {
		if source.GetName() == name {
			return source, true
		}
	}
	return nil, false
}

func (o *PropertySources) Add(propertySource PropertySource) {
	o.RemoveIfPresent(propertySource)
	o.sources = append(o.sources, propertySource)
}

func (o *PropertySources) Remove(name string) PropertySource {
	source, index := o.findPropertySourceByName(name)
	if index != -1 {
		o.sources = append(o.sources[:index], o.sources[index+1:]...)
	}
	return source
}

func (o *PropertySources) Replace(name string, propertySource PropertySource) {
	_, index := o.findPropertySourceByName(name)
	if index != -1 {
		o.sources[index] = propertySource
	}
}

func (o *PropertySources) RemoveIfPresent(propertySource PropertySource) {
	if propertySource == nil {
		return
	}
	_, index := o.findPropertySourceByName(propertySource.GetName())
	if index != -1 {
		o.sources = append(o.sources[:index], o.sources[index+1:]...)
	}
}

func (o *PropertySources) findPropertySourceByName(name string) (PropertySource, int) {
	for index, source := range o.sources {
		if source.GetName() == name {
			return source, index
		}
	}
	return nil, -1
}

func (o *PropertySources) GetSize() int {
	return len(o.sources)
}

func (o *PropertySources) GetPropertyResources() []PropertySource {
	return o.sources
}
