package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

const ProcyonAppFilePropertySource = "ProcyonAppFilePropertySource"

const ProcyonAppFilePrefix = "procyon"
const ProcyonAppFilePath = "resources/"
const ProcyonAppFileSuffix = ".yaml"
const ProcyonDefaultProfile = "default"

type AppFilePropertySource struct {
	propertyMap map[string]interface{}
}

func NewAppFilePropertySource(profiles string) *AppFilePropertySource {
	if profiles == "" {
		profiles = ProcyonDefaultProfile
	}

	profileArr := strings.FieldsFunc(profiles, func(r rune) bool {
		return r == ','
	})

	filePaths := make([]string, 0)

	for _, profile := range profileArr {
		path := ProcyonAppFilePath + ProcyonAppFilePrefix

		if profile == ProcyonDefaultProfile {
			path = path + ProcyonAppFileSuffix
		} else {
			path = path + "." + strings.Trim(profile, " ") + ProcyonAppFileSuffix
		}

		filePaths = append(filePaths, path)
	}

	propertyMap, err := NewAppFileParser().Parse(filePaths)

	if err != nil {
		panic(err)
	}

	propertySource := &AppFilePropertySource{
		propertyMap: propertyMap,
	}

	return propertySource
}

func (propertySource *AppFilePropertySource) GetName() string {
	return ProcyonAppFilePropertySource
}

func (propertySource *AppFilePropertySource) GetSource() interface{} {
	return propertySource.propertyMap
}

func (propertySource *AppFilePropertySource) GetProperty(name string) interface{} {
	if propertySource.ContainsProperty(name) {
		return propertySource.propertyMap[name]
	}

	return nil
}

func (propertySource *AppFilePropertySource) ContainsProperty(name string) bool {
	if _, ok := propertySource.propertyMap[name]; ok {
		return true
	}

	return false
}

func (propertySource *AppFilePropertySource) GetPropertyNames() []string {
	keys := make([]string, 0, len(propertySource.propertyMap))

	for key, _ := range propertySource.propertyMap {
		keys = append(keys, key)
	}

	return keys
}

type AppFileParser struct {
}

func NewAppFileParser() *AppFileParser {
	return &AppFileParser{}
}

func (parser *AppFileParser) Parse(filePaths []string) (map[string]interface{}, error) {
	propertyMap := make(map[string]interface{})

	for _, filePath := range filePaths {
		resultMap, err := parser.parseFile(filePath)

		if err != nil {
			return nil, err
		}

		propertyMap = parser.mergeFlattenMap(propertyMap, resultMap)
	}

	return propertyMap, nil
}

func (parser *AppFileParser) parseFile(filePath string) (map[string]interface{}, error) {
	if _, err := os.Stat(filePath); err != nil {
		return nil, fmt.Errorf("app file does not exist : %s", filePath)
	}

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("could not open app file '%s', %s", filePath, err.Error())
	}

	propertyMap := make(map[string]interface{})

	err = yaml.Unmarshal(data, propertyMap)

	if err != nil {
		return nil, fmt.Errorf("could not read app file '%s', %s", filePath, err.Error())
	}

	flattenMap := FlatMap(propertyMap)

	return flattenMap, nil
}

func (parser *AppFileParser) mergeFlattenMap(map1, map2 map[string]interface{}) map[string]interface{} {
	for key, value := range map2 {
		map1[key] = value
	}

	return map1
}
