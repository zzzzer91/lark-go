package lark

import "github.com/bytedance/sonic"

type ImV1Msg struct {
	Content string
	MsgType string
}

func NewTextMsg(text string) *ImV1Msg {
	return &ImV1Msg{
		MsgType: MsgTypeText,
		Content: buildImV1MessageTextContent(text),
	}
}

func NewCardMsg(content string) *ImV1Msg {
	return &ImV1Msg{
		MsgType: MsgTypeCard,
		Content: content,
	}
}

func buildImV1MessageTextContent(text string) string {
	content := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}
	res, _ := sonic.Marshal(content)
	return string(res)
}

func IsChallengeStage(tp string) bool {
	return tp == MsgTypeChallengeFlag
}

func NewEventSubscriptionMessageChallengeResponse(challenge string) *EventSubscriptionMessageChallengeResponse {
	return &EventSubscriptionMessageChallengeResponse{
		Challenge: challenge,
	}
}
