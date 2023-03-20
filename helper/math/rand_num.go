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
