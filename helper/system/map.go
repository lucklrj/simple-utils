package system

import (
	"encoding/json"
	"reflect"
	"regexp"

	"github.com/spf13/cast"
)

func Map2Struct(mapData map[string]interface{}, obj interface{}) error {
	arr, err := json.Marshal(mapData)
	if err != nil {
		return err
	}

	// 反序列化
	err = json.Unmarshal(arr, &obj)
	if err != nil {
		return err
	}
	return nil
}

// FetchByKey 从map中获取key的值，如果key不存在，则获取默认值
func FetchByKey(data interface{}, key interface{}, defaultValue interface{}) interface{} {
	dateType := reflect.TypeOf(data).String()
	if dateType[0:3] != "map" {
		return defaultValue
	}
	reg := regexp.MustCompile("map\\[([a-z]*?)[0-9]*\\].*")
	dataKeyType := reg.ReplaceAllString(dateType, "$1")
	if dataKeyType != reflect.TypeOf(key).String() {
		return defaultValue
	}

	values := reflect.ValueOf(data)
	keys := values.MapKeys()
	for _, _key := range keys {
		switch dataKeyType {
		case "uint":
			if _key.Uint() == cast.ToUint64(key) {
				return values.MapIndex(_key).Interface()
			}
		case "int":
			if _key.Int() == cast.ToInt64(key) {
				return values.MapIndex(_key).Interface()
			}
		case "string":
			if _key.String() == cast.ToString(key) {
				return values.MapIndex(_key).Interface()
			}
		case "float":
			if _key.Float() == cast.ToFloat64(key) {
				return values.MapIndex(_key).Interface()
			}
		case "bool":
			if _key.Bool() == cast.ToBool(key) {
				return values.MapIndex(_key).Interface()
			}
		}

	}

	return defaultValue
}
