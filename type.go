package lark

type (
	// refer: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
	EventSubscriptionMessage struct {
		Type      string `json:"type"`
		Challenge string `json:"challenge"`
		Schema    string `json:"schema"`
		Header    struct {
			EventID    string `json:"event_id"`
			EventType  string `json:"event_type"`
			CreateTime string `json:"create_time"`
			Token      string `json:"token"`
			AppID      string `json:"app_id"`
			TenantKey  string `json:"tenant_key"`
		} `json:"header"`
		Event struct {
			Sender struct {
				SenderID struct {
					UnionID string `json:"union_id"`
					UserID  string `json:"user_id"`
					OpenID  string `json:"open_id"`
				} `json:"sender_id"`
				SenderType string `json:"sender_type"`
				TenantKey  string `json:"tenant_key"`
			} `json:"sender"`
			Message struct {
				MessageID   string `json:"message_id"`
				RootID      string `json:"root_id"`
				ParentID    string `json:"parent_id"`
				CreateTime  string `json:"create_time"`
				ChatID      string `json:"chat_id"`
				ChatType    string `json:"chat_type"`
				MessageType string `json:"message_type"`
				Content     string `json:"content"`
				Mentions    []struct {
					Key string `json:"key"`
					ID  struct {
						UnionID string `json:"union_id"`
						UserID  string `json:"user_id"`
						OpenID  string `json:"open_id"`
					} `json:"id"`
					Name      string `json:"name"`
					TenantKey string `json:"tenant_key"`
				} `json:"mentions"`
			} `json:"message"`
		} `json:"event"`
	}

	EventSubscriptionMessageChallengeResponse struct {
		Challenge string `json:"challenge"`
	}
)

type (
	fetchTenantAccessTokenRequest struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}

	fetchTenantAccessTokenResponse struct {
		TenantAccessToken string `json:"tenant_access_token"`
		// The default expired period is two hours.
		// When token's left duration is less than 30 minutes, service will return a new token.
		// At the same time, the old token can be continue used before it's expired.
		Expire int `json:"expire"`
	}
)

type (
	// refer: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json
	imMsgRequest struct {
		ReceiveID string `json:"receive_id,omitempty"`
		MsgType   string `json:"msg_type"`
		Content   string `json:"content"`
	}

	ImMsg struct {
		Content string
		MsgType string
	}

	ImMsgResponse struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			MessageID  string `json:"message_id"`
			RootID     string `json:"root_id"`
			ParentID   string `json:"parent_id"`
			MsgType    string `json:"msg_type"`
			CreateTime string `json:"create_time"`
			UpdateTime string `json:"update_time"`
			Deleted    bool   `json:"deleted"`
			Updated    bool   `json:"updated"`
			ChatID     string `json:"chat_id"`
			Sender     struct {
				ID         string `json:"id"`
				IDType     string `json:"id_type"`
				SenderType string `json:"sender_type"`
				TenantKey  string `json:"tenant_key"`
			} `json:"sender"`
			Body struct {
				Content string `json:"content"`
			} `json:"body"`
			Mentions []struct {
				Key       string `json:"key"`
				ID        string `json:"id"`
				IDType    string `json:"id_type"`
				Name      string `json:"name"`
				TenantKey string `json:"tenant_key"`
			} `json:"mentions"`
			UpperMessageID string `json:"upper_message_id"`
		}
	}
)

type (
	DocxBasicInfoResponse struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Document struct {
				DocumentID string `json:"document_id"`
				RevisionID int    `json:"revision_id"`
				Title      string `json:"title"`
			} `json:"document"`
		} `json:"data"`
	}

	DocxRawContentResponse struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Content string `json:"content"`
		} `json:"data"`
	}
)
