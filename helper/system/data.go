package system

func GetData(data map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if _, ok := data[key]; ok {
		return data[key]
	} else {
		return defaultValue
	}
}
