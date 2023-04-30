package lark_core

type (
	fetchTenantAccessTokenRequest struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}

	fetchTenantAccessTokenResponse struct {
		CodeMsg
		TenantAccessToken string `json:"tenant_access_token"`
		// The default expired period is two hours.
		// When token's left duration is less than 30 minutes, service will return a new token.
		// At the same time, the old token can be continue used before it's expired.
		Expire int `json:"expire"`
	}
)

type CodeMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (m *CodeMsg) GetCode() int {
	return m.Code
}

func (m *CodeMsg) GetMsg() string {
	return m.Msg
}
