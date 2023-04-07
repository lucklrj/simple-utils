package beanstalkd

import (
	"encoding/base64"
	"time"

	"github.com/beanstalkd/go-beanstalk"
	"github.com/golang/protobuf/proto"
	poolUtils "github.com/lucklrj/simple-utils/component/system/pool"
	eventPB "github.com/lucklrj/simple-utils/helper/push/beanstalkd/event"
	"github.com/mitchellh/mapstructure"
)

func Push(pool *poolUtils.Pool, businessID, eventName string, eventData map[string]interface{},
	messagePb proto.Message) error {

	err := mapstructure.Decode(eventData, messagePb)
	if err != nil {
		return err
	}

	//生成事件主体pb流
	body, err := proto.Marshal(messagePb)
	if err != nil {
		return err
	}

	worker, err := pool.Get()
	defer pool.Put(worker)
	if err != nil {
		return err
	}
	conn := worker.Handler.(*beanstalk.Conn)
	eventPb := eventPB.EventReq{}
	eventPb.BusinessID = businessID
	eventPb.Name = eventName
	eventPb.Body = base64.StdEncoding.EncodeToString(body)

	messageEvent, err := proto.Marshal(&eventPb)
	if err != nil {
		return err
	}

	//触发create_new_account事件
	_, err = conn.Put(messageEvent, 1, 0, 6*time.Hour)
	return err

}
