package docx

import lark_core "github.com/zzzzer91/lark-go/core"

type (
	DocxBasicInfoResponse struct {
		lark_core.CodeMsg
		Data struct {
			Document struct {
				DocumentID string `json:"document_id"`
				RevisionID int    `json:"revision_id"`
				Title      string `json:"title"`
			} `json:"document"`
		} `json:"data"`
	}

	DocxRawContentResponse struct {
		lark_core.CodeMsg
		Data struct {
			Content string `json:"content"`
		} `json:"data"`
	}

	DocxBlocksResponse struct {
		lark_core.CodeMsg
		Data struct {
			Items     []*Block `json:"items,omitempty"`      // block 的 children 列表
			PageToken *string  `json:"page_token,omitempty"` // 下一个分页的分页标记
			HasMore   bool     `json:"has_more,omitempty"`   // 是否还有下一个分页
		} `json:"data"`
	}
)

type DocxBlockType int64

const (
	DocxBlockTypePage           DocxBlockType = 1   // 文档 Block
	DocxBlockTypeText           DocxBlockType = 2   // 文本 Block
	DocxBlockTypeHeading1       DocxBlockType = 3   // 一级标题 Block
	DocxBlockTypeHeading2       DocxBlockType = 4   // 二级标题 Block
	DocxBlockTypeHeading3       DocxBlockType = 5   // 三级标题 Block
	DocxBlockTypeHeading4       DocxBlockType = 6   // 四级标题 Block
	DocxBlockTypeHeading5       DocxBlockType = 7   // 五级标题 Block
	DocxBlockTypeHeading6       DocxBlockType = 8   // 六级标题 Block
	DocxBlockTypeHeading7       DocxBlockType = 9   // 七级标题 Block
	DocxBlockTypeHeading8       DocxBlockType = 10  // 八级标题 Block
	DocxBlockTypeHeading9       DocxBlockType = 11  // 九级标题 Block
	DocxBlockTypeBullet         DocxBlockType = 12  // 无序列表 Block
	DocxBlockTypeOrdered        DocxBlockType = 13  // 有序列表 Block
	DocxBlockTypeCode           DocxBlockType = 14  // 代码块 Block
	DocxBlockTypeQuote          DocxBlockType = 15  // 引用 Block
	DocxBlockTypeEquation       DocxBlockType = 16  // 公式 Block
	DocxBlockTypeTodo           DocxBlockType = 17  // 任务 Block
	DocxBlockTypeBitable        DocxBlockType = 18  // 多维表格 Block
	DocxBlockTypeCallout        DocxBlockType = 19  // 高亮块 Block
	DocxBlockTypeChatCard       DocxBlockType = 20  // 群聊卡片 Block
	DocxBlockTypeDiagram        DocxBlockType = 21  // 流程图/UML Block
	DocxBlockTypeDivider        DocxBlockType = 22  // 分割线 Block
	DocxBlockTypeFile           DocxBlockType = 23  // 文件 Block
	DocxBlockTypeGrid           DocxBlockType = 24  // 分栏 Block
	DocxBlockTypeGridColumn     DocxBlockType = 25  // 分栏列 Block
	DocxBlockTypeIframe         DocxBlockType = 26  // 内嵌 Block
	DocxBlockTypeImage          DocxBlockType = 27  // 图片 Block
	DocxBlockTypeISV            DocxBlockType = 28  // 三方 Block
	DocxBlockTypeMindnote       DocxBlockType = 29  // 思维笔记 Block
	DocxBlockTypeSheet          DocxBlockType = 30  // 电子表格 Block
	DocxBlockTypeTable          DocxBlockType = 31  // 表格 Block
	DocxBlockTypeTableCell      DocxBlockType = 32  // 单元格 Block
	DocxBlockTypeView           DocxBlockType = 33  // 视图 Block
	DocxBlockTypeQuoteContainer DocxBlockType = 34  // 引用容器 Block
	DocxBlockTypeTask           DocxBlockType = 35  // 任务容器 Block
	DocxBlockTypeOKR            DocxBlockType = 36  // OKR容器 Block
	DocxBlockTypeOKRObjective   DocxBlockType = 37  // OKR Objective容器 Block
	DocxBlockTypeOKRKeyResult   DocxBlockType = 38  // OKR KeyResult容器 Block
	DocxBlockTypeProgress       DocxBlockType = 39  // Progress容器 Block
	DocxBlockTypeUndefined      DocxBlockType = 999 // 未支持 Block
)

