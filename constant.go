package lark

// lark open api URL
const (
	urlBasic                     = "https://open.feishu.cn/open-apis"
	urlTenantAccessToken         = urlBasic + "/auth/v3/tenant_access_token/internal" // Getting token API
	urlImV1Messages              = urlBasic + "/im/v1/messages"                       // Sending Message API
	urlImV1MessagesRelayTemplate = urlBasic + "/im/v1/messages/%s/reply"              // Reply Message API
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
