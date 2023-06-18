package im

import (
	"context"
	"fmt"

	lark_core "github.com/zzzzer91/lark-go/core"
)

type ImService interface {
	RelayMsgByMsgID(ctx context.Context, msgID string, msg *ImMsg) (*ImMsgResponse, error)
	SendMsgByOpenID(ctx context.Context, openID string, msg *ImMsg) (*ImMsgResponse, error)
	SendMsgByChatID(ctx context.Context, chatID string, msg *ImMsg) (*ImMsgResponse, error)
}

func NewService(cli *lark_core.Client) ImService {
	return &imServiceImpl{
		cli: cli,
	}
}

type imServiceImpl struct {
	cli *lark_core.Client
}

func (s *imServiceImpl) postImV1Message(ctx context.Context, url, receiveID string, msg *ImMsg) (*ImMsgResponse, error) {
	req := &imMsgRequest{
		ReceiveID: receiveID,
		MsgType:   msg.MsgType,
		Content:   msg.Content,
	}
	jd := new(ImMsgResponse)
	err := s.cli.PostJson(ctx, url, req, jd)
	if err != nil {
		return nil, err
	}
	return jd, nil
}

func (s *imServiceImpl) RelayMsgByMsgID(ctx context.Context, msgID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := fmt.Sprintf(urlImV1MessagesRelayTemplate, msgID)
	return s.postImV1Message(ctx, url, "", msg)
}

func (s *imServiceImpl) SendMsgByOpenID(ctx context.Context, openID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := urlImV1Messages + "?receive_id_type=" + msgReceiveIdTypeOpenID
	return s.postImV1Message(ctx, url, openID, msg)
}

func (s *imServiceImpl) SendMsgByChatID(ctx context.Context, chatID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := urlImV1Messages + "?receive_id_type=" + msgReceiveIdTypeChatID
	return s.postImV1Message(ctx, url, chatID, msg)
}
