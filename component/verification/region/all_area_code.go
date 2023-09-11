package region

import (
	"fmt"

	"github.com/lucklrj/simple-utils/component/system/log"
	"github.com/lucklrj/simple-utils/component/verification/data"
	"github.com/tidwall/gjson"
)

type SingleArea struct {
	EnName    string "json:en_name'"
	CnName    string "json:cn_name'"
	Code      uint   "json:code'"
	ShortName string "json:short_name'"
}

var allAreaCode = map[uint]SingleArea{}

func init() {
	//解析json
	areaCodeFile := "region/area_code.json"
	json, err := data.StaticFile.ReadFile(areaCodeFile)
	if err != nil {
		errMsg := fmt.Sprintf("加载区域文件:%s,出错:%s.", areaCodeFile, err.Error())
		log.Logger.Error(errMsg)
		return
	}
	//json整理为map
	var code uint = 0

	allAreaArray := gjson.ParseBytes(json).Array()
	for _, line := range allAreaArray {
		code = uint(line.Get("code").Uint())
		allAreaCode[code] = SingleArea{
			EnName:    line.Get("en_name").String(),
			CnName:    line.Get("cn_name").String(),
			Code:      code,
			ShortName: line.Get("short_name").String(),
		}
	}
}

func CheckAreaCode(areaCode uint) bool {
	_, ok := allAreaCode[areaCode]
	return ok
}
