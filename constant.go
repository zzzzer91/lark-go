package lark

// lark open api URL
const (
	urlBasic                     = "https://open.feishu.cn/open-apis"
	urlTenantAccessToken         = urlBasic + "/auth/v3/tenant_access_token/internal" // 获取token
	urlImV1Messages              = urlBasic + "/im/v1/messages"                       // 发送消息接口
	urlImV1MessagesRelayTemplate = urlBasic + "/im/v1/messages/%s/reply"              // 发送回复消息接口
)

const (
	MsgTypeChallengeFlag = "url_verification"
)

const (
	msgReceiveIdTypeOpenID  = "open_id"
	msgReceiveIdTypeUserID  = "user_id"
	msgReceiveIdTypeUnionID = "union_id"
	msgReceiveIdTypeEmail   = "email"
	msgReceiveIdTypeChatID  = "chat_id"
)

const (
	MsgTypeText      = "text"
	MsgTypeRichText  = "rich_text"
	MsgTypeCard      = "interactive"
	MsgTypePost      = "post"
	MsgTypeImage     = "image"
	MsgTypeShareCard = "share_chat"
)

const (
	EventChatTypePrivate = "p2p"
	EventChatTypeGroup   = "group"
)
