package core

import (
	"os"
	"strings"
)

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

const ProcyonSystemEnvironmentPropertySource = "ProcyonSystemEnvironmentPropertySource"

type SystemEnvironmentPropertySource struct {
	environmentProperties map[string]string
}

func NewSystemEnvironmentPropertySource() SystemEnvironmentPropertySource {
	propertySource := SystemEnvironmentPropertySource{
		environmentProperties: make(map[string]string, 0),
	}

	environmentProperties := os.Environ()
	for _, property := range environmentProperties {
		index := strings.Index(property, "=")
		if index != -1 {
			propertySource.environmentProperties[property[:index]] = property[index+1:]
		}
	}
	return propertySource
}

func (propertySource SystemEnvironmentPropertySource) GetName() string {
	return ProcyonSystemEnvironmentPropertySource
}

func (propertySource SystemEnvironmentPropertySource) GetSource() interface{} {
	return propertySource.environmentProperties
}

func (propertySource SystemEnvironmentPropertySource) GetProperty(name string) interface{} {
	actualPropertyName := propertySource.checkPropertyName(strings.ToLower(name))
	if actualPropertyName != nil {
		return propertySource.environmentProperties[actualPropertyName.(string)]
	}

	actualPropertyName = propertySource.checkPropertyName(strings.ToUpper(name))
	if actualPropertyName != nil {
		return propertySource.environmentProperties[actualPropertyName.(string)]
	}
	return nil
}

func (propertySource SystemEnvironmentPropertySource) ContainsProperty(name string) bool {
	return propertySource.checkPropertyName(strings.ToUpper(name)) != nil || propertySource.checkPropertyName(strings.ToLower(name)) != nil
}

func (propertySource SystemEnvironmentPropertySource) GetPropertyNames() []string {
	keys := make([]string, 0, len(propertySource.environmentProperties))
	for key, _ := range propertySource.environmentProperties {
		keys = append(keys, key)
	}
	return keys
}

func (propertySource SystemEnvironmentPropertySource) checkIfPresent(propertyName string) bool {
	if _, ok := propertySource.environmentProperties[propertyName]; ok {
		return true
	}
	return false
}

func (propertySource SystemEnvironmentPropertySource) checkPropertyName(propertyName string) interface{} {
	if propertySource.checkIfPresent(propertyName) {
		return propertyName
	}

	noHyphenPropertyName := strings.ReplaceAll(propertyName, "-", "_")
	if propertyName != noHyphenPropertyName && propertySource.checkIfPresent(noHyphenPropertyName) {
		return noHyphenPropertyName
	}

	noDotPropertyName := strings.ReplaceAll(propertyName, ".", "_")
	if propertyName != noDotPropertyName && propertySource.checkIfPresent(noDotPropertyName) {
		return noDotPropertyName
	}

	noHyphenAndNoDotName := strings.ReplaceAll(noDotPropertyName, "-", "_")
	if noDotPropertyName != noHyphenAndNoDotName && propertySource.checkIfPresent(noHyphenAndNoDotName) {
		return noHyphenAndNoDotName
	}

	return nil
}
