package lark_im

import (
	"fmt"

	lark_core "github.com/zzzzer91/lark-go/core"
)

type ImService interface {
	RelayMsgByMsgID(msgID string, msg *ImMsg) (*ImMsgResponse, error)
	SendMsgByOpenID(openID string, msg *ImMsg) (*ImMsgResponse, error)
	SendMsgByChatID(chatID string, msg *ImMsg) (*ImMsgResponse, error)
}

func NewService(cli *lark_core.Client) ImService {
	return &imServiceImpl{
		cli: cli,
	}
}

type imServiceImpl struct {
	cli *lark_core.Client
}

func (s *imServiceImpl) postImV1Message(url, receiveID string, msg *ImMsg) (*ImMsgResponse, error) {
	req := &imMsgRequest{
		ReceiveID: receiveID,
		MsgType:   msg.MsgType,
		Content:   msg.Content,
	}
	jd := new(ImMsgResponse)
	err := s.cli.PostJson(url, req, jd)
	if err != nil {
		return nil, err
	}
	return jd, nil
}

func (s *imServiceImpl) RelayMsgByMsgID(msgID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := fmt.Sprintf(urlImV1MessagesRelayTemplate, msgID)
	return s.postImV1Message(url, "", msg)
}

func (s *imServiceImpl) SendMsgByOpenID(openID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := urlImV1Messages + "?receive_id_type=" + msgReceiveIdTypeOpenID
	return s.postImV1Message(url, openID, msg)
}

func (s *imServiceImpl) SendMsgByChatID(chatID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := urlImV1Messages + "?receive_id_type=" + msgReceiveIdTypeChatID
	return s.postImV1Message(url, chatID, msg)
}
