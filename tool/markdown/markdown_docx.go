package lark_markdown

import (
	"strings"

	"github.com/zzzzer91/gopkg/urlx"
	lark_docx "github.com/zzzzer91/lark-go/service/docx/v1"
)

func ParseDocxContent(blocks []*lark_docx.Block) (string, []string) {
	p := &docxParser{
		sb:       new(strings.Builder),
		blockMap: make(map[string]*lark_docx.Block, len(blocks)),
	}
	p.sb.Grow(stringBuilderInitSize)
	p.parseContent(blocks)
	return p.sb.String(), p.ImgTokens
}

type docxParser struct {
	sb        *strings.Builder
	blockMap  map[string]*lark_docx.Block
	ImgTokens []string // 所有图片的 token
}

func (p *docxParser) parseContent(blocks []*lark_docx.Block) {
	var pageBlock *lark_docx.Block
	for _, block := range blocks {
		if block.BlockType == lark_docx.DocxBlockTypePage {
			pageBlock = block
		}
		p.blockMap[block.BlockId] = block
	}
	if pageBlock != nil {
		p.sb.WriteString("# ")
		p.parseDocxBlockText(pageBlock.Page)
		p.sb.WriteString(markdownSeparator)
		for _, blockId := range pageBlock.Children {
			p.parseSingleBlock(p.blockMap[blockId], 0)
		}
	}
}

func (p *docxParser) parseSingleBlock(b *lark_docx.Block, indentLevel int) {
	sb := p.sb
	switch b.BlockType {
	case lark_docx.DocxBlockTypeText:
		p.writeIndentSpaces(indentLevel)
		p.parseDocxBlockText(b.Text)
	case lark_docx.DocxBlockTypeHeading1:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("## ")
		p.parseDocxBlockText(b.Heading1)
	case lark_docx.DocxBlockTypeHeading2:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("### ")
		p.parseDocxBlockText(b.Heading2)
	case lark_docx.DocxBlockTypeHeading3:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("#### ")
		p.parseDocxBlockText(b.Heading3)
	case lark_docx.DocxBlockTypeHeading4:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("##### ")
		p.parseDocxBlockText(b.Heading4)
	case lark_docx.DocxBlockTypeHeading5:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("###### ")
		p.parseDocxBlockText(b.Heading5)
	case lark_docx.DocxBlockTypeHeading6:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("####### ")
		p.parseDocxBlockText(b.Heading6)
	case lark_docx.DocxBlockTypeHeading7:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("######## ")
		p.parseDocxBlockText(b.Heading7)
	case lark_docx.DocxBlockTypeHeading8:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("######### ")
		p.parseDocxBlockText(b.Heading8)
	case lark_docx.DocxBlockTypeHeading9:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("########## ")
		p.parseDocxBlockText(b.Heading9)
	case lark_docx.DocxBlockTypeBullet:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("- ")
		p.parseDocxBlockText(b.Bullet)
	case lark_docx.DocxBlockTypeOrdered:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("1. ")
		p.parseDocxBlockText(b.Ordered)
	case lark_docx.DocxBlockTypeCode:
		sb.WriteString("```")
		sb.WriteString(b.Code.Style.Language.String())
		sb.WriteString("\n")
		p.parseDocxBlockText(b.Code)
		sb.WriteString("\n```")
	case lark_docx.DocxBlockTypeQuote:
		p.writeIndentSpaces(indentLevel)
		sb.WriteString("> ")
		p.parseDocxBlockText(b.Quote)
	case lark_docx.DocxBlockTypeEquation:
		sb.WriteString("$$\n")
		p.parseDocxBlockText(b.Equation)
		sb.WriteString("\n$$")
	case lark_docx.DocxBlockTypeTodo:
		p.writeIndentSpaces(indentLevel)
		if *b.Todo.Style.Done {
			sb.WriteString("- [x] ")
		} else {
			sb.WriteString("- [ ] ")
		}
		p.parseDocxBlockText(b.Todo)
	case lark_docx.DocxBlockTypeImage:
		p.parseBlockImage(b.Image)
	case lark_docx.DocxBlockTypeTable:
		p.parseBlockTable(b.Table)
		b.Children = nil // 防止下面重复解析 children
	case lark_docx.DocxBlockTypeQuoteContainer:
		p.writeIndentSpaces(indentLevel)
		p.parseBlockQuoteContainer(b)
		b.Children = nil // 防止下面重复解析 children
	}
	sb.WriteString(markdownSeparator)

	for _, blockId := range b.Children {
		p.parseSingleBlock(p.blockMap[blockId], indentLevel+1)
	}
}

