package redis

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

func newSingle(host, port string, maxIdle, maxActive, idleTimeout int) (*redigo.Pool, error) {
	pool := &redigo.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Minute,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", host+":"+port)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	return pool, nil

}
