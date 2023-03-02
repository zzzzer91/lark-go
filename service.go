package lark

type Service interface {
	RelayMsgByMsgID(msgID string, msg *ImV1Msg) (*ImV1MessageResponse, error)
	SendMsgByOpenID(openID string, msg *ImV1Msg) (*ImV1MessageResponse, error)
	SendMsgByChatID(chatID string, msg *ImV1Msg) (*ImV1MessageResponse, error)
}

func NewService(appID, appSecret string) Service {
	s := &serviceImpl{
		appID:     appID,
		appSecret: appSecret,
	}
	err := s.refreshTenantAccessTokenRegularly()
	if err != nil {
		panic(err)
	}
	return s
}
