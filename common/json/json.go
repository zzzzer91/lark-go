package json

import (
	"io"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

type CodeMsgIface interface {
	GetCode() int
	GetMsg() string
}

func DecodeBody(body io.Reader, out CodeMsgIface) error {
	if err := sonic.ConfigDefault.NewDecoder(body).Decode(out); err != nil {
		return errors.WithStack(err)
	}
	if out.GetCode() != 0 {
		return errors.Errorf("json code %d is invalid, msg: %s", out.GetCode(), out.GetMsg())
	}
	return nil
}
