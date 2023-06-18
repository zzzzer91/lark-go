package wiki

import "context"

type wikiServiceMock struct {
}

func NewServiceMock() WikiService {
	return &wikiServiceMock{}
}

func (m *wikiServiceMock) GetWikiSpaceNodes(ctx context.Context, spaceId string) (*WikiSpaceNodesResponse, error) {
	panic("not implemented") // TODO: Implement
}
