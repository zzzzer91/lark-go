package lark

import (
	"time"

	lark_core "github.com/zzzzer91/lark-go/core"
	lark_docx "github.com/zzzzer91/lark-go/service/docx/v1"
	lark_driver "github.com/zzzzer91/lark-go/service/driver/v1"
	lark_im "github.com/zzzzer91/lark-go/service/im/v1"
	lark_wiki "github.com/zzzzer91/lark-go/service/wiki/v2"
)

type LarkService interface {
	lark_im.ImService
	lark_docx.DocxService
	lark_wiki.WikiService
	lark_driver.DriverService
}

type larkServiceImpl struct {
	lark_im.ImService
	lark_docx.DocxService
	lark_wiki.WikiService
	lark_driver.DriverService
}

func NewService(appID, appSecret string, timeout time.Duration) LarkService {
	cli := lark_core.NewClient(appID, appSecret, timeout)
	return &larkServiceImpl{
		ImService:     lark_im.NewService(cli),
		DocxService:   lark_docx.NewService(cli),
		WikiService:   lark_wiki.NewService(cli),
		DriverService: lark_driver.NewService(cli),
	}
}

type larkServiceMock struct {
	lark_im.ImService
	lark_docx.DocxService
	lark_wiki.WikiService
	lark_driver.DriverService
}

func NewServiceMock() LarkService {
	return &larkServiceMock{
		ImService:     lark_im.NewServiceMock(),
		DocxService:   lark_docx.NewServiceMock(),
		WikiService:   lark_wiki.NewServiceMock(),
		DriverService: lark_driver.NewServiceMock(),
	}
}
