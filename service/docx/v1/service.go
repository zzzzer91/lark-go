package docx

import (
	"context"
	"fmt"

	lark_core "github.com/zzzzer91/lark-go/core"
)

type DocxService interface {
	GetDocxBasicInfo(ctx context.Context, documentId string) (*DocxBasicInfoResponse, error)
	GetDocxRawContent(ctx context.Context, documentId string) (*DocxRawContentResponse, error)
	GetDocxBlocks(ctx context.Context, documentId string) (*DocxBlocksResponse, error)
}

func NewService(cli *lark_core.Client) DocxService {
	return &docxServiceImpl{
		cli: cli,
	}
}

type docxServiceImpl struct {
	cli *lark_core.Client
}

func (s *docxServiceImpl) GetDocxBasicInfo(ctx context.Context, documentId string) (*DocxBasicInfoResponse, error) {
	url := fmt.Sprintf(urlDocxV1DocumentsTemplate, documentId)
	jd := new(DocxBasicInfoResponse)
	err := s.cli.GetJson(ctx, url, jd)
	if err != nil {
		return nil, err
	}
	return jd, nil
}

func (s *docxServiceImpl) GetDocxRawContent(ctx context.Context, documentId string) (*DocxRawContentResponse, error) {
	url := fmt.Sprintf(urlDocxV1RawContentTemplate, documentId)
	jd := new(DocxRawContentResponse)
	err := s.cli.GetJson(ctx, url, jd)
	if err != nil {
		return nil, err
	}
	return jd, nil
}

func (s *docxServiceImpl) GetDocxBlocks(ctx context.Context, documentId string) (*DocxBlocksResponse, error) {
	url := fmt.Sprintf(urlDocxV1BlocksTemplate, documentId)
	jd := new(DocxBlocksResponse)
	err := s.cli.GetJson(ctx, url, jd)
	if err != nil {
		return nil, err
	}
	return jd, nil
}
