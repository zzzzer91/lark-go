package im

import lark_core "github.com/zzzzer91/lark-go/core"

const (
	MsgTypeText      = "text"
	MsgTypeRichText  = "rich_text"
	MsgTypeCard      = "interactive"
	MsgTypePost      = "post"
	MsgTypeImage     = "image"
	MsgTypeShareCard = "share_chat"
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
		lark_core.CodeMsg
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
