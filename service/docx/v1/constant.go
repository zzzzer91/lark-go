package lark_docx

import (
	lark_core "github.com/zzzzer91/lark-go/core"
)

const (
	urlDocxV1DocumentsTemplate  = lark_core.UrlBasic + "/docx/v1/documents/%s"             // Getting document basic information APi.
	urlDocxV1RawContentTemplate = lark_core.UrlBasic + "/docx/v1/documents/%s/raw_content" // Getting raw content of document APi.
	urlDocxV1BlocksTemplate     = lark_core.UrlBasic + "/docx/v1/documents/%s/blocks"      // Getting blocks of document APi.
)
