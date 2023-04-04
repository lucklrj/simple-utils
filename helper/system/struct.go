package system

import (
	"errors"
	"reflect"
)

func StructCopyTo(src, dst interface{}) error {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Struct {
		return errors.New("源数据格式不对，必须是struct")
	}
	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr || dstValue.Elem().Kind() != reflect.Struct {
		return errors.New("目标必须是结构体指针")
	}
	dstElem := dstValue.Elem()
	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		dstField := dstElem.FieldByName(srcValue.Type().Field(i).Name)
		if dstField.IsValid() && dstField.Type() == srcField.Type() {
			dstField.Set(srcField)
		}
	}
	return nil
}
