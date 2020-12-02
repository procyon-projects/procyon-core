package core

import (
	"errors"
	"github.com/codnect/goo"
)

type Component interface{}

type ComponentProcessor interface {
	SupportsComponent(typ goo.Type) bool
	ProcessComponent(typ goo.Type) error
}

var (
	componentTypes     = make(map[string]goo.Type, 0)
	componentProcessor = make(map[string]goo.Type, 0)
)

func Register(components ...Component) {
	for _, component := range components {
		typ := goo.GetType(component)
		if isSupportComponent(typ) {
			fun := typ.ToFunctionType()
			retType := fun.GetFunctionReturnTypes()[0].ToStructType()
			compressorType := goo.GetType((*ComponentProcessor)(nil)).ToInterfaceType()
			if retType.Implements(compressorType) {
				registerComponentProcessor(typ.GetFullName(), typ)
			} else {
				registerComponentType(typ.GetFullName(), typ)
			}
		} else {
			panic("It supports only constructor functions")
		}
	}
}

func registerComponentType(name string, typ goo.Type) {
	if _, ok := componentTypes[name]; ok {
		panic("You have already registered the same component : " + name)
	}
	componentTypes[name] = typ
}

func registerComponentProcessor(name string, typ goo.Type) {
	if _, ok := componentProcessor[name]; ok {
		panic("You have already registered the same component processor : " + name)
	}
	componentProcessor[name] = typ
}

func isSupportComponent(typ goo.Type) bool {
	if typ.IsFunction() {
		fun := typ.ToFunctionType()
		if fun.GetFunctionReturnTypeCount() != 1 {
			panic("Constructor functions are only supported, that why's your function must have only one return type")
		}
		retType := fun.GetFunctionReturnTypes()[0]
		if !retType.IsStruct() {
			panic("Constructor functions must only return struct instances : " + retType.GetPackageFullName())
		}
		return true
	}
	return false
}

func GetComponentTypes(requestedType goo.Type) ([]goo.Type, error) {
	return GetComponentTypesWithParam(requestedType, nil)
}

func GetComponentTypesWithParam(requestedType goo.Type, paramTypes []goo.Type) ([]goo.Type, error) {
	if requestedType == nil {
		return nil, errors.New("type must not be null")
	}
	if !requestedType.IsStruct() && !requestedType.IsInterface() {
		panic("Requested type must be only interface or struct")
	}
	result := make([]goo.Type, 0)
	for _, componentType := range componentTypes {
		fun := componentType.ToFunctionType()
		returnType := fun.GetFunctionReturnTypes()[0].ToStructType()
		match := false
		if requestedType.IsInterface() && returnType.Implements(requestedType.ToInterfaceType()) {
			match = true
		} else if requestedType.IsStruct() {
			if requestedType.GetGoType() == returnType.GetGoType() || requestedType.ToStructType().EmbeddedStruct(returnType) {
				match = true
			}
		}
		if match && hasFunctionSameParametersWithGivenParameters(componentType, paramTypes) {
			result = append(result, componentType)
		}
	}
	return result, nil
}

func ForEachComponentType(callback func(string, goo.Type) error) (err error) {
	for componentName := range componentTypes {
		component := componentTypes[componentName]
		err = callback(componentName, component)
		if err != nil {
			break
		}
	}
	return nil
}

func ForEachComponentProcessor(callback func(string, goo.Type) error) (err error) {
	for componentProcessorName := range componentProcessor {
		componentProcessor := componentProcessor[componentProcessorName]
		err = callback(componentProcessorName, componentProcessor)
		if err != nil {
			break
		}
	}
	return
}