// DocxAlign Block 的排版方式，比如居左等
type DocxAlign int64

const (
	DocxAlignLeft   DocxAlign = 1 // 居左排版
	DocxAlignCenter DocxAlign = 2 // 居中排版
	DocxAlignRight  DocxAlign = 3 // 居右排版
)

// DocxCodeLanguage 代码块语言
type DocxCodeLanguage int64

func (code DocxCodeLanguage) String() string {
	if int(code) >= len(docxCodeLang2MdStr) {
		return ""
	}
	return docxCodeLang2MdStr[code]
}

var docxCodeLang2MdStr = [...]string{
	DocxCodeLanguagePlainText:    "",
	DocxCodeLanguageABAP:         "abap",
	DocxCodeLanguageAda:          "ada",
	DocxCodeLanguageApache:       "apache",
	DocxCodeLanguageApex:         "apex",
	DocxCodeLanguageAssembly:     "assembly",
	DocxCodeLanguageBash:         "bash",
	DocxCodeLanguageCSharp:       "csharp",
	DocxCodeLanguageCPlusPlus:    "cpp",
	DocxCodeLanguageC:            "c",
	DocxCodeLanguageCOBOL:        "cobol",
	DocxCodeLanguageCSS:          "css",
	DocxCodeLanguageCoffeeScript: "coffeescript",
	DocxCodeLanguageD:            "d",
	DocxCodeLanguageDart:         "dart",
	DocxCodeLanguageDelphi:       "delphi",
	DocxCodeLanguageDjango:       "django",
	DocxCodeLanguageDockerfile:   "dockerfile",
	DocxCodeLanguageErlang:       "erlang",
	DocxCodeLanguageFortran:      "fortran",
	DocxCodeLanguageFoxPro:       "foxpro",
	DocxCodeLanguageGo:           "go",
	DocxCodeLanguageGroovy:       "groovy",
	DocxCodeLanguageHTML:         "html",
	DocxCodeLanguageHTMLBars:     "htmlbars",
	DocxCodeLanguageHTTP:         "http",
	DocxCodeLanguageHaskell:      "haskell",
	DocxCodeLanguageJSON:         "json",
	DocxCodeLanguageJava:         "java",
	DocxCodeLanguageJavaScript:   "javascript",
	DocxCodeLanguageJulia:        "julia",
	DocxCodeLanguageKotlin:       "kotlin",
	DocxCodeLanguageLateX:        "latex",
	DocxCodeLanguageLisp:         "lisp",
	DocxCodeLanguageLogo:         "logo",
	DocxCodeLanguageLua:          "lua",
	DocxCodeLanguageMATLAB:       "matlab",
	DocxCodeLanguageMakefile:     "makefile",
	DocxCodeLanguageMarkdown:     "markdown",
	DocxCodeLanguageNginx:        "nginx",
	DocxCodeLanguageObjective:    "objectivec",
	DocxCodeLanguageOpenEdgeABL:  "openedge-abl",
	DocxCodeLanguagePHP:          "php",
	DocxCodeLanguagePerl:         "perl",
	DocxCodeLanguagePostScript:   "postscript",
	DocxCodeLanguagePower:        "powershell",
	DocxCodeLanguageProlog:       "prolog",
	DocxCodeLanguageProtoBuf:     "protobuf",
	DocxCodeLanguagePython:       "python",
	DocxCodeLanguageR:            "r",
	DocxCodeLanguageRPG:          "rpg",
	DocxCodeLanguageRuby:         "ruby",
	DocxCodeLanguageRust:         "rust",
	DocxCodeLanguageSAS:          "sas",
	DocxCodeLanguageSCSS:         "scss",
	DocxCodeLanguageSQL:          "sql",
	DocxCodeLanguageScala:        "scala",
	DocxCodeLanguageScheme:       "scheme",
	DocxCodeLanguageScratch:      "scratch",
	DocxCodeLanguageShell:        "shell",
	DocxCodeLanguageSwift:        "swift",
	DocxCodeLanguageThrift:       "thrift",
	DocxCodeLanguageTypeScript:   "typescript",
	DocxCodeLanguageVBScript:     "vbscript",
	DocxCodeLanguageVisual:       "vbnet",
	DocxCodeLanguageXML:          "xml",
	DocxCodeLanguageYAML:         "yaml",
}

