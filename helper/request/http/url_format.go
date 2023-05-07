package http

import (
	"sort"
	"strings"

	"github.com/goware/urlx"
)

type singleQuery struct {
	key   string `json:"key"`
	value string `json:"value"`
}

type queryOrder []singleQuery

func (array queryOrder) Len() int {
	return len(array)
}
func (array queryOrder) Less(i, j int) bool {
	return array[i].key < array[j].key
}
func (array queryOrder) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

// UrlQueryOrder 对url进行整理：相同参数均为一个
func UrlQueryOrder(path string) string {
	result, _ := urlx.Parse(path)
	pathFormat := result.Scheme + "://" + result.User.String() + result.Host + result.Path
	rawQueryFormat := packQuery(splitQuery(result.RawQuery))
	if rawQueryFormat != "" {
		pathFormat = pathFormat + "?" + rawQueryFormat
	}
	return pathFormat

}

func splitQuery(query string) queryOrder {
	var result queryOrder
	if query != "" {
		queryGroup := strings.Split(query, "&")
		if len(queryGroup) > 0 {
			for _, singleGroup := range queryGroup {
				if strings.Contains(singleGroup, "=") {
					pair := strings.Split(singleGroup, "=")
					result = append(result, singleQuery{key: pair[0], value: pair[1]})
				} else {
					result = append(result, singleQuery{key: singleGroup, value: ""})
				}
			}
			sort.Sort(result)
		}

	}
	return result
}
func packQuery(query queryOrder) string {
	pairs := make([]string, 0)
	for _, singleQuery := range query {
		pairs = append(pairs, singleQuery.key+"="+singleQuery.value)
	}
	return strings.Join(pairs, "&")
}
