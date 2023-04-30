package lark_wiki

import (
	"fmt"

	lark_core "github.com/zzzzer91/lark-go/core"
)

type WikiService interface {
	GetWikiSpaceNodes(spaceId string) (*WikiSpaceNodesResponse, error)
}

func NewService(cli *lark_core.Client) WikiService {
	return &wikiServiceImpl{
		cli: cli,
	}
}

type wikiServiceImpl struct {
	cli *lark_core.Client
}

func (s *wikiServiceImpl) GetWikiSpaceNodes(spaceId string) (*WikiSpaceNodesResponse, error) {
	url := fmt.Sprintf(urlWikiV2GetSpaceNodesTemplate, spaceId)
	jd := new(WikiSpaceNodesResponse)
	err := s.cli.GetJson(url, jd)
	if err != nil {
		return nil, err
	}
	return jd, nil
}
