package arrays

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cast"
)

func InArray(obj interface{}, target interface{}, ignoreType bool) bool {
	targetValue := reflect.ValueOf(target)

	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			targetKind := targetValue.Index(i).Kind().String()

			var result bool
			if ignoreType { //忽略类型
				switch {
				case targetKind == "func":
					result = fmt.Sprintf("%p", targetValue.Index(i).Interface()) == fmt.Sprintf("%p", obj)
				case targetKind == "bool":
					result = cast.ToBool(targetValue.Index(i).Interface()) == cast.ToBool(obj)
				case targetKind == "interface":
					result = cast.ToString(targetValue.Index(i).Interface()) == cast.ToString(obj)
				case targetKind == "string":
					result = cast.ToString(targetValue.Index(i).Interface()) == cast.ToString(obj)
				case strings.HasPrefix(targetKind, "uint"):
					result = cast.ToUint(targetValue.Index(i).Interface()) == cast.ToUint(obj)
				case strings.HasPrefix(targetKind, "int"):
					result = cast.ToInt(targetValue.Index(i).Interface()) == cast.ToInt(obj)
				case strings.HasPrefix(targetKind, "float"):
					result = cast.ToFloat64(targetValue.Index(i).Interface()) == cast.ToFloat64(obj)
				default: //ptr, slice, struct, array,chan, map,
					result = reflect.DeepEqual(targetValue.Index(i).Interface(), obj)
				}
				if result == true {
					return true
				}
			} else {
				if targetValue.Index(i).Interface() == obj {
					return true
				}
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func RemoveDuplication[T string | uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64](arr []T) []T {
	set := make(map[T]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func Filter[T interface{}](data []T, filter map[string]interface{}) []T {
	var result = make([]T, 0)
	if len(data) > 0 {
		for _, line := range data {
			_type := reflect.TypeOf(line)
			if _type.Kind() == reflect.Struct {
				_value := reflect.ValueOf(line)
				thisLevelMatch := true
				for filterName, filterValue := range filter {
					obj, isExists := _type.FieldByName(filterName)
					if isExists {
						index := obj.Index[0]
						switch _value.Field(index).Type().String() {
						case "string":
							if _value.Field(index).String() != cast.ToString(filterValue) {
								thisLevelMatch = false
								break
							}
						case "int":
							if _value.Field(index).Int() != cast.ToInt64(filterValue) {
								thisLevelMatch = false
								break
							}
						case "uint":
							if _value.Field(index).Uint() != cast.ToUint64(filterValue) {
								thisLevelMatch = false
								break
							}
						case "bool":
							if _value.Field(index).Bool() != cast.ToBool(filterValue) {
								thisLevelMatch = false
								break
							}
						}

					}
				}
				if thisLevelMatch {
					result = append(result, line)
				}
			}
		}
	}

	return result
}

func GetKeysFromMap(mapData interface{}) []interface{} {
	result := make([]interface{}, 0)

	data := reflect.ValueOf(mapData)
	for _, single := range data.MapKeys() {
		result = append(result, single.Interface())
	}
	return result
}

func RandChoice(source interface{}) (interface{}, error) {
	sourceValue := reflect.ValueOf(source)
	kind := reflect.TypeOf(source).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return nil, errors.New("不支持该类型")
	}
	sourceLen := sourceValue.Len()

	rand.Seed(time.Now().UnixNano())
	return sourceValue.Index(rand.Intn(sourceLen)).Interface(), nil
}

func Shuffle[T any](data []T) []T {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	return data
}
