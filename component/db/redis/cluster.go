package redis

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/lucklrj/simple-utils/component/system/env"
	"github.com/mna/redisc"
	"github.com/spf13/cast"
	"overseas/helper/output"
)

func newCluster(host, port string) *redisc.Cluster {

	cluster := &redisc.Cluster{
		StartupNodes: []string{host + ":" + port},
		DialOptions:  []redis.DialOption{redis.DialConnectTimeout(5 * time.Second)},
		CreatePool:   createPool,
	}
	//defer cluster.Close()
	if err := cluster.Refresh(); err != nil {
		output.Error("Refresh failed: " + err.Error())
		os.Exit(0)
	}

	return cluster
}
func createPool(addr string,maxIdle,maxActive,idleTimeout int, opts ...redis.DialOption) (*redis.Pool, error) {

	//maxIdle := cast.ToInt(env.Envs["redis_max_idle"])
	//maxActive := cast.ToInt(env.Envs["redis_max_active"])
	//idleTimeout := cast.ToInt(env.Envs["redis_max_idle_timeout"])
	//return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, opts...)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}
