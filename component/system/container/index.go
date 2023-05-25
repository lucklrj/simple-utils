package container

import (
	"sync"

	dateHelper "github.com/lucklrj/simple-utils/helper/date"
)

type Container struct {
	data sync.Map
}
type singleData struct {
	ExpireTime uint64
	Data       interface{}
}

func (c *Container) Init() {
	c.data = sync.Map{}
}
func (c *Container) Set(key string, value interface{}, expire uint64) {
	var expireTime uint64 = 0
	if expire > 0 {
		expireTime = dateHelper.GetNowUnixTimeStamp() + expire
	}
	data := singleData{
		ExpireTime: expireTime,
		Data:       value,
	}
	c.data.Store(key, data)
}
func (c *Container) Check(key string) bool {
	obj, isExits := c.data.Load(key)

	if isExits == false {
		return false
	} else {
		value := obj.(singleData)
		if value.ExpireTime > 0 && value.ExpireTime < dateHelper.GetNowUnixTimeStamp() {
			return false
		} else {
			return true
		}
	}
}
func (c *Container) Get(key string, defaultValue interface{}) interface{} {
	obj, isExists := c.data.Load(key)
	if isExists == false {
		return defaultValue
	}
	value := obj.(singleData)
	if value.ExpireTime > 0 && value.ExpireTime < dateHelper.GetNowUnixTimeStamp() {
		return defaultValue
	} else {
		return value.Data
	}

}

func (c *Container) Delete(key string) {
	c.Delete(key)
}
func (c *Container) Items() map[any]any {
	result := map[any]any{}
	c.data.Range(func(key, value any) bool {
		result[key] = value
		return true
	})
	return result
}
