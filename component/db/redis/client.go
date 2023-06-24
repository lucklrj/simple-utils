package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/lucklrj/simple-utils/component/system/env"
	"github.com/mna/redisc"
	"github.com/spf13/cast"
	"overseas/helper/output"
)

var DefaultConn redisClient

type redisClient struct {
	obj       interface{}
	isCluster bool
}

func (r *redisClient) Get(key string) (string, error) {
	link, err := r.getLink()
	if err != nil {
		return "", err
	}
	defer link.Close()
	return redigo.String(link.Do("GET", key))

}

func (r *redisClient) Set(key string, value string, expire int64) error {
	link, err := r.getLink()
	if err != nil {
		return err
	}
	//defer link.Close()
	_, err = link.Do("SET", key, value)
	if expire > 0 {
		_, err = link.Do("EXPIRE", key, expire)
	}
	return err
}

func (r *redisClient) Delete(key string) error {
	link, err := r.getLink()
	if err != nil {
		return err
	}
	defer link.Close()
	_, err = link.Do("Del", key)

	return err
}
func (r *redisClient) Expire(key string, expire int64) error {
	link, err := r.getLink()
	if err != nil {
		return err
	}
	defer link.Close()
	_, err = link.Do("EXPIRE", key, expire)
	return err
}
func (r *redisClient) Incr(key string) (interface{}, error) {
	link, err := r.getLink()
	if err != nil {
		return "", err
	}
	defer link.Close()
	return link.Do("INCR", key)
}
func (r *redisClient) getLink() (redis.Conn, error) {
	var err error
	var link redis.Conn
	if r.isCluster == true {
		rc, _ := r.obj.(*redisc.Cluster).Dial()
		link, err = redisc.RetryConn(rc, 10, 100*time.Millisecond)
		if err != nil {
			output.Info("RetryConn failed:" + err.Error())
			return nil, err
		}
	} else {
		link = r.obj.(*redigo.Pool).Get()
	}
	return link, err
}
func MakeDefault() {
	isCluster := cast.ToBool(env.Envs["redis_cluster"])
	if isCluster == true {
		DefaultConn = redisClient{
			obj:       newCluster(),
			isCluster: true,
		}
	} else {
		DefaultConn = redisClient{
			obj:       newSingle(),
			isCluster: false,
		}
	}
}
