package lark_markdown

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	lark_core "github.com/zzzzer91/lark-go/core"
	lark_docx "github.com/zzzzer91/lark-go/service/docx/v1"
)

var (
	appID, appSecret = os.Getenv("LARK_APP_ID"), os.Getenv("LARK_APP_SECRET")
	cli              = lark_core.NewClient(appID, appSecret, time.Second*5)
	svc              = lark_docx.NewService(cli)
)

func TestParseDocxContent(t *testing.T) {
	docId := ""
	res, err := svc.GetDocxBlocks(context.Background(), docId)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	s, ImgTokens := ParseDocxContent(res.Data.Items)
	f, _ := os.OpenFile("test.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	f.Write([]byte(s))
	t.Log(ImgTokens)
}

func BenchmarkParseDocxContent(b *testing.B) {
	docId := ""
	res, err := svc.GetDocxBlocks(context.Background(), docId)
	assert.Nil(b, err)
	assert.NotNil(b, res)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseDocxContent(res.Data.Items)
	}
}
