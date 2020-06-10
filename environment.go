package core

import "os"

type Environment interface {
	PropertyResolver
}

type ConfigurableEnvironment interface {
	Environment
	GetPropertySources() *PropertySources
	GetSystemEnvironment() []string
	GetTypeConverterService() TypeConverterService
}

type StandardEnvironment struct {
	propertySources  *PropertySources
	converterService TypeConverterService
	propertyResolver PropertyResolver
}

func NewStandardEnvironment() StandardEnvironment {
	env := StandardEnvironment{
		propertySources:  NewPropertySources(),
		converterService: NewDefaultTypeConverterService(),
	}
	env.propertyResolver = NewSimplePropertyResolver(env.propertySources)
	return env
}

func (env StandardEnvironment) GetPropertySources() *PropertySources {
	return env.propertySources
}

func (env StandardEnvironment) GetSystemEnvironment() []string {
	return os.Environ()
}

func (env StandardEnvironment) ContainsProperty(name string) bool {
	return env.propertyResolver.ContainsProperty(name)
}

func (env StandardEnvironment) GetProperty(name string, defaultValue string) interface{} {
	return env.propertyResolver.GetProperty(name, defaultValue)
}

func (env StandardEnvironment) GetTypeConverterService() TypeConverterService {
	return env.converterService
}
