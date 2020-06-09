package core

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"
)

type TypeConverter interface {
	Support(sourceTyp *Type, targetTyp *Type) bool
	Convert(source interface{}, sourceTyp *Type, targetTyp *Type) (interface{}, error)
}

type StringToNumberConverter struct {
}

func NewStringToNumberConverter() StringToNumberConverter {
	return StringToNumberConverter{}
}

func (converter StringToNumberConverter) Support(sourceTyp *Type, targetTyp *Type) bool {
	if sourceTyp.Val.Kind() == reflect.String {
		switch targetTyp.Val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return true
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return true
		case reflect.Float32, reflect.Float64:
			return true
		}
	}
	return false
}

func (converter StringToNumberConverter) Convert(source interface{}, sourceTyp *Type, targetTyp *Type) (interface{}, error) {
	if sourceTyp.Val.Kind() == reflect.String {
		switch targetTyp.Val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result, err := strconv.ParseInt(source.(string), 10, 64)
			if err != nil {
				return nil, nil
			}
			if targetTyp.Val.OverflowInt(result) {
				return nil, errors.New("incompatible int type " + sourceTyp.String() + " to " + targetTyp.String())
			}
			val := reflect.New(targetTyp.Typ)
			val.SetInt(result)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result, err := strconv.ParseUint(source.(string), 10, 64)
			if err != nil {
				return nil, nil
			}
			if targetTyp.Val.OverflowUint(result) {
				return nil, errors.New("incompatible uint type " + sourceTyp.String() + " to " + targetTyp.String())
			}
			val := reflect.New(targetTyp.Typ)
			val.SetUint(result)
		case reflect.Float32, reflect.Float64:
			result, err := strconv.ParseFloat(source.(string), 64)
			if err != nil {
				return nil, nil
			}
			if targetTyp.Val.OverflowFloat(result) {
				return nil, errors.New("incompatible float type " + sourceTyp.String() + " to " + targetTyp.String())
			}
			val := reflect.New(targetTyp.Typ)
			val.SetFloat(result)
		}
	}
	return nil, errors.New("unsupported type")
}

type NumberToStringConverter struct {
}

func NewNumberToStringConverter() NumberToStringConverter {
	return NumberToStringConverter{}
}

func (converter NumberToStringConverter) Support(sourceTyp *Type, targetTyp *Type) bool {
	if targetTyp.Val.Kind() == reflect.String {
		switch sourceTyp.Val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return true
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return true
		case reflect.Float32, reflect.Float64:
			return true
		}
	}
	return false
}

func (converter NumberToStringConverter) Convert(source interface{}, sourceTyp *Type, targetTyp *Type) (interface{}, error) {
	if targetTyp.Val.Kind() == reflect.String {
		switch sourceTyp.Val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return fmt.Sprintf("%d", source), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return fmt.Sprintf("%d", source), nil
		case reflect.Float32, reflect.Float64:
			return fmt.Sprintf("%f", source), nil
		}
	}
	return nil, errors.New("unsupported type")
}

type StringToBooleanConverter struct {
}

func NewStringToBooleanConverter() StringToBooleanConverter {
	return StringToBooleanConverter{}
}

func (converter StringToBooleanConverter) Support(sourceTyp *Type, targetTyp *Type) bool {
	if sourceTyp.Val.Kind() == reflect.String {
		switch targetTyp.Val.Kind() {
		case reflect.Bool:
			return true
		}
	}
	return false
}

func (converter StringToBooleanConverter) Convert(source interface{}, sourceTyp *Type, targetTyp *Type) (interface{}, error) {
	if sourceTyp.Val.Kind() == reflect.String {
		switch targetTyp.Val.Kind() {
		case reflect.Bool:
			result, err := strconv.ParseBool(source.(string))
			if err != nil {
				return nil, nil
			}
			return result, nil
		}
	}
	return nil, errors.New("unsupported type")
}

type BooleanToStringConverter struct {
}

func NewBooleanToStringConverter() BooleanToStringConverter {
	return BooleanToStringConverter{}
}

func (converter BooleanToStringConverter) Support(sourceTyp *Type, targetTyp *Type) bool {
	if targetTyp.Val.Kind() == reflect.String {
		switch sourceTyp.Val.Kind() {
		case reflect.Bool:
			return true
		}
	}
	return false
}

func (converter BooleanToStringConverter) Convert(source interface{}, sourceTyp *Type, targetTyp *Type) (interface{}, error) {
	if targetTyp.Val.Kind() == reflect.String {
		switch sourceTyp.Val.Kind() {
		case reflect.Bool:
			return strconv.FormatBool(source.(bool)), nil
		}
	}
	return nil, errors.New("unsupported type")
}

type TypeConverterRegistry interface {
	RegisterConverter(converter TypeConverter)
}

type TypeConverterService interface {
	TypeConverterRegistry
	CanConvert(source *Type, target *Type) bool
	Convert(source interface{}, sourceTyp *Type, targetTyp *Type) interface{}
}

type DefaultTypeConverterService struct {
	converters map[reflect.Type]TypeConverter
	mu         sync.RWMutex
}

func NewDefaultTypeConverterService() *DefaultTypeConverterService {
	converterService := &DefaultTypeConverterService{
		converters: make(map[reflect.Type]TypeConverter, 0),
	}
	converterService.registerDefaultConverters()
	return converterService
}

func (cs *DefaultTypeConverterService) registerDefaultConverters() {
	/* number to string and string to number */
	cs.RegisterConverter(NewNumberToStringConverter())
	cs.RegisterConverter(NewStringToNumberConverter())
	/* bool to string and string to bool */
	cs.RegisterConverter(NewBooleanToStringConverter())
	cs.RegisterConverter(NewStringToBooleanConverter())
}

func (cs *DefaultTypeConverterService) CanConvert(source *Type, target *Type) bool {
	var result bool
	cs.mu.Lock()
	for _, converter := range cs.converters {
		if converter.Support(source, target) {
			result = true
			break
		}
	}
	cs.mu.Unlock()
	return result
}

func (cs *DefaultTypeConverterService) Convert(source interface{}, sourceTyp *Type, targetTyp *Type) interface{} {
	var typConverter TypeConverter
	cs.mu.Lock()
	for _, converter := range cs.converters {
		if converter.Support(sourceTyp, targetTyp) {
			typConverter = converter
		}
	}
	cs.mu.Unlock()
	if typConverter != nil {
		defer func() {
			log.Println("converting error has just occurred")
		}()
		value, err := typConverter.Convert(source, sourceTyp, targetTyp)
		if err == nil {
			return value
		}
	}
	return nil
}

func (cs *DefaultTypeConverterService) RegisterConverter(converter TypeConverter) {
	cs.mu.Lock()
	cs.converters[GetType(converter).Typ] = converter
	cs.mu.Unlock()
}
