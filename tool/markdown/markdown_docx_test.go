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

func TestGetDocBasicInfo(t *testing.T) {
	appID, appSecret := "", ""
	cli := lark_core.NewClient(appID, appSecret, time.Second*5)
	svc := lark_docx.NewService(cli)
	docId := ""
	res, err := svc.GetDocxBlocks(context.Background(), docId)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	p := NewParser()
	s := p.ParseDocxContent(res.Data.Items)
	f, _ := os.OpenFile("test.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	f.Write([]byte(s))
}
