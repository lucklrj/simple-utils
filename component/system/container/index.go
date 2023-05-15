package container

import (
	dateHelper "github.com/lucklrj/simple-utils/helper/date"
	"github.com/orcaman/concurrent-map/v2"
)

type Container struct {
	data cmap.ConcurrentMap[interface{}]
}
type singleData struct {
	ExpireTime uint64
	Data       interface{}
}

func (c *Container) Init() {
	c.data = cmap.New[interface{}]()
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
	c.data.Set(key, data)
}
func (c *Container) Check(key string) bool {
	if c.data.Has(key) == false {
		return false
	} else {
		obj, _ := c.data.Get(key)
		value := obj.(singleData)
		if value.ExpireTime > 0 && value.ExpireTime < dateHelper.GetNowUnixTimeStamp() {
			return false
		} else {
			return true
		}
	}
}
func (c *Container) Get(key string, defaultValue interface{}) interface{} {
	obj, isExists := c.data.Get(key)
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
	c.data.Items()
	c.data.Remove(key)
}
func (c *Container) Items() map[string]interface{} {
	return c.data.Items()
}
