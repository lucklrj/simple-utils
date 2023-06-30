package beanstalkd

import (
	"github.com/beanstalkd/go-beanstalk"
	_ "github.com/beanstalkd/go-beanstalk"
	poolUtils "github.com/lucklrj/simple-utils/component/system/pool"
	"github.com/spf13/cast"
)

func Run(url string, port uint, workerNamePrefix string, maxOpenWorkers, workerMaxLifeTime int) (*poolUtils.Pool,
	error) {
	pool := new(poolUtils.Pool)
	//defer pool.Close()

	//todo 暂时统一参数，稍后可以通过配置来
	pool.MaxOpenWorkers = maxOpenWorkers
	pool.WorkerMaxLifeTime = workerMaxLifeTime
	pool.WorkerTimeOut = 10
	pool.WorkerNamePrefix = workerNamePrefix
	pool.Url = url

	pool.CreateWorker = func() interface{} {
		conn, err := beanstalk.Dial("tcp", pool.Url+":"+cast.ToString(port))
		if err != nil {
			//todo 处理：创建链接时出错
			return nil
		} else {
			return conn
		}

	}
	pool.DestroyWorker = func(c interface{}) {
		c.(*beanstalk.Conn).Close()
	}

	err := pool.Init()
	return pool, err
}
