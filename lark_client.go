package lark

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"github.com/zzzzer91/httpgo"
)

type larkClient struct {
	cli               *httpgo.Client
	appID             string
	appSecret         string
	tenantAccessToken atomic.Value
}

// fetchTenantAccessToken get tenant access token
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token_internal
func (lc *larkClient) fetchTenantAccessToken() (*fetchTenantAccessTokenResponse, error) {
	req := &fetchTenantAccessTokenRequest{
		AppID:     lc.appID,
		AppSecret: lc.appSecret,
	}
	resp, err := lc.cli.PostJSON(urlTenantAccessToken, req)
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

func (lc *larkClient) fetchAndSetTenantAccessToken() (time.Duration, error) {
	res, err := lc.fetchTenantAccessToken()
	if err != nil {
		return 0, err
	}
	lc.tenantAccessToken.Store("Bearer " + res.TenantAccessToken)
	return time.Duration(res.Expire), nil
}

func (lc *larkClient) refreshTenantAccessTokenRegularly() error {
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
				logger.Error(err)
				t.Reset(time.Minute)
			} else {
				// refresh 20 minutes in advance
				t.Reset((expire - 20*60) * time.Second)
			}
		}
	}()
	return nil
}

func (lc *larkClient) getTenantAccessToken() string {
	return lc.tenantAccessToken.Load().(string)
}

func (lc *larkClient) GetJson(url string) (*http.Response, error) {
	return lc.cli.GetJsonWithAuth(url, lc.getTenantAccessToken())
}

func (lc *larkClient) PostJson(url string, data interface{}) (*http.Response, error) {
	return lc.cli.PostJsonWithAuth(url, data, lc.getTenantAccessToken())
}
