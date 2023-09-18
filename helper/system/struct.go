package system

import (
	"errors"
	"reflect"

	"github.com/spf13/cast"
)

func StructCopyTo(src, dst interface{}, ignoreZeroValue bool) error {
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

		if dstField.IsValid() {
			if dstField.Type() == srcField.Type() {
				if ignoreZeroValue == true {
					dstField.Set(srcField)
				} else {
					if reflect.ValueOf(srcField.Interface()).IsZero() == false {
						dstField.Set(srcField)
					}
				}
			} else {
				value, err := changeType(dstField.Type().String(), srcField.Interface())
				if err == nil {
					tmp := reflect.ValueOf(value)
					if ignoreZeroValue == true {
						dstField.Set(tmp)
					} else {
						if reflect.ValueOf(tmp.Interface()).IsZero() == false {
							dstField.Set(tmp)
						}
					}
				}
			}
		}

	}
	return nil
}

func changeType(targetType string, value interface{}) (interface{}, error) {
	switch targetType {
	case "string":
		return cast.ToString(value), nil
	case "bool":
		return cast.ToBool(value), nil
	case "int":
		return cast.ToInt(value), nil
	case "int8":
		return cast.ToInt8(value), nil
	case "int16":
		return cast.ToInt16(value), nil
	case "int32":
		return cast.ToInt32(value), nil
	case "int64":
		return cast.ToInt64(value), nil
	case "uint":
		return cast.ToUint(value), nil
	case "uint8":
		return cast.ToUint8(value), nil
	case "uint16":
		return cast.ToUint16(value), nil
	case "uint32":
		return cast.ToUint32(value), nil
	case "uint64":
		return cast.ToUint64(value), nil
	case "float32":
		return cast.ToFloat32(value), nil
	case "float64":
		return cast.ToFloat64(value), nil
	default:
		return nil, errors.New("不支持该类型的转换")
	}
}
