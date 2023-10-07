package format

import (
	"reflect"

	"github.com/fatih/structs"
)

func Format(obj interface{}, callbacks map[uint]func(interface{}) interface{}) interface{} {
	objAddress := reflect.ValueOf(&obj)
	objValue := reflect.ValueOf(obj)
	tmp := reflect.New(objValue.Type()).Elem()

	kind := objValue.Kind()
	switch kind { //只支持简单类型修改
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.String:
		callback, ok := callbacks[uint(kind)]
		if ok {
			return callback(obj)
		}
		return obj

	case reflect.Struct:
		for _, field := range structs.Fields(obj) {
			tmp.FieldByName(field.Name()).Set(reflect.ValueOf(Format(field.Value(), callbacks)))
		}
		objAddress.Elem().Set(tmp)
	case reflect.Slice, reflect.Array:
		for i := 0; i < objValue.Len(); i++ {
			objValue.Index(i).Set(reflect.ValueOf(Format(objValue.Index(i).Interface(), callbacks)))
		}
	default:
		return obj
	}
	return obj
}