const (
	DocxCodeLanguagePlainText    DocxCodeLanguage = 1  // PlainText
	DocxCodeLanguageABAP         DocxCodeLanguage = 2  // ABAP
	DocxCodeLanguageAda          DocxCodeLanguage = 3  // Ada
	DocxCodeLanguageApache       DocxCodeLanguage = 4  // Apache
	DocxCodeLanguageApex         DocxCodeLanguage = 5  // Apex
	DocxCodeLanguageAssembly     DocxCodeLanguage = 6  // Assembly
	DocxCodeLanguageBash         DocxCodeLanguage = 7  // Bash
	DocxCodeLanguageCSharp       DocxCodeLanguage = 8  // CSharp
	DocxCodeLanguageCPlusPlus    DocxCodeLanguage = 9  // C++
	DocxCodeLanguageC            DocxCodeLanguage = 10 // C
	DocxCodeLanguageCOBOL        DocxCodeLanguage = 11 // COBOL
	DocxCodeLanguageCSS          DocxCodeLanguage = 12 // CSS
	DocxCodeLanguageCoffeeScript DocxCodeLanguage = 13 // CoffeeScript
	DocxCodeLanguageD            DocxCodeLanguage = 14 // D
	DocxCodeLanguageDart         DocxCodeLanguage = 15 // Dart
	DocxCodeLanguageDelphi       DocxCodeLanguage = 16 // Delphi
	DocxCodeLanguageDjango       DocxCodeLanguage = 17 // Django
	DocxCodeLanguageDockerfile   DocxCodeLanguage = 18 // Dockerfile
	DocxCodeLanguageErlang       DocxCodeLanguage = 19 // Erlang
	DocxCodeLanguageFortran      DocxCodeLanguage = 20 // Fortran
	DocxCodeLanguageFoxPro       DocxCodeLanguage = 21 // FoxPro
	DocxCodeLanguageGo           DocxCodeLanguage = 22 // Go
	DocxCodeLanguageGroovy       DocxCodeLanguage = 23 // Groovy
	DocxCodeLanguageHTML         DocxCodeLanguage = 24 // HTML
	DocxCodeLanguageHTMLBars     DocxCodeLanguage = 25 // HTMLBars
	DocxCodeLanguageHTTP         DocxCodeLanguage = 26 // HTTP
	DocxCodeLanguageHaskell      DocxCodeLanguage = 27 // Haskell
	DocxCodeLanguageJSON         DocxCodeLanguage = 28 // JSON
	DocxCodeLanguageJava         DocxCodeLanguage = 29 // Java
	DocxCodeLanguageJavaScript   DocxCodeLanguage = 30 // JavaScript
	DocxCodeLanguageJulia        DocxCodeLanguage = 31 // Julia
	DocxCodeLanguageKotlin       DocxCodeLanguage = 32 // Kotlin
	DocxCodeLanguageLateX        DocxCodeLanguage = 33 // LateX
	DocxCodeLanguageLisp         DocxCodeLanguage = 34 // Lisp
	DocxCodeLanguageLogo         DocxCodeLanguage = 35 // Logo
	DocxCodeLanguageLua          DocxCodeLanguage = 36 // Lua
	DocxCodeLanguageMATLAB       DocxCodeLanguage = 37 // MATLAB
	DocxCodeLanguageMakefile     DocxCodeLanguage = 38 // Makefile
	DocxCodeLanguageMarkdown     DocxCodeLanguage = 39 // Markdown
	DocxCodeLanguageNginx        DocxCodeLanguage = 40 // Nginx
	DocxCodeLanguageObjective    DocxCodeLanguage = 41 // Objective
	DocxCodeLanguageOpenEdgeABL  DocxCodeLanguage = 42 // OpenEdgeABL
	DocxCodeLanguagePHP          DocxCodeLanguage = 43 // PHP
	DocxCodeLanguagePerl         DocxCodeLanguage = 44 // Perl
	DocxCodeLanguagePostScript   DocxCodeLanguage = 45 // PostScript
	DocxCodeLanguagePower        DocxCodeLanguage = 46 // Power
	DocxCodeLanguageProlog       DocxCodeLanguage = 47 // Prolog
	DocxCodeLanguageProtoBuf     DocxCodeLanguage = 48 // ProtoBuf
	DocxCodeLanguagePython       DocxCodeLanguage = 49 // Python
	DocxCodeLanguageR            DocxCodeLanguage = 50 // R
	DocxCodeLanguageRPG          DocxCodeLanguage = 51 // RPG
	DocxCodeLanguageRuby         DocxCodeLanguage = 52 // Ruby
	DocxCodeLanguageRust         DocxCodeLanguage = 53 // Rust
	DocxCodeLanguageSAS          DocxCodeLanguage = 54 // SAS
	DocxCodeLanguageSCSS         DocxCodeLanguage = 55 // SCSS
	DocxCodeLanguageSQL          DocxCodeLanguage = 56 // SQL
	DocxCodeLanguageScala        DocxCodeLanguage = 57 // Scala
	DocxCodeLanguageScheme       DocxCodeLanguage = 58 // Scheme
	DocxCodeLanguageScratch      DocxCodeLanguage = 59 // Scratch
	DocxCodeLanguageShell        DocxCodeLanguage = 60 // Shell
	DocxCodeLanguageSwift        DocxCodeLanguage = 61 // Swift
	DocxCodeLanguageThrift       DocxCodeLanguage = 62 // Thrift
	DocxCodeLanguageTypeScript   DocxCodeLanguage = 63 // TypeScript
	DocxCodeLanguageVBScript     DocxCodeLanguage = 64 // VBScript
	DocxCodeLanguageVisual       DocxCodeLanguage = 65 // Visual
	DocxCodeLanguageXML          DocxCodeLanguage = 66 // XML
	DocxCodeLanguageYAML         DocxCodeLanguage = 67 // YAML
)

