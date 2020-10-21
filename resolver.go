package core

type PropertyResolver interface {
	ContainsProperty(name string) bool
	GetProperty(name string, defaultValue string) interface{}
}

type SimplePropertyResolver struct {
	sources *PropertySources
}

func NewSimplePropertyResolver(sources *PropertySources) *SimplePropertyResolver {
	return &SimplePropertyResolver{
		sources,
	}
}

func (resolver *SimplePropertyResolver) ContainsProperty(name string) bool {
	for _, propertySource := range resolver.sources.GetPropertyResources() {
		if propertySource.ContainsProperty(name) {
			return true
		}
	}
	return false
}

func (resolver *SimplePropertyResolver) GetProperty(name string, defaultValue string) interface{} {
	for _, propertySource := range resolver.sources.GetPropertyResources() {
		if propertySource.ContainsProperty(name) {
			return propertySource.GetProperty(name).(string)
		}
	}
	return defaultValue
}
