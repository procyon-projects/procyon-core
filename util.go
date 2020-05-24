package core

import (
	"errors"
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type TaskWatch struct {
	taskName  string
	startTime int64
	totalTime int64
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
	watch.startTime = time.Now().Unix()
	return nil
}

func (watch *TaskWatch) Stop() error {
	if watch.taskName == "" {
		return errors.New("TaskWatch is not running")
	}
	watch.totalTime = time.Now().Unix() - watch.startTime
	watch.taskName = ""
	return nil
}

func (watch *TaskWatch) IsRunning() bool {
	return watch.taskName != ""
}

func (watch *TaskWatch) GetTotalTime() int64 {
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
	Typ reflect.Type
	Val reflect.Value
}

func sanitizedName(str string) string {
	name := strings.ReplaceAll(str, "/", ".")
	name = strings.ReplaceAll(name, "-", ".")
	name = strings.ReplaceAll(name, "_", ".")
	return name
}

func getStructName(typ Type) string {
	if isStruct(typ) {
		name := sanitizedName(typ.Typ.PkgPath())
		name = name + componentStructSeparator + typ.Typ.Name()
		return name
	}
	return "<nil>"
}

func getFunctionName(component Component) string {
	funcFullName := getFullFunctionName(component)
	lastIndexForDot := strings.LastIndex(funcFullName, ".")
	funcFullName = funcFullName[0:lastIndexForDot] + componentFunctionSeparator + funcFullName[lastIndexForDot+1:]
	name := sanitizedName(funcFullName)
	return name
}

func GetType(component Component) Type {
	typ := reflect.TypeOf(component)
	if typ == nil {
		log.Fatal("Type cannot be determined.")
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	val := reflect.ValueOf(component)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return Type{
		Typ: typ,
		Val: val,
	}
}

func getFullFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func getFuncReturnType(typ Type) Type {
	returnType := typ.Typ.Out(0)
	if returnType.Kind() == reflect.Ptr {
		returnType = returnType.Elem()
	}
	val := reflect.ValueOf(typ.Typ)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return Type{
		Typ: returnType,
		Val: val,
	}
}

func isStruct(typ Type) bool {
	return typ.Typ.Kind() == reflect.Struct
}

func isFunc(typ Type) bool {
	return typ.Typ.Kind() == reflect.Func
}

func isInterface(typ Type) bool {
	return typ.Typ.Kind() == reflect.Interface
}

func getNumField(typ Type) int {
	return typ.Typ.NumField()
}

func getFieldByIndex(typ Type, index int) reflect.StructField {
	return typ.Typ.Field(index)
}

func isAnonymous(typ reflect.StructField) bool {
	return typ.Anonymous
}

func getTypeFromStructField(field reflect.StructField) Type {
	return Type{
		Typ: field.Type,
	}
}

func isEmbeddedStruct(parentStructType Type, childStructType Type) bool {
	childMethodNum := getNumField(childStructType)
	for index := 0; index < childMethodNum; index++ {
		field := getFieldByIndex(childStructType, index)
		fieldTyp := getTypeFromStructField(field)
		if isAnonymous(field) && isStruct(fieldTyp) {
			if getStructName(fieldTyp) == getStructName(parentStructType) {
				return true
			}
			if getNumField(fieldTyp) > 0 {
				return isEmbeddedStruct(parentStructType, fieldTyp)
			}
		}
	}
	return false
}
