package arrays

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/spf13/cast"
)

func InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func RemoveDuplication[T string | uint](arr []T) []T {
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

func RandChoice(source []interface{}) interface{} {
	rand.Seed(time.Now().Unix())
	return source[rand.Intn(len(source))]
}
