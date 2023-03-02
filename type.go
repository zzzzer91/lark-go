package lark

type (
	fetchTenantAccessTokenRequest struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}

	fetchTenantAccessTokenResponse struct {
		TenantAccessToken string `json:"tenant_access_token"`
		Expire            int    `json:"expire"` // 默认两小时超时，但这个值必然大于30分钟，因为还有30分钟时，就会获取新token，同时老token继续生效
	}

	// 使用顺序：chat_id > open_id > user_id > email，只会用一个
	ImV1MessageRequest struct {
		ReceiveID string `json:"receive_id,omitempty"` // 依据receive_id_type的值，填写对应的消息接收者id
		MsgType   string `json:"msg_type"`             // 参考：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json
		Content   string `json:"content"`              // 消息内容，json结构，同上
	}

	ImV1MessageResponse struct {
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
	// EventSubscriptionMessage 接受 lark 机器人消息回调
	// 字段含义见：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
	EventSubscriptionMessage struct {
		Type      string `json:"type"` // 用于判断 challenge
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
