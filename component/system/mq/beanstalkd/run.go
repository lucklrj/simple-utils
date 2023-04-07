package beanstalkd

import (
	"github.com/beanstalkd/go-beanstalk"
	_ "github.com/beanstalkd/go-beanstalk"
	poolUtils "github.com/lucklrj/simple-utils/component/system/pool"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
)

func run(url string, port uint) (*poolUtils.Pool, error) {
	pool := new(poolUtils.Pool)
	//defer pool.Close()

	//todo 暂时统一参数，稍后可以通过配置来
	pool.MaxOpenWorkers = 20
	pool.WorkerMaxLifeTime = 600
	pool.WorkerTimeOut = 10

	pool.CreateWorker = func() interface{} {
		conn, err := beanstalk.Dial("tcp", url+":"+cast.ToString(port))
		if err != nil {
			//todo 处理：创建链接时出错
			return nil
		} else {
			return conn
		}

	}
	pool.DestroyWorker = func(c interface{}) {
		c.(*grpc.ClientConn).Close()
	}

	err := pool.Init()
	return pool, err
}
