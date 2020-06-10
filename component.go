package core

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
			Logger.Panic("It supports only constructor functions")
		}
	}
}

func registerComponentType(name string, typ *Type) {
	if _, ok := componentTypes[name]; ok {
		Logger.Panic("You have already registered the same component : " + name)
	}
	componentTypes[name] = typ
}

func isSupportComponent(typ *Type) bool {
	if IsFunc(typ) {
		if typ.Typ.NumOut() > 1 || typ.Typ.NumOut() == 0 {
			Logger.Panic("Constructor functions are only supported, that why's your function must have only one return type")
		}
		retType := GetFunctionFirstReturnType(typ)
		if !IsStruct(retType) {
			Logger.Panic("Constructor functions must only return struct instances : " + retType.Typ.String())
		}
		return true
	}
	return false
}

func GetComponentTypes(typ *Type) []*Type {
	return GetComponentTypesWithParam(typ, nil)
}

func GetComponentTypesWithParam(typ *Type, paramTypes []*Type) []*Type {
	if typ == nil {
		Logger.Panic("Type must not be null")
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
	return result
}
