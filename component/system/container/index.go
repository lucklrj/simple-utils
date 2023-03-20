package container

import (
	"github.com/orcaman/concurrent-map/v2"
)

type Container struct {
	data cmap.ConcurrentMap[interface{}]
}

func (c *Container) Init() {
	c.data = cmap.New[interface{}]()
}
func (c *Container) Set(key string, value interface{}) {
	c.data.Set(key, value)

}
func (c *Container) Check(key string) bool {
	return c.data.Has(key)
}
func (c *Container) Get(key string, defaultValue interface{}) interface{} {
	value, _ := c.data.Get(key)
	return value

}

func (c *Container) Delete(key string) {
	c.data.Remove(key)
}
