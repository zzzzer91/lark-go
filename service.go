package lark

import (
	"time"

	"github.com/zzzzer91/httpgo"
)

type ImV1Iface interface {
	RelayMsgByMsgID(msgID string, msg *ImMsg) (*ImMsgResponse, error)
	SendMsgByOpenID(openID string, msg *ImMsg) (*ImMsgResponse, error)
	SendMsgByChatID(chatID string, msg *ImMsg) (*ImMsgResponse, error)
}

type DocxV1Iface interface {
	GetDocxBasicInfo(documentId string) (*DocxBasicInfoResponse, error)
	GetDocxRawContent(documentId string) (*DocxRawContentResponse, error)
}

type Service interface {
	ImV1Iface
	DocxV1Iface
}

func NewService(appID, appSecret string, timeout time.Duration) Service {
	lc := &larkClient{
		cli:       httpgo.NewClient(timeout, nil),
		appID:     appID,
		appSecret: appSecret,
	}
	err := lc.refreshTenantAccessTokenRegularly()
	if err != nil {
		panic(err)
	}
	s := &serviceImpl{
		lc: lc,
	}
	return s
}
