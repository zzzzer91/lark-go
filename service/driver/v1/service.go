package lark_driver

import (
	"context"
	"fmt"
	"io"

	lark_core "github.com/zzzzer91/lark-go/core"
)

type DriverService interface {
	DownloadMedia(ctx context.Context, fileToken string) (io.ReadCloser, string, error)
}

func NewService(cli *lark_core.Client) DriverService {
	return &driverServiceImpl{
		cli: cli,
	}
}

type driverServiceImpl struct {
	cli *lark_core.Client
}

func (s *driverServiceImpl) DownloadMedia(ctx context.Context, fileToken string) (io.ReadCloser, string, error) {
	url := fmt.Sprintf(urlDriverV1DownloadMediaTemplate, fileToken)
	return s.cli.Download(ctx, url)
}
