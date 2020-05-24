package core

import "log"

type Component interface{}

const componentFunctionSeparator = "#"
const componentStructSeparator = "$"

var (
	componentTypes = make(map[string]*Type, 0)
)

func Register(components ...Component) {
	for _, component := range components {
		typ := GetType(component)
		if isSupportComponent(typ) {
			name := getComponentName(component)
			registerComponentType(name, typ)
		} else {
			log.Fatal("It supports only struct and function")
		}
	}
}

func registerComponentType(name string, typ *Type) {
	if isFunc(typ) && (typ.Typ.NumOut() > 1 || typ.Typ.NumOut() == 0) {
		log.Fatal("Constructor functions are only supported, that why's your function must have only one return type : " + name)
	}
	if _, ok := componentTypes[name]; ok {
		log.Fatal("You have already registered the same component : " + name)
	}
	componentTypes[name] = typ
}

func isSupportComponent(typ *Type) bool {
	if isFunc(typ) {
		retType := getFuncReturnType(typ)
		if !isStruct(retType) {
			log.Fatal("Constructor functions must only return struct instances : " + retType.Typ.String())
		}
		return true
	}
	return isStruct(typ)
}

func getComponentName(component Component) string {
	typ := GetType(component)
	var name string
	if isStruct(typ) {
		name = getStructName(typ)
	} else {
		name = getFunctionName(component)
	}
	return name
}

func GetComponentTypes(typ *Type) []*Type {
	if typ == nil {
		log.Fatal("Type must not be null")
	}
	result := make([]*Type, 0)
	for key, componentType := range componentTypes {
		log.Printf(key)
		if isFunc(componentType) {
			funcReturnType := getFuncReturnType(componentType)
			funcReturnType.Typ.ConvertibleTo(typ.Typ)
			if isInterface(typ) && funcReturnType.Typ.Implements(typ.Typ) {
				result = append(result, componentType)
			} else if isStruct(typ) && (typ.Typ == funcReturnType.Typ || isEmbeddedStruct(typ, funcReturnType)) {
				result = append(result, componentType)
			}
		} else if isStruct(componentType) {
			if isStruct(typ) && (typ == componentType || isEmbeddedStruct(typ, componentType)) {
				result = append(result, componentType)
			}
		}
	}
	return result
}
