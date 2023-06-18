package wiki

import lark_core "github.com/zzzzer91/lark-go/core"

type (
	WikiSpaceNodesResponse struct {
		lark_core.CodeMsg
		Data struct {
			HasMore bool `json:"has_more"`
			Items   []*struct {
				Creator         string `json:"creator"`
				HasChild        bool   `json:"has_child"`
				NodeCreateTime  string `json:"node_create_time"`
				NodeToken       string `json:"node_token"`
				NodeType        string `json:"node_type"`
				ObjCreateTime   string `json:"obj_create_time"`
				ObjEditTime     string `json:"obj_edit_time"`
				ObjToken        string `json:"obj_token"`
				ObjType         string `json:"obj_type"`
				OriginNodeToken string `json:"origin_node_token"`
				OriginSpaceID   string `json:"origin_space_id"`
				Owner           string `json:"owner"`
				ParentNodeToken string `json:"parent_node_token"`
				SpaceID         string `json:"space_id"`
				Title           string `json:"title"`
			} `json:"items"`
			PageToken string `json:"page_token"`
		} `json:"data"`
	}
)
