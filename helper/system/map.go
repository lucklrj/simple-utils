package system

import (
	"encoding/json"
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

func GetMapKey[T int | string](data map[T]interface{}) []interface{} {
	keys := make([]interface{}, len(data))
	index := 0
	for key := range data {
		keys[index] = key
		index += 1
	}
	return keys
}
