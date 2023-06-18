package driver

import (
	"context"
	"io"
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
	content, _, err := svc.DownloadMedia(context.Background(), fileToken)
	assert.Nil(t, err)
	assert.NotNil(t, content)

	f, _ := os.OpenFile("test.png", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	io.Copy(f, content)
}
