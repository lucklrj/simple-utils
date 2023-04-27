package math

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func CreateRandNum(digit uint) string {
	return fmt.Sprintf("%0"+strconv.Itoa(int(digit))+"v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(math.Pow(10, float64(digit)))))
}
func GetRandFromArea(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

type ItemWeighted struct {
	Value    interface{}
	Weighted uint
}

func GetRandWithWeighted(items []ItemWeighted) interface{} {
	itemsLength := len(items)
	preSum := make([]uint, itemsLength)
	preSum[0] = items[0].Weighted
	for i := 1; i < itemsLength; i++ {
		preSum[i] = preSum[i-1] + items[i].Weighted
	}
	randNum := GetRandFromArea(1, int(preSum[itemsLength-1]))

	//2分法查找区间
	start, end := 0, itemsLength-1
	runNum := 0
	for start < end {
		mid := (end + start) / 2
		if preSum[mid] >= uint(randNum) {
			end = mid
		} else {
			start = mid + 1
		}
		runNum++
		if runNum > 5 {
			break
		}
	}
	return items[end].Value
}

// 测试GetRandWithWeighted的分布情况
//func test(){
//	items := []ItemWeighted{
//		{Value: 1, Weighted: 1},
//		{Value: 2, Weighted: 2},
//		{Value: 3, Weighted: 5},
//		{Value: 4, Weighted: 10},
//		{Value: 5, Weighted: 20},
//	}
//	result := map[int]uint{}
//
//	max := 10000000
//	run := 0
//	for {
//		num := GetRandWithWeighted(items).(int)
//		if _, ok := result[num]; ok {
//			result[num]++
//		} else {
//			result[num] = 1
//		}
//		run++
//		if run > max {
//			break
//		}
//	}
//
//	for key, val := range result {
//		fmt.Println(key, " = ", val)
//	}
//}
