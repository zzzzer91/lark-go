package lark_driver

import (
	"context"
	"fmt"
	"io"

	lark_core "github.com/zzzzer91/lark-go/core"
)

type DriverService interface {
	DownloadMedia(ctx context.Context, fileToken string) (io.ReadCloser, error)
}

func NewService(cli *lark_core.Client) DriverService {
	return &driverServiceImpl{
		cli: cli,
	}
}

type driverServiceImpl struct {
	cli *lark_core.Client
}

func (s *driverServiceImpl) DownloadMedia(ctx context.Context, fileToken string) (io.ReadCloser, error) {
	url := fmt.Sprintf(urlDriverV1DownloadMediaTemplate, fileToken)
	res, err := s.cli.Download(ctx, url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
