package date

import (
	"fmt"

	"github.com/golang-module/carbon"
	"github.com/spf13/cast"
)

func GetTodayDate() string {
	return carbon.Now().ToDateString()
}
func GetTodayUnixTime() int64 {
	return carbon.Parse(carbon.Now().ToDateString()).Timestamp()
}
func FormatHour(hour uint) string {
	hourStr := cast.ToString(hour)
	if hour <= 9 {
		return "0" + hourStr + ":00"
	} else {
		return hourStr + ":00"
	}
}

func GetDailyHourArea(format bool) []string {
	result := make([]string, 0)
	var i uint
	for i = 0; i <= 23; i++ {
		start := ""
		end := ""
		if format == true {
			start = FormatHour(i)
			end = FormatHour(i + 1)
		} else {
			start = cast.ToString(i)
			end = cast.ToString(i + 1)
		}
		result = append(result, start+"-"+end)
	}
	return result
}

func TimeStamp2Datetime(time int64) string {
	return carbon.CreateFromTimestamp(time).Format("Y-m-d H:i:s")
}

func TimeStamp2Date(time int64) string {
	return carbon.CreateFromTimestamp(time).Format("Y-m-d")
}

func Datetime2TimeStamp(datetime string) int64 {
	return carbon.Parse(datetime).Timestamp()
}

func GetNowUnixTimeStamp() uint64 {
	return uint64(carbon.Now().Timestamp())
}
func GetNowUnixTimeStampMilli() uint64 {
	return uint64(carbon.Now().TimestampMilli())
}

func GetUnixTimeStamp(datetime string, timezone string) uint64 {
	return uint64(carbon.Parse(datetime, timezone).Timestamp())
}

// TimeStamp2DateAndHour 拆分一个时间戳，返回日期，小时
func TimeStamp2DateAndHour(time uint64) (uint, uint) {
	obj := carbon.CreateFromTimestamp(int64(time))
	day := obj.Format("Y-m-d")
	return uint(carbon.Parse(day).Timestamp()), cast.ToUint(obj.Format("H"))
}

func GetNowFromTimestampMicro() string {
	nowTimestampMicro := carbon.Now().TimestampMicro()

	nowSecond := nowTimestampMicro / 1000000
	return fmt.Sprintf("%s.%d", TimeStamp2Datetime(nowSecond), nowTimestampMicro-nowSecond*1000000)

}
