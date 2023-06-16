package lark_docx

import "context"

type docxServiceMock struct {
}

func NewServiceMock() DocxService {
	return &docxServiceMock{}
}

func (m *docxServiceMock) GetDocxBasicInfo(ctx context.Context, documentId string) (*DocxBasicInfoResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (m *docxServiceMock) GetDocxRawContent(ctx context.Context, documentId string) (*DocxRawContentResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (m *docxServiceMock) GetDocxBlocks(ctx context.Context, documentId string) (*DocxBlocksResponse, error) {
	panic("not implemented") // TODO: Implement
}
