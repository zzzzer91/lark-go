package lark

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var larkService Service

func init() {
	appID, appSecret := "", ""
	larkService = NewService(appID, appSecret, time.Second*5)
}

func TestGetDocBasicInfo(t *testing.T) {
	docId := ""
	res, err := larkService.GetDocBasicInfo(docId)
	assert.Nil(t, err)
	assert.Nil(t, res)
}