type (
	Block struct {
		BlockId        string          `json:"block_id,omitempty"`        // Block 唯一标识
		BlockType      DocxBlockType   `json:"block_type,omitempty"`      // block 类型
		ParentId       string          `json:"parent_id,omitempty"`       // block 的父亲 id
		Children       []string        `json:"children,omitempty"`        // block 的孩子 id 列表
		Page           *Text           `json:"page,omitempty"`            // 文档 Block
		Text           *Text           `json:"text,omitempty"`            // 文本 Block
		Heading1       *Text           `json:"heading1,omitempty"`        // 一级标题 Block
		Heading2       *Text           `json:"heading2,omitempty"`        // 二级标题 Block
		Heading3       *Text           `json:"heading3,omitempty"`        // 三级标题 Block
		Heading4       *Text           `json:"heading4,omitempty"`        // 四级标题 Block
		Heading5       *Text           `json:"heading5,omitempty"`        // 五级标题 Block
		Heading6       *Text           `json:"heading6,omitempty"`        // 六级标题 Block
		Heading7       *Text           `json:"heading7,omitempty"`        // 七级标题 Block
		Heading8       *Text           `json:"heading8,omitempty"`        // 八级标题 Block
		Heading9       *Text           `json:"heading9,omitempty"`        // 九级标题 Block
		Bullet         *Text           `json:"bullet,omitempty"`          // 无序列表 Block
		Ordered        *Text           `json:"ordered,omitempty"`         // 有序列表 Block
		Code           *Text           `json:"code,omitempty"`            // 代码块 Block
		Quote          *Text           `json:"quote,omitempty"`           // 引用 Block
		Equation       *Text           `json:"equation,omitempty"`        // 公式 Block
		Todo           *Text           `json:"todo,omitempty"`            // 待办事项 Block
		Bitable        *Bitable        `json:"bitable,omitempty"`         // 多维表格 Block
		Callout        *Callout        `json:"callout,omitempty"`         // 高亮块 Block
		ChatCard       *ChatCard       `json:"chat_card,omitempty"`       // 群聊卡片 Block
		Diagram        *Diagram        `json:"diagram,omitempty"`         // 流程图/UML Block
		Divider        *Divider        `json:"divider,omitempty"`         // 分割线 Block
		File           *File           `json:"file,omitempty"`            // 文件 Block
		Grid           *Grid           `json:"grid,omitempty"`            // 分栏 Block
		GridColumn     *GridColumn     `json:"grid_column,omitempty"`     // 分栏列 Block
		Iframe         *Iframe         `json:"iframe,omitempty"`          // 内嵌 Block
		Image          *Image          `json:"image,omitempty"`           // 图片 Block
		Isv            *Isv            `json:"isv,omitempty"`             // 三方 Block
		AddOns         *AddOns         `json:"add_ons,omitempty"`         // Add-ons
		Mindnote       *Mindnote       `json:"mindnote,omitempty"`        // 思维笔记 Block
		Sheet          *Sheet          `json:"sheet,omitempty"`           // 电子表格 Block
		Table          *Table          `json:"table,omitempty"`           // 表格 Block
		TableCell      *TableCell      `json:"table_cell,omitempty"`      // 单元格 Block
		View           *View           `json:"view,omitempty"`            // 视图 Block
		Undefined      *Undefined      `json:"undefined,omitempty"`       // 未支持 Block
		QuoteContainer *QuoteContainer `json:"quote_container,omitempty"` // 引用容器 Block
	}

	Text struct {
		Style    *TextStyle     `json:"style,omitempty"`    // 文本样式
		Elements []*TextElement `json:"elements,omitempty"` // 文本元素
	}

	Bitable struct {
		Token    *string `json:"token,omitempty"`     // 多维表格文档 Token
		ViewType *int    `json:"view_type,omitempty"` // 类型
	}

	Callout struct {
		BackgroundColor *int    `json:"background_color,omitempty"` // 高亮块背景色
		BorderColor     *int    `json:"border_color,omitempty"`     // 边框色
		TextColor       *int    `json:"text_color,omitempty"`       // 文字颜色
		EmojiId         *string `json:"emoji_id,omitempty"`         // 高亮块图标
	}

	ChatCard struct {
		ChatId *string `json:"chat_id,omitempty"` // 群聊天会话 ID
		Align  *int    `json:"align,omitempty"`   // 对齐方式
	}

	Diagram struct {
		DiagramType *int `json:"diagram_type,omitempty"` // 绘图类型
	}

	Divider struct {
	}

	File struct {
		Token *string `json:"token,omitempty"` // 附件 Token
		Name  *string `json:"name,omitempty"`  // 文件名
	}

	Grid struct {
		ColumnSize *int `json:"column_size,omitempty"` // 分栏列数量
	}

	GridColumn struct {
		WidthRatio *int `json:"width_ratio,omitempty"` // 当前分栏列占整个分栏的比例
	}

	Iframe struct {
		Component *IframeComponent `json:"component,omitempty"` // iframe 的组成元素
	}

	Image struct {
		Width  int    `json:"width,omitempty"`  // 宽度单位 px
		Height int    `json:"height,omitempty"` // 高度
		Token  string `json:"token,omitempty"`  // 图片 Token
	}

	Isv struct {
		ComponentId     *string `json:"component_id,omitempty"`      // 团队互动应用唯一ID
		ComponentTypeId *string `json:"component_type_id,omitempty"` // 团队互动应用类型，比如信息收集"blk_5f992038c64240015d280958"
	}

	AddOns struct {
		ComponentId     *string `json:"component_id,omitempty"`      // 团队互动应用唯一ID
		ComponentTypeId *string `json:"component_type_id,omitempty"` // 团队互动应用类型，比如问答互动"blk_636a0a6657db8001c8df5488"
		Record          *string `json:"record,omitempty"`            // 文档小组件内容数据，JSON 字符串
	}

	Mindnote struct {
		Token *string `json:"token,omitempty"` // 思维导图 token
	}

	Sheet struct {
		Token      *string `json:"token,omitempty"`       // 电子表格 block 的 token
		RowSize    *int    `json:"row_size,omitempty"`    // 电子表格行数量
		ColumnSize *int    `json:"column_size,omitempty"` // 电子表格列数量
	}

	Table struct {
		Cells    []string       `json:"cells,omitempty"`    // 单元格数组，数组元素为 Table Cell Block 的 ID
		Property *TableProperty `json:"property,omitempty"` // 表格属性
	}

	TableCell struct {
	}

	View struct {
		ViewType *int `json:"view_type,omitempty"` // 视图类型
	}

	Undefined struct {
	}

	QuoteContainer struct {
	}
)

