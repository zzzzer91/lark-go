package lark

import (
	"time"

	lark_core "github.com/zzzzer91/lark-go/core"
	lark_docx "github.com/zzzzer91/lark-go/service/docx/v1"
	lark_im "github.com/zzzzer91/lark-go/service/im/v1"
	lark_wiki "github.com/zzzzer91/lark-go/service/wiki/v2"
)

type LarkService interface {
	lark_im.ImService
	lark_docx.DocxService
	lark_wiki.WikiService
}

func NewService(appID, appSecret string, timeout time.Duration) LarkService {
	cli := lark_core.NewClient(appID, appSecret, timeout)
	return &larkServiceImpl{
		ImService:   lark_im.NewService(cli),
		DocxService: lark_docx.NewService(cli),
		WikiService: lark_wiki.NewService(cli),
	}
}

type larkServiceImpl struct {
	lark_im.ImService
	lark_docx.DocxService
	lark_wiki.WikiService
}
