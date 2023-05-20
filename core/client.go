package lark_core

import (
	"context"
	"io"
	"sync/atomic"
	"time"

	"github.com/zzzzer91/gopkg/logx"
	"github.com/zzzzer91/httpgo"
	"github.com/zzzzer91/lark-go/common/json"
)

type Client struct {
	cli               *httpgo.Client
	appID             string
	appSecret         string
	tenantAccessToken atomic.Value
}

func NewClient(appID, appSecret string, timeout time.Duration) *Client {
	lc := &Client{
		cli:       httpgo.NewClient(timeout, nil),
		appID:     appID,
		appSecret: appSecret,
	}
	err := lc.refreshTenantAccessTokenRegularly()
	if err != nil {
		panic(err)
	}
	return lc
}

// fetchTenantAccessToken get tenant access token
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token_internal
func (lc *Client) fetchTenantAccessToken() (*fetchTenantAccessTokenResponse, error) {
	req := &fetchTenantAccessTokenRequest{
		AppID:     lc.appID,
		AppSecret: lc.appSecret,
	}
	resp, err := lc.cli.PostJSON(context.Background(), urlTenantAccessToken, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := new(fetchTenantAccessTokenResponse)
	if err := json.DecodeBody(resp.Body, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (lc *Client) fetchAndSetTenantAccessToken() (time.Duration, error) {
	res, err := lc.fetchTenantAccessToken()
	if err != nil {
		return 0, err
	}
	lc.tenantAccessToken.Store("Bearer " + res.TenantAccessToken)
	return time.Duration(res.Expire), nil
}

func (lc *Client) refreshTenantAccessTokenRegularly() error {
	expire, err := lc.fetchAndSetTenantAccessToken()
	if err != nil {
		return err
	}
	go func() {
		t := time.NewTimer((expire - 20*60) * time.Second)
		defer t.Stop()
		for range t.C {
			expire, err := lc.fetchAndSetTenantAccessToken()
			if err != nil {
				// retry
				logx.Error(err)
				t.Reset(time.Minute)
			} else {
				// refresh 20 minutes in advance
				t.Reset((expire - 20*60) * time.Second)
			}
		}
	}()
	return nil
}

func (lc *Client) getTenantAccessToken() string {
	return lc.tenantAccessToken.Load().(string)
}

func (lc *Client) GetJson(ctx context.Context, url string, out json.CodeMsgIface) error {
	resp, err := lc.cli.GetJsonWithAuth(ctx, url, lc.getTenantAccessToken())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.DecodeBody(resp.Body, out); err != nil {
		return err
	}
	return nil
}

func (lc *Client) PostJson(ctx context.Context, url string, data interface{}, out json.CodeMsgIface) error {
	resp, err := lc.cli.PostJsonWithAuth(ctx, url, lc.getTenantAccessToken(), data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.DecodeBody(resp.Body, out); err != nil {
		return err
	}
	return nil
}

func (lc *Client) Download(ctx context.Context, url string) (io.ReadCloser, error) {
	resp, err := lc.cli.GetWithAuth(ctx, url, lc.getTenantAccessToken())
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
