package beanstalkd

import (
	"encoding/base64"
	"reflect"
	"time"

	"github.com/beanstalkd/go-beanstalk"
	"github.com/golang/protobuf/proto"
	eventPB "github.com/lucklrj/simple-utils/helper/push/beanstalkd/event"
	"github.com/mitchellh/mapstructure"
)

func Push(conn *beanstalk.Conn, businessID, eventName string, eventData map[string]interface{},
	messagePb proto.Message) error {

	if eventData != nil {
		eventData["BusinessId"] = businessID

		err := mapstructure.Decode(eventData, messagePb)
		if err != nil {
			return err
		}
	} else {
		obj := reflect.ValueOf(messagePb).Elem()
		obj.FieldByName("BusinessId").SetString(businessID)
	}

	//生成事件主体pb流
	body, err := proto.Marshal(messagePb)
	if err != nil {
		return err
	}

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