type (
	TextStyle struct {
		Align    *DocxAlign        `json:"align,omitempty"`    // 对齐方式
		Done     *bool             `json:"done,omitempty"`     // todo 的完成状态
		Folded   *bool             `json:"folded,omitempty"`   // 文本的折叠状态
		Language *DocxCodeLanguage `json:"language,omitempty"` // 代码块语言
		Wrap     *bool             `json:"wrap,omitempty"`     // 代码块是否自动换行
	}

	TextElement struct {
		TextRun     *TextRun          `json:"text_run,omitempty"`     // 文字
		MentionUser *MentionUser      `json:"mention_user,omitempty"` // @用户
		MentionDoc  *MentionDoc       `json:"mention_doc,omitempty"`  // @文档
		Reminder    *Reminder         `json:"reminder,omitempty"`     // 日期提醒
		File        *InlineFile       `json:"file,omitempty"`         // 内联附件
		Undefined   *UndefinedElement `json:"undefined,omitempty"`    // 未支持的 TextElement
		InlineBlock *InlineBlock      `json:"inline_block,omitempty"` // 内联 block
		Equation    *Equation         `json:"equation,omitempty"`     // 公式
	}

	TextRun struct {
		Content          string            `json:"content,omitempty"`            // 文本内容
		TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"` // 文本局部样式
	}

	MentionUser struct {
		UserId           *string           `json:"user_id,omitempty"`            // 用户 OpenID
		TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"` // 文本局部样式
	}

	MentionDoc struct {
		Token            *string           `json:"token,omitempty"`              // 云文档 token
		ObjType          *int              `json:"obj_type,omitempty"`           // 云文档类型
		Url              *string           `json:"url,omitempty"`                // 云文档链接（需要 url_encode)
		Title            *string           `json:"title,omitempty"`              // 文档标题，只读属性
		TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"` // 文本局部样式
	}

	Reminder struct {
		CreateUserId     *string           `json:"create_user_id,omitempty"`     // 创建者用户 ID
		IsNotify         *bool             `json:"is_notify,omitempty"`          // 是否通知
		IsWholeDay       *bool             `json:"is_whole_day,omitempty"`       // 是日期还是整点小时
		ExpireTime       *string           `json:"expire_time,omitempty"`        // 事件发生的时间（毫秒级事件戳）
		NotifyTime       *string           `json:"notify_time,omitempty"`        // 触发通知的时间（毫秒级时间戳）
		TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"` // 文本局部样式
	}

	InlineFile struct {
		FileToken        *string           `json:"file_token,omitempty"`         // 附件 token
		SourceBlockId    *string           `json:"source_block_id,omitempty"`    // 当前文档中该附件所处的 block 的 id
		TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"` // 文本局部样式
	}

	UndefinedElement struct {
	}

	InlineBlock struct {
		BlockId          *string           `json:"block_id,omitempty"`           // 关联的内联状态的 block 的 block_id
		TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"` // 文本局部样式
	}

	Equation struct {
		Content          *string           `json:"content,omitempty"`            // 符合 KaTeX 语法的公式内容，语法规则请参考：https://katex.org/docs/supported.html
		TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"` // 文本局部样式
	}

	TextElementStyle struct {
		Bold            bool     `json:"bold,omitempty"`             // 加粗
		Italic          bool     `json:"italic,omitempty"`           // 斜体
		Strikethrough   bool     `json:"strikethrough,omitempty"`    // 删除线
		Underline       bool     `json:"underline,omitempty"`        // 下划线
		InlineCode      bool     `json:"inline_code,omitempty"`      // inline 代码
		BackgroundColor *int     `json:"background_color,omitempty"` // 背景色
		TextColor       *int     `json:"text_color,omitempty"`       // 字体颜色
		Link            *Link    `json:"link,omitempty"`             // 链接
		CommentIds      []string `json:"comment_ids,omitempty"`      // 评论 id 列表
	}

	Link struct {
		Url *string `json:"url,omitempty"` // 超链接指向的 url (需要 url_encode)
	}
)

type (
	TableProperty struct {
		RowSize     int              `json:"row_size,omitempty"`     // 行数
		ColumnSize  int              `json:"column_size,omitempty"`  // 列数
		ColumnWidth []int             `json:"column_width,omitempty"` // 列宽，单位px
		MergeInfo   []*TableMergeInfo `json:"merge_info,omitempty"`   // 单元格合并信息
	}

	TableMergeInfo struct {
		RowSpan *int `json:"row_span,omitempty"` // 从当前行索引起被合并的连续行数
		ColSpan *int `json:"col_span,omitempty"` // 从当前列索引起被合并的连续列数
	}
)

type (
	IframeComponent struct {
		IframeType *int    `json:"iframe_type,omitempty"` // iframe 类型
		Url        *string `json:"url,omitempty"`         // iframe 目标 url（需要进行 url_encode）
	}
)
