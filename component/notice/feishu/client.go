package feishu

import (
	"context"

	"github.com/chyroc/lark"
)

var Client *lark.Lark

func MakeDefault(appId, appSecret string) {
	//发送飞书，要求待审
	Client = lark.New(lark.WithAppCredential(appId, appSecret))
}

func SendTextMessage(content string, receiveUserId string) (*lark.SendRawMessageResp, *lark.Response, error) {
	ctx := context.Background()
	return Client.Message.SendRawMessage(ctx, &lark.SendRawMessageReq{
		ReceiveIDType: "open_id",
		ReceiveID:     receiveUserId,
		Content:       "{\"text\":\"" + content + "\"}",
		MsgType:       "text",
	})
}
