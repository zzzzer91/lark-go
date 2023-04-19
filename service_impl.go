package lark

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type serviceImpl struct {
	lc *larkClient
}

func (s *serviceImpl) postImV1Message(url, receiveID string, msg *ImMsg) (*ImMsgResponse, error) {
	req := &imMsgRequest{
		ReceiveID: receiveID,
		MsgType:   msg.MsgType,
		Content:   msg.Content,
	}
	resp, err := s.lc.PostJson(url, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result ImMsgResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.WithStack(err)
	}
	if result.Code != 0 {
		return nil, errors.Errorf("json code is %d, msg: %s", result.Code, result.Msg)
	}
	return &result, nil
}

func (s *serviceImpl) RelayMsgByMsgID(msgID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := fmt.Sprintf(urlImV1MessagesRelayTemplate, msgID)
	return s.postImV1Message(url, "", msg)
}

func (s *serviceImpl) SendMsgByOpenID(openID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := urlImV1Messages + "?receive_id_type=" + msgReceiveIdTypeOpenID
	return s.postImV1Message(url, openID, msg)
}

func (s *serviceImpl) SendMsgByChatID(chatID string, msg *ImMsg) (*ImMsgResponse, error) {
	url := urlImV1Messages + "?receive_id_type=" + msgReceiveIdTypeChatID
	return s.postImV1Message(url, chatID, msg)
}

func (s *serviceImpl) GetDocxBasicInfo(documentId string) (*DocxBasicInfoResponse, error) {
	url := fmt.Sprintf(urlDocxV1DocumentsTemplate, documentId)
	resp, err := s.lc.GetJson(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result DocxBasicInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.WithStack(err)
	}
	if result.Code != 0 {
		return nil, errors.Errorf("json code is %d, msg: %s", result.Code, result.Msg)
	}
	return &result, nil
}

func (s *serviceImpl) GetDocxRawContent(documentId string) (*DocxRawContentResponse, error) {
	url := fmt.Sprintf(urlDocxV1RawContentTemplate, documentId)
	resp, err := s.lc.GetJson(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result DocxRawContentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.WithStack(err)
	}
	if result.Code != 0 {
		return nil, errors.Errorf("json code is %d, msg: %s", result.Code, result.Msg)
	}
	return &result, nil
}
