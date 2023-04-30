package lark_im

import "github.com/bytedance/sonic"

func NewTextMsg(text string) *ImMsg {
	return &ImMsg{
		MsgType: MsgTypeText,
		Content: buildImV1MessageTextContent(text),
	}
}

func NewCardMsg(content string) *ImMsg {
	return &ImMsg{
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
