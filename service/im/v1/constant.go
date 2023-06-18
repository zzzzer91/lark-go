package im

import lark_core "github.com/zzzzer91/lark-go/core"

const (
	urlImV1Messages              = lark_core.UrlBasic + "/im/v1/messages"          // Sending Message API
	urlImV1MessagesRelayTemplate = lark_core.UrlBasic + "/im/v1/messages/%s/reply" // Reply Message API
)

const (
	msgReceiveIdTypeOpenID  = "open_id"
	msgReceiveIdTypeUserID  = "user_id"
	msgReceiveIdTypeUnionID = "union_id"
	msgReceiveIdTypeEmail   = "email"
	msgReceiveIdTypeChatID  = "chat_id"
)
