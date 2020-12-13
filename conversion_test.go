package core

import (
	"github.com/procyon-projects/goo"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
}

func testConverter(t *testing.T, converter TypeConverter, sourceType goo.Type, targetType goo.Type, value interface{}, expected interface{}) {
	assert.True(t, converter.Support(sourceType, targetType))
	assert.False(t, converter.Support(goo.GetType(testStruct{}), targetType))

	result, err := converter.Convert(value, sourceType, targetType)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, targetType.GetGoType(), goo.GetType(result).GetGoType())

	result, err = converter.Convert(value, goo.GetType(testStruct{}), targetType)
	assert.NotNil(t, err)
	assert.Equal(t, "unsupported type", err.Error())
	assert.Nil(t, result)
}

func testDefaultTypeConverterService(t *testing.T, converter TypeConverterService, sourceType goo.Type, targetType goo.Type, value interface{}, expected interface{}) {
	assert.True(t, converter.CanConvert(sourceType, targetType))
	result, err := converter.Convert(value, sourceType, targetType)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, targetType.GetGoType(), goo.GetType(result).GetGoType())
}

func TestStringToNumberConverter(t *testing.T) {
	converter := NewStringToNumberConverter()

	strValue := "41"
	sourceType := goo.GetType(strValue)

	// target type : int
	testConverter(t, converter, sourceType, goo.GetType(0), strValue, 41)
	// target type : int8
	testConverter(t, converter, sourceType, goo.GetType(int8(0)), strValue, int8(41))
	// target type : int16
	testConverter(t, converter, sourceType, goo.GetType(int16(0)), strValue, int16(41))
	// target type : int32
	testConverter(t, converter, sourceType, goo.GetType(int32(0)), strValue, int32(41))
	// target type : int64
	testConverter(t, converter, sourceType, goo.GetType(int64(0)), strValue, int64(41))

	// target type : uint
	testConverter(t, converter, sourceType, goo.GetType(uint(0)), strValue, uint(41))
	// target type : uint8
	testConverter(t, converter, sourceType, goo.GetType(uint8(0)), strValue, uint8(41))
	// target type : uint16
	testConverter(t, converter, sourceType, goo.GetType(uint16(0)), strValue, uint16(41))
	// target type : uint32
	testConverter(t, converter, sourceType, goo.GetType(uint32(0)), strValue, uint32(41))
	// target type : uint64
	testConverter(t, converter, sourceType, goo.GetType(uint64(0)), strValue, uint64(41))

	// target type : float32
	strValue = "41.5"
	testConverter(t, converter, sourceType, goo.GetType(float32(0)), strValue, float32(41.5))
	// target type : float64
	testConverter(t, converter, sourceType, goo.GetType(float64(0)), strValue, float64(41.5))
}

func TestNumberToStringConverter(t *testing.T) {
	converter := NewNumberToStringConverter()
	targetType := goo.GetType("")
	expectedValue := "41"

	// source type : int, target type : string
	testConverter(t, converter, goo.GetType(0), targetType, 41, expectedValue)
	// source type : int8, target type : string
	testConverter(t, converter, goo.GetType(0), targetType, int8(41), expectedValue)
	// source type : int16, target type : string
	testConverter(t, converter, goo.GetType(0), targetType, int16(41), expectedValue)
	// source type : int32, target type : string
	testConverter(t, converter, goo.GetType(0), targetType, int32(41), expectedValue)
	// source type : int64, target type : string
	testConverter(t, converter, goo.GetType(0), targetType, int64(41), expectedValue)

	// source type : uint, target type : string
	testConverter(t, converter, goo.GetType(uint(0)), targetType, uint(41), expectedValue)
	// source type : uint8, target type : string
	testConverter(t, converter, goo.GetType(uint(0)), targetType, uint8(41), expectedValue)
	// source type : uint16, target type : string
	testConverter(t, converter, goo.GetType(uint(0)), targetType, uint16(41), expectedValue)
	// source type : uint32, target type : string
	testConverter(t, converter, goo.GetType(uint(0)), targetType, uint32(41), expectedValue)
	// source type : uint64, target type : string
	testConverter(t, converter, goo.GetType(uint(0)), targetType, uint64(41), expectedValue)

	// target type : float32
	testConverter(t, converter, goo.GetType(float32(0)), targetType, float32(41.5), "41.500000")
	// target type : float64
	testConverter(t, converter, goo.GetType(float64(0)), targetType, float64(41.5), "41.500000")
}

