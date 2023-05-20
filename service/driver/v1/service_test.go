package lark_driver

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	lark_core "github.com/zzzzer91/lark-go/core"
)

var (
	appID, appSecret = os.Getenv("LARK_APP_ID"), os.Getenv("LARK_APP_SECRET")
	cli              = lark_core.NewClient(appID, appSecret, time.Second*5)
	svc              = NewService(cli)
)

func Test_DownloadMedia(t *testing.T) {
	fileToken := ""
	_, err := svc.DownloadMedia(context.Background(), fileToken)
	assert.Nil(t, err)
}