func (p *docxParser) writeIndentSpaces(indentLevel int) {
	for indentLevel > 0 {
		p.sb.WriteString(markdownIndentSpaces)
		indentLevel--
	}
}

func (p *docxParser) parseDocxBlockText(t *lark_docx.Text) {
	if t == nil {
		return
	}
	for _, e := range t.Elements {
		p.parseDocxTextElement(e)
	}
}

func (p *docxParser) parseDocxTextElement(e *lark_docx.TextElement) {
	if e.TextRun != nil {
		p.parseDocxTextElementTextRun(e.TextRun)
	}
	if e.MentionUser != nil {
		p.sb.WriteString(*e.MentionUser.UserId)
	}
	if e.MentionDoc != nil {
		p.sb.WriteString("[")
		p.sb.WriteString(*e.MentionDoc.Title)
		p.sb.WriteString("](")
		p.sb.WriteString(urlx.UnescapeURL(*e.MentionDoc.Url))
		p.sb.WriteString(")")
	}
	if e.Equation != nil {
		p.sb.WriteString("$$")
		p.sb.WriteString(strings.TrimSuffix(*e.Equation.Content, "\n"))
		p.sb.WriteString("$$")
	}
}

func (p *docxParser) parseDocxTextElementTextRun(tr *lark_docx.TextRun) {
	postWrite := ""
	if style := tr.TextElementStyle; style != nil {
		if style.Bold {
			p.sb.WriteString("**")
			postWrite = "**"
		} else if style.Italic {
			p.sb.WriteString("*")
			postWrite = "*"
		} else if style.Strikethrough {
			p.sb.WriteString("~~")
			postWrite = "~~"
		} else if style.Underline {
			// ignore underline
		} else if style.InlineCode {
			p.sb.WriteString("`")
			postWrite = "`"
		} else if link := style.Link; link != nil {
			p.sb.WriteString("[")
			postWrite = "](" + urlx.UnescapeURL(*link.Url) + ")"
		}
	}
	p.sb.WriteString(tr.Content)
	p.sb.WriteString(postWrite)
}

func (p *docxParser) parseBlockImage(img *lark_docx.Image) {
	p.sb.WriteString("![](")
	p.sb.WriteString(img.Token)
	p.sb.WriteString(")")
	p.ImgTokens = append(p.ImgTokens, img.Token)
}

func (p *docxParser) parseBlockTable(t *lark_docx.Table) {
	colSize := t.Property.ColumnSize
	rowSize := t.Property.RowSize
	rows := make([][]string, rowSize)
	buffer := make([]string, colSize*rowSize)
	for i := range rows {
		rows[i] = buffer[i*colSize : (i+1)*colSize][:0]
	}
	blockIds := t.Cells
	// 解析成 table
	// table 头部
	p.sb.WriteString("|")
	for j := 0; j < colSize; j++ {
		p.parseTableCell(p.blockMap[blockIds[j]])
		p.sb.WriteString("|")
	}
	p.sb.WriteString("\n")
	// table 的 |---|---|
	p.sb.WriteString("|")
	for j := 0; j < colSize; j++ {
		p.sb.WriteString("-")
		p.sb.WriteString("|")
	}
	if rowSize > 1 {
		p.sb.WriteString("\n")
	}
	// table 内容
	for i := 0; i < rowSize; i++ {
		p.sb.WriteString("|")
		for j := 0; j < colSize; j++ {
			p.parseTableCell(p.blockMap[blockIds[i*colSize+j]])
			p.sb.WriteString("|")
		}
		if i != rowSize-1 {
			p.sb.WriteString("\n")
		}
	}
}

func (p *docxParser) parseBlockQuoteContainer(b *lark_docx.Block) {
	p.sb.WriteString("> ")
	for _, childBlockId := range b.Children {
		p.parseDocxBlockText(p.blockMap[childBlockId].Text)
	}
}

func (p *docxParser) parseTableCell(b *lark_docx.Block) {
	t := p.blockMap[b.Children[0]].Text
	if t == nil {
		return
	}
	for _, e := range t.Elements {
		if e.TextRun != nil {
			e.TextRun.Content = strings.ReplaceAll(e.TextRun.Content, "|", "\\|") // 转译 table 中的 “｜”
		}
		p.parseDocxTextElement(e)
	}
}
