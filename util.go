package core

import (
	"errors"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type TaskWatch struct {
	taskName  string
	startTime int
	totalTime int
}

func NewTaskWatch() *TaskWatch {
	return &TaskWatch{
		taskName: "[empty_task]",
	}
}

func NewTaskWatchWithName(taskName string) *TaskWatch {
	return &TaskWatch{
		taskName: taskName,
	}
}

func (watch *TaskWatch) Start() error {
	if watch.taskName != "" && watch.startTime != 0 {
		return errors.New("TaskWatch is already running")
	}
	watch.startTime = time.Now().Nanosecond()
	return nil
}

func (watch *TaskWatch) Stop() error {
	if watch.taskName == "" {
		return errors.New("TaskWatch is not running")
	}
	watch.totalTime = time.Now().Nanosecond() - watch.startTime
	watch.taskName = ""
	return nil
}

func (watch *TaskWatch) IsRunning() bool {
	return watch.taskName != ""
}

func (watch *TaskWatch) GetTotalTime() int {
	return watch.totalTime
}

func GetMapKeys(mapObj interface{}) []string {
	argMapKeys := reflect.ValueOf(mapObj).MapKeys()
	mapKeys := make([]string, len(argMapKeys))
	for i := 0; i < len(argMapKeys); i++ {
		mapKeys[i] = argMapKeys[i].String()
	}
	return mapKeys
}

type Type struct {
	Typ  reflect.Type
	Val  reflect.Value
	name string
}

func (typ *Type) String() string {
	return typ.name
}

func GetType(obj interface{}) *Type {
	typ := reflect.TypeOf(obj)
	if typ == nil {
		panic("Type cannot be determined.")
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	result := &Type{
		Typ: typ,
		Val: val,
	}
	var name string
	if IsFunc(result) {
		returnTypeNames := GetFunctionReturnTypeNames(result)
		if len(returnTypeNames) == 1 {
			name = returnTypeNames[0]
		}
	} else {
		name = GetTypeName(result)
	}
	result.name = name
	return result
}

func sanitizedName(str string) string {
	name := strings.ReplaceAll(str, "/", ".")
	name = strings.ReplaceAll(name, "-", ".")
	name = strings.ReplaceAll(name, "_", ".")
	return name
}

func getTypeBaseName(typ reflect.Type) string {
	name := sanitizedName(typ.PkgPath())
	if name != "" {
		name = name + "." + typ.Name()
	} else {
		name = typ.Name()
	}
	if name == "" {
		name = typ.String()
	}
	return name
}

func GetTypeName(typ *Type) string {
	if typ == nil {
		panic("it must not be null")
	}
	if typ.Typ.Kind() == reflect.Func {
		panic("Must use core.GetFunctionReturnTypeNames for functions")
	}
	return getTypeBaseName(typ.Typ)
}

func GetFunctionReturnParamCount(typ *Type) int {
	if typ == nil {
		panic("it must not be null")
	}
	if typ.Typ.Kind() != reflect.Func {
		panic("You cannot use it except function")
	}
	return typ.Typ.NumOut()
}

func GetFunctionParameterCount(typ *Type) int {
	if typ == nil {
		panic("it must not be null")
	}
	return typ.Typ.NumIn()
}

func GetFunctionReturnTypeNames(typ *Type) []string {
	if typ.Typ.Kind() != reflect.Func {
		panic("It is not function type")
	}
	typeNames := make([]string, 0)
	returnTypeCount := typ.Typ.NumOut()
	for index := 0; index < returnTypeCount; index++ {
		param := typ.Typ.Out(index)
		if param.Kind() == reflect.Ptr {
			param = param.Elem()
		}
		typeNames = append(typeNames, getTypeBaseName(param))
	}
	return typeNames
}

func GetFullFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func GetFunctionInputTypes(typ *Type) []*Type {
	if typ == nil {
		panic("it must not be null")
	}
	inputParameterCount := typ.Typ.NumIn()
	inputTypes := make([]*Type, inputParameterCount)
	for index := 0; index < inputParameterCount; index++ {
		typ := typ.Typ.In(index)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		inputType := &Type{
			Typ:  typ,
			name: getTypeBaseName(typ),
		}
		inputTypes[index] = inputType
	}
	return inputTypes
}

func GetFunctionFirstReturnType(typ *Type) *Type {
	returnType := typ.Typ.Out(0)
	if returnType.Kind() == reflect.Ptr {
		returnType = returnType.Elem()
	}
	val := reflect.ValueOf(typ.Typ)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return &Type{
		Typ:  returnType,
		Val:  val,
		name: getTypeBaseName(returnType),
	}
}

func IsPtr(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Ptr
}

func IsStruct(typ *Type) bool {
	if typ == nil {
		panic("it must not be null")
	}
	return typ.Typ.Kind() == reflect.Struct
}

func IsFunc(typ *Type) bool {
	if typ == nil {
		panic("it must not be null")
	}
	return typ.Typ.Kind() == reflect.Func
}

func IsInterface(typ *Type) bool {
	if typ == nil {
		panic("it must not be null")
	}
	return typ.Typ.Kind() == reflect.Interface
}

func GetNumField(typ *Type) int {
	if typ == nil {
		panic("it must not be null")
	}
	return typ.Typ.NumField()
}

func GetFieldTypeByIndex(typ *Type, index int) reflect.StructField {
	if typ == nil {
		panic("it must not be null")
	}
	return typ.Typ.Field(index)
}

func GetFieldValueByIndex(typ *Type, index int) reflect.Value {
	if typ == nil {
		panic("it must not be null")
	}
	return typ.Val.Field(index)
}

func GetStructFieldByIndex(typ *Type, index int) reflect.StructField {
	if typ == nil {
		panic("it must not be null")
	}
	return typ.Typ.Field(index)
}

func IsAnonymous(typ reflect.StructField) bool {
	return typ.Anonymous
}

func GetTypeFromStructField(field reflect.StructField) *Type {
	return &Type{
		Typ: field.Type,
	}
}

func IsEmbeddedStruct(parentStructType *Type, childStructType *Type) bool {
	if parentStructType == nil || childStructType == nil {
		panic("it must not be null")
	}
	childMethodNum := GetNumField(childStructType)
	for index := 0; index < childMethodNum; index++ {
		field := GetStructFieldByIndex(childStructType, index)
		fieldTyp := GetTypeFromStructField(field)
		if IsAnonymous(field) && IsStruct(fieldTyp) {
			if GetTypeName(fieldTyp) == GetTypeName(parentStructType) {
				return true
			}
			if GetNumField(fieldTyp) > 0 {
				return IsEmbeddedStruct(parentStructType, fieldTyp)
			}
		}
	}
	return false
}

func HasFunctionSameParametersWithGivenParameters(typ *Type, parameters []*Type) bool {
	functionParameterCount := GetFunctionParameterCount(typ)
	if len(parameters) != functionParameterCount {
		return false
	}
	inputTypes := GetFunctionInputTypes(typ)
	for index, inputType := range inputTypes {
		if parameters[index].Typ != inputType.Typ {
			return false
		}
	}
	return true
}
