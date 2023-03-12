package lark

import (
	"time"

	"github.com/zzzzer91/httpgo"
)

type Service interface {
	RelayMsgByMsgID(msgID string, msg *ImV1Msg) (*ImV1MessageResponse, error)
	SendMsgByOpenID(openID string, msg *ImV1Msg) (*ImV1MessageResponse, error)
	SendMsgByChatID(chatID string, msg *ImV1Msg) (*ImV1MessageResponse, error)
}

func NewService(appID, appSecret string, timeout time.Duration) Service {
	s := &serviceImpl{
		cli:       httpgo.NewClient(timeout, nil),
		appID:     appID,
		appSecret: appSecret,
	}
	err := s.refreshTenantAccessTokenRegularly()
	if err != nil {
		panic(err)
	}
	return s
}
