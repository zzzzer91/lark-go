package lark

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"github.com/zzzzer91/httpgo"
	"github.com/zzzzer91/zlog"
)

type serviceImpl struct {
	cli               *httpgo.Client
	appID             string
	appSecret         string
	tenantAccessToken atomic.Value
}

// fetchTenantAccessToken get tenant access token
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token_internal
func (s *serviceImpl) fetchTenantAccessToken() (*fetchTenantAccessTokenResponse, error) {
	req := &fetchTenantAccessTokenRequest{
		AppID:     s.appID,
		AppSecret: s.appSecret,
	}
	resp, err := s.cli.PostJSON(urlTenantAccessToken, req)
	if err != nil {
		return nil, errors.Wrap(err, "PostJSON error")
	}
	defer resp.Body.Close()
	result := &fetchTenantAccessTokenResponse{}
	err = sonic.ConfigDefault.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal error")
	}

	return result, nil
}

func (s *serviceImpl) fetchAndSetTenantAccessToken() (time.Duration, error) {
	res, err := s.fetchTenantAccessToken()
	if err != nil {
		return 0, err
	}
	zlog.Ctx(context.Background()).Infof("fetchTenantAccessToken: %#v", res)
	s.tenantAccessToken.Store("Bearer " + res.TenantAccessToken)
	return time.Duration(res.Expire), nil
}

func (s *serviceImpl) refreshTenantAccessTokenRegularly() error {
	expire, err := s.fetchAndSetTenantAccessToken()
	if err != nil {
		return err
	}
	go func() {
		t := time.NewTimer((expire - 20*60) * time.Second)
		defer t.Stop()
		for range t.C {
			expire, err := s.fetchAndSetTenantAccessToken()
			if err != nil {
				// retry
				zlog.Ctx(context.Background()).WithError(err).Error()
				t.Reset(time.Minute)
			} else {
				// refresh 20 minutes in advance
				t.Reset((expire - 20*60) * time.Second)
			}
		}
	}()
	return nil
}

func (s *serviceImpl) getTenantAccessToken() string {
	return s.tenantAccessToken.Load().(string)
}

func (s *serviceImpl) postImV1Message(url, receiveID string, msg *ImV1Msg) (*ImV1MessageResponse, error) {
	token := s.getTenantAccessToken()
	req := &ImV1MessageRequest{
		ReceiveID: receiveID,
		MsgType:   msg.MsgType,
		Content:   msg.Content,
	}
	resp, err := s.cli.PostJsonWithAuth(url, req, token)
	if err != nil {
		return nil, errors.Wrap(err, "PostJsonWithAuth error")
	}
	defer resp.Body.Close()
	var result ImV1MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.WithStack(err)
	}
	if result.Code != 0 {
		return nil, errors.Errorf("json code is %d, msg: %s", result.Code, result.Msg)
	}
	return &result, nil
}

func (s *serviceImpl) RelayMsgByMsgID(msgID string, msg *ImV1Msg) (*ImV1MessageResponse, error) {
	url := fmt.Sprintf(urlImV1MessagesRelayTemplate, msgID)
	return s.postImV1Message(url, "", msg)
}

func (s *serviceImpl) SendMsgByOpenID(openID string, msg *ImV1Msg) (*ImV1MessageResponse, error) {
	url := urlImV1Messages + "?receive_id_type=" + msgReceiveIdTypeOpenID
	return s.postImV1Message(url, openID, msg)
}

func (s *serviceImpl) SendMsgByChatID(chatID string, msg *ImV1Msg) (*ImV1MessageResponse, error) {
	url := urlImV1Messages + "?receive_id_type=" + msgReceiveIdTypeChatID
	return s.postImV1Message(url, chatID, msg)
}
