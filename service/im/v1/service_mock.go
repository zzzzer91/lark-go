package im

import "context"

type imServiceMock struct {
}

func NewServiceMock() ImService {
	return &imServiceMock{}
}

func (m *imServiceMock) RelayMsgByMsgID(ctx context.Context, msgID string, msg *ImMsg) (*ImMsgResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (m *imServiceMock) SendMsgByOpenID(ctx context.Context, openID string, msg *ImMsg) (*ImMsgResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (m *imServiceMock) SendMsgByChatID(ctx context.Context, chatID string, msg *ImMsg) (*ImMsgResponse, error) {
	panic("not implemented") // TODO: Implement
}
