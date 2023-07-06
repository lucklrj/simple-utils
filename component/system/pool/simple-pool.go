package pool

import (
	"errors"
	"sync"
	"time"

	dateHelper "github.com/lucklrj/simple-utils/helper/date"
	"github.com/lucklrj/simple-utils/helper/strings"
)

type Worker struct {
	Name     string
	Handler  interface{}
	LeftTime int64
}

type Pool struct {
	MaxOpenWorkers    int
	WorkerMaxLifeTime int
	WorkerTimeOut     int
	WorkNum           int
	mu                sync.RWMutex
	Url               string

	CreateWorker     func() interface{}
	DestroyWorker    func(interface{})
	IdleWorkers      chan *Worker
	WorkerNamePrefix string
}

//idle
func (p *Pool) create() {
	defer func() { p.WorkNum++ }()

	handler := p.CreateWorker()
	if handler == nil {
		return
	}
	var leftTime int64 = 0
	if p.WorkerMaxLifeTime > 0 {
		leftTime = int64(dateHelper.GetNowUnixTimeStamp()) + int64(p.WorkerMaxLifeTime)
	}

	Worker := new(Worker)
	Worker.Handler = handler
	Worker.LeftTime = leftTime
	Worker.Name = p.WorkerNamePrefix + "_" + strings.GetRand(6)
	p.mu.RLock()
	p.IdleWorkers <- Worker
	p.mu.RUnlock()
}

func (p *Pool) Init() error {
	p.WorkNum = 0
	if p.MaxOpenWorkers > 0 {
		p.IdleWorkers = make(chan *Worker, p.MaxOpenWorkers)
		return nil
	} else {
		return errors.New("parameter:MaxOpenWorkers must be greater than zero.")
	}
}
func (p *Pool) Get() (*Worker, error) {
	//线程池=0，但未达到max，直接创建
	//if len(p.IdleWorkers) == 0 && p.WorkNum < p.MaxOpenWorkers {
	//当死亡worker未剔除时，此时IdleWorkers不为空
	if p.WorkNum < p.MaxOpenWorkers {
		p.create()
	}
	for {
		select {
		//case <-time.After(time.Duration(p.WorkerTimeOut / 1000)):
		//	if len(p.IdleWorkers) < p.MaxOpenWorkers {
		//		p.Create()
		//	}
		case <-time.After(time.Duration(p.WorkerTimeOut) * time.Second):
			return nil, errors.New("get Workerection time out")
		case Worker := <-p.IdleWorkers:
			if Worker.LeftTime > 0 && Worker.LeftTime < int64(dateHelper.GetNowUnixTimeStamp()) {
				p.DestroyWorker((*Worker).Handler)
				p.WorkNum--
				continue
			} else {
				return Worker, nil
			}
		}
	}
}
func (p *Pool) Put(c *Worker) {
	if len(p.IdleWorkers) > p.MaxOpenWorkers {
		p.DestroyWorker((*c).Handler)
	} else {
		p.mu.RLock()
		p.IdleWorkers <- c
		p.mu.RUnlock()
	}
}
func (p *Pool) Close() {
	p.mu.Lock()
	idleWorkers := p.IdleWorkers
	p.IdleWorkers = nil
	p.mu.Unlock()

	if idleWorkers != nil {
		close(idleWorkers)
		for Worker := range idleWorkers {
			p.DestroyWorker((*Worker).Handler)
		}
	}
}
