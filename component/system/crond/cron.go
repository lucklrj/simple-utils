package crond

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tovenja/cron/v3"
)

type Crond struct {
	Timer *cron.Cron
	Bind  map[string]SingleJob //任务名:运行ID
	Sig   chan string
}
type SingleJob struct {
	JobId cron.EntryID
	Spec  string
}

func (c *Crond) Init() {
	c.Bind = map[string]SingleJob{}
	c.Sig = make(chan string, 10)

	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)
	c.Timer = cron.New(cron.WithParser(parser))
	c.Timer.Start()
}

func (c *Crond) Register(jobName, spec string, cmd func()) error {
	jobId, err := c.Timer.AddFunc(spec, cmd)
	if err != nil {
		return err
	}
	c.Bind[jobName] = SingleJob{
		JobId: jobId,
		Spec:  spec,
	}
	return nil
}
func (c *Crond) Remove(taskName string) error {
	jobInfo, ok := c.Bind[taskName]
	if !ok {
		return errors.New(taskName + "not register")
	}
	c.Timer.Remove(jobInfo.JobId)

	delete(c.Bind, taskName)
	return nil
}
func (c *Crond) Print(interval uint) {
	fmt.Println("c.Bind", c.Bind)
	sp := strings.Repeat("*", 30)
	for {
		fmt.Println(sp)
		for jobName, jobInfo := range c.Bind {
			fmt.Printf("任务名称:%s,job_id:%d,执行频率：%s\n", jobName, jobInfo.JobId, jobInfo.Spec)
		}
		fmt.Println(sp)

		time.Sleep(time.Duration(interval) * time.Second)
	}

}
