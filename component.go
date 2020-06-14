package core

import "errors"

type Component interface{}

var (
	componentTypes = make(map[string]*Type, 0)
)

func Register(components ...Component) {
	for _, component := range components {
		typ := GetType(component)
		if isSupportComponent(typ) {
			registerComponentType(typ.String(), typ)
		} else {
			panic("It supports only constructor functions")
		}
	}
}

func registerComponentType(name string, typ *Type) {
	if _, ok := componentTypes[name]; ok {
		panic("You have already registered the same component : " + name)
	}
	componentTypes[name] = typ
}

func isSupportComponent(typ *Type) bool {
	if IsFunc(typ) {
		if typ.Typ.NumOut() > 1 || typ.Typ.NumOut() == 0 {
			panic("Constructor functions are only supported, that why's your function must have only one return type")
		}
		retType := GetFunctionFirstReturnType(typ)
		if !IsStruct(retType) {
			panic("Constructor functions must only return struct instances : " + retType.Typ.String())
		}
		return true
	}
	return false
}

func GetComponentTypes(typ *Type) ([]*Type, error) {
	return GetComponentTypesWithParam(typ, nil)
}

func GetComponentTypesWithParam(typ *Type, paramTypes []*Type) ([]*Type, error) {
	if typ == nil {
		return nil, errors.New("type must not be null")
	}
	result := make([]*Type, 0)
	for _, componentType := range componentTypes {
		if IsFunc(componentType) {
			funcReturnType := GetFunctionFirstReturnType(componentType)
			if (IsInterface(typ) && funcReturnType.Typ.Implements(typ.Typ)) ||
				(IsStruct(typ) && (typ.Typ == funcReturnType.Typ)) ||
				(IsStruct(typ) && IsEmbeddedStruct(typ, funcReturnType)) {
				if HasFunctionSameParametersWithGivenParameters(componentType, paramTypes) {
					result = append(result, componentType)
				}
			}
		} else if IsStruct(componentType) {
			if IsStruct(typ) && (typ == componentType || IsEmbeddedStruct(typ, componentType)) {
				result = append(result, componentType)
			}
		}
	}
	return result, nil
}
