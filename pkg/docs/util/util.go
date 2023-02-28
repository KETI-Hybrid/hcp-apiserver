package util

import (
	"reflect"
)

func PrintStructFields(t reflect.Type, m map[string]interface{}) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		m[field.Name] = field.Type.String()

		if field.Type.Kind() == reflect.Struct {
			subM := make(map[string]interface{})
			PrintStructFields(field.Type, subM)
			m[field.Name] = subM
		}
	}
}

func DocWithReq(req interface{}, resp interface{}) (interface{}, interface{}) {
	inputTypeMap := make(map[string]interface{})
	t := reflect.TypeOf(req)
	PrintStructFields(t, inputTypeMap)

	typeMap := make(map[string]interface{})

	t2 := reflect.TypeOf(resp)
	PrintStructFields(t2, typeMap)

	return inputTypeMap, typeMap
}

func DocWithoutReq(resp interface{}) interface{} {
	typeMap := make(map[string]interface{})

	t2 := reflect.TypeOf(resp)
	PrintStructFields(t2, typeMap)

	return typeMap
}