func TestStringToBooleanConverter(t *testing.T) {
	converter := NewStringToBooleanConverter()
	sourceType := goo.GetType("true")
	targetType := goo.GetType(true)

	testConverter(t, converter, sourceType, targetType, "true", true)
	testConverter(t, converter, sourceType, targetType, "false", false)
}

func TestBooleanToStringConverter(t *testing.T) {
	converter := NewBooleanToStringConverter()
	sourceType := goo.GetType(true)
	targetType := goo.GetType("true")

	testConverter(t, converter, sourceType, targetType, true, "true")
	testConverter(t, converter, sourceType, targetType, false, "false")
}

func TestDefaultTypeConverterService_RegisterConverterWithNil(t *testing.T) {
	converterService := NewDefaultTypeConverterService()
	assert.Panics(t, func() {
		converterService.RegisterConverter(nil)
	})
}

func TestDefaultTypeConverterService_StringToNumberConverter(t *testing.T) {
	converterService := NewDefaultTypeConverterService()

	strValue := "41"
	sourceType := goo.GetType(strValue)

	// target type : int
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(0), strValue, 41)
	// target type : int8
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(int8(0)), strValue, int8(41))
	// target type : int16
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(int16(0)), strValue, int16(41))
	// target type : int32
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(int32(0)), strValue, int32(41))
	// target type : int64
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(int64(0)), strValue, int64(41))

	// target type : uint
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(uint(0)), strValue, uint(41))
	// target type : uint8
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(uint8(0)), strValue, uint8(41))
	// target type : uint16
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(uint16(0)), strValue, uint16(41))
	// target type : uint32
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(uint32(0)), strValue, uint32(41))
	// target type : uint64
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(uint64(0)), strValue, uint64(41))

	// target type : float32
	strValue = "41.5"
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(float32(0)), strValue, float32(41.5))
	// target type : float64
	testDefaultTypeConverterService(t, converterService, sourceType, goo.GetType(float64(0)), strValue, float64(41.5))
}

func TestDefaultTypeConverterService_NumberToStringConverter(t *testing.T) {
	converterService := NewDefaultTypeConverterService()
	targetType := goo.GetType("")
	expectedValue := "41"

	// source type : int, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(0), targetType, 41, expectedValue)
	// source type : int8, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(0), targetType, int8(41), expectedValue)
	// source type : int16, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(0), targetType, int16(41), expectedValue)
	// source type : int32, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(0), targetType, int32(41), expectedValue)
	// source type : int64, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(0), targetType, int64(41), expectedValue)

	// source type : uint, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(uint(0)), targetType, uint(41), expectedValue)
	// source type : uint8, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(uint(0)), targetType, uint8(41), expectedValue)
	// source type : uint16, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(uint(0)), targetType, uint16(41), expectedValue)
	// source type : uint32, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(uint(0)), targetType, uint32(41), expectedValue)
	// source type : uint64, target type : string
	testDefaultTypeConverterService(t, converterService, goo.GetType(uint(0)), targetType, uint64(41), expectedValue)

	// target type : float32
	testDefaultTypeConverterService(t, converterService, goo.GetType(float32(0)), targetType, float32(41.5), "41.500000")
	// target type : float64
	testDefaultTypeConverterService(t, converterService, goo.GetType(float64(0)), targetType, float64(41.5), "41.500000")
}

func TestDefaultTypeConverterService_StringToBooleanConverter(t *testing.T) {
	converterService := NewDefaultTypeConverterService()

	sourceType := goo.GetType("true")
	targetType := goo.GetType(true)

	testDefaultTypeConverterService(t, converterService, sourceType, targetType, "true", true)
	testDefaultTypeConverterService(t, converterService, sourceType, targetType, "false", false)
}

func TestDefaultTypeConverterService_BooleanToStringConverter(t *testing.T) {
	converterService := NewDefaultTypeConverterService()

	sourceType := goo.GetType(true)
	targetType := goo.GetType("true")

	testDefaultTypeConverterService(t, converterService, sourceType, targetType, true, "true")
	testDefaultTypeConverterService(t, converterService, sourceType, targetType, false, "false")
}
