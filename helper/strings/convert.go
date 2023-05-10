package strings

import "strings"

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

//UnderscoreToCamelCase 下划线转驼峰
func UnderscoreToCamelCase(source string, isSmall bool) string {
	if source == "" {
		return ""
	}
	result := ""
	for index, group := range strings.Split(source, "_") {
		var format func(string) string
		if index == 0 && isSmall {
			format = FirstLower
		} else {
			format = FirstUpper
		}
		result = result + format(group)
	}
	return result
}

func CamelCaseToUnderscore(source string) string {
	if source == "" {
		return ""
	}

	result := make([]byte, 0)
	sourceLength := len(source)
	for i := 0; i < sourceLength; i++ {
		d := source[i]
		if i > 0 && d >= 'A' && d <= 'Z' {
			result = append(result, []byte("_")...)
		}
		result = append(result, d)
	}
	return strings.ToLower(string(result[:]))
}
