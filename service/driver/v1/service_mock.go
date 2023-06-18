package driver

import (
	"context"
	"io"
)

type driverServiceMock struct {
}

func NewServiceMock() DriverService {
	return &driverServiceMock{}
}

func (m *driverServiceMock) DownloadMedia(ctx context.Context, fileToken string) (io.ReadCloser, string, error) {
	panic("not implemented") // TODO: Implement
}
