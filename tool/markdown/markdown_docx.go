package lark_markdown

import (
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/zzzzer91/gopkg/urlx"
	lark_docx "github.com/zzzzer91/lark-go/service/docx/v1"
)

func ParseDocxContent(blocks []*lark_docx.Block) (string, []string) {
	p := &docxParser{
		sb:       new(strings.Builder),
		blockMap: make(map[string]*lark_docx.Block),
	}
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
		parseDocxBlockText(p.sb, pageBlock.Page)
		p.sb.WriteString("\n\n")
		for _, blockId := range pageBlock.Children {
			p.parseSingleBlock(p.blockMap[blockId], 0)
		}
	}
}

func (p *docxParser) parseSingleBlock(b *lark_docx.Block, indentLevel int) {
	sb := p.sb
	sb.WriteString(strings.Repeat(" ", markdownIndentCount*indentLevel))
	switch b.BlockType {
	case lark_docx.DocxBlockTypeText:
		parseDocxBlockText(sb, b.Text)
	case lark_docx.DocxBlockTypeHeading1:
		sb.WriteString("## ")
		parseDocxBlockText(sb, b.Heading1)
	case lark_docx.DocxBlockTypeHeading2:
		sb.WriteString("### ")
		parseDocxBlockText(sb, b.Heading2)
	case lark_docx.DocxBlockTypeHeading3:
		sb.WriteString("#### ")
		parseDocxBlockText(sb, b.Heading3)
	case lark_docx.DocxBlockTypeHeading4:
		sb.WriteString("##### ")
		parseDocxBlockText(sb, b.Heading4)
	case lark_docx.DocxBlockTypeHeading5:
		sb.WriteString("###### ")
		parseDocxBlockText(sb, b.Heading5)
	case lark_docx.DocxBlockTypeHeading6:
		sb.WriteString("####### ")
		parseDocxBlockText(sb, b.Heading6)
	case lark_docx.DocxBlockTypeHeading7:
		sb.WriteString("######## ")
		parseDocxBlockText(sb, b.Heading7)
	case lark_docx.DocxBlockTypeHeading8:
		sb.WriteString("######### ")
		parseDocxBlockText(sb, b.Heading8)
	case lark_docx.DocxBlockTypeHeading9:
		sb.WriteString("########## ")
		parseDocxBlockText(sb, b.Heading9)
	case lark_docx.DocxBlockTypeBullet:
		sb.WriteString("- ")
		parseDocxBlockText(sb, b.Bullet)
	case lark_docx.DocxBlockTypeOrdered:
		sb.WriteString("1. ")
		parseDocxBlockText(sb, b.Ordered)
	case lark_docx.DocxBlockTypeCode:
		sb.WriteString("```" + b.Code.Style.Language.String() + "\n")
		parseDocxBlockText(sb, b.Code)
		sb.WriteString("\n```")
	case lark_docx.DocxBlockTypeQuote:
		sb.WriteString("> ")
		parseDocxBlockText(sb, b.Quote)
	case lark_docx.DocxBlockTypeEquation:
		sb.WriteString("$$\n")
		parseDocxBlockText(sb, b.Equation)
		sb.WriteString("\n$$")
	case lark_docx.DocxBlockTypeTodo:
		if *b.Todo.Style.Done {
			sb.WriteString("- [x] ")
		} else {
			sb.WriteString("- [ ] ")
		}
		parseDocxBlockText(sb, b.Todo)
	case lark_docx.DocxBlockTypeImage:
		p.parseBlockImage(b.Image)
	case lark_docx.DocxBlockTypeTable:
		p.parseBlockTable(b)
		b.Children = nil
	case lark_docx.DocxBlockTypeQuoteContainer:
		p.parseBlockQuoteContainer(b)
		b.Children = nil
	}
	sb.WriteString(separator)
	for _, blockId := range b.Children {
		p.parseSingleBlock(p.blockMap[blockId], indentLevel+1)
	}
}

func (p *docxParser) parseBlockImage(img *lark_docx.Image) {
	p.sb.WriteString("![](")
	p.sb.WriteString(img.Token)
	p.sb.WriteString(")")
	p.ImgTokens = append(p.ImgTokens, img.Token)
}

func (p *docxParser) parseBlockTable(b *lark_docx.Block) {
	var rows [][]string
	for i, blockId := range b.Table.Cells {
		content := p.parseBlockTableCell(p.blockMap[blockId])
		rowIndex := i / b.Table.Property.ColumnSize
		if len(rows) < int(rowIndex)+1 {
			rows = append(rows, []string{})
		}
		rows[rowIndex] = append(rows[rowIndex], content)
	}

	p.sb.WriteString(renderMarkdownTable(rows))
}

func (p *docxParser) parseBlockQuoteContainer(b *lark_docx.Block) {
	p.sb.WriteString("> ")
	for _, childBlockId := range b.Children {
		parseDocxBlockText(p.sb, p.blockMap[childBlockId].Text)
	}
}

func (p *docxParser) parseBlockTableCell(b *lark_docx.Block) string {
	sb := getStringBuilder()
	defer putStringBuilder(sb)
	parseDocxBlockText(sb, p.blockMap[b.Children[0]].Text)
	return sb.String()
}

func parseDocxBlockText(sb *strings.Builder, t *lark_docx.Text) {
	if t == nil {
		return
	}
	for _, e := range t.Elements {
		parseDocxTextElement(sb, e)
	}
}

func parseDocxTextElement(sb *strings.Builder, e *lark_docx.TextElement) {
	if e.TextRun != nil {
		parseDocxTextElementTextRun(sb, e.TextRun)
	}
	if e.MentionUser != nil {
		sb.WriteString(*e.MentionUser.UserId)
	}
	if e.MentionDoc != nil {
		sb.WriteString("[")
		sb.WriteString(*e.MentionDoc.Title)
		sb.WriteString("](")
		sb.WriteString(urlx.UnescapeURL(*e.MentionDoc.Url))
		sb.WriteString(")")
	}
	if e.Equation != nil {
		sb.WriteString("$$" + strings.TrimSuffix(*e.Equation.Content, "\n") + "$$")
	}
}

func parseDocxTextElementTextRun(sb *strings.Builder, tr *lark_docx.TextRun) {
	postWrite := ""
	if style := tr.TextElementStyle; style != nil {
		if style.Bold {
			sb.WriteString("**")
			postWrite = "**"
		} else if style.Italic {
			sb.WriteString("*")
			postWrite = "*"
		} else if style.Strikethrough {
			sb.WriteString("~~")
			postWrite = "~~"
		} else if style.Underline {
			// ignore underline
		} else if style.InlineCode {
			sb.WriteString("`")
			postWrite = "`"
		} else if link := style.Link; link != nil {
			sb.WriteString("[")
			postWrite = fmt.Sprintf("](%s)", urlx.UnescapeURL(*link.Url))
		}
	}
	sb.WriteString(tr.Content)
	sb.WriteString(postWrite)
}

func renderMarkdownTable(data [][]string) string {
	sb := getStringBuilder()
	defer putStringBuilder(sb)
	table := tablewriter.NewWriter(sb)
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetAutoMergeCells(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetHeader(data[0])
	table.AppendBulk(data[1:])
	table.Render()
	return sb.String()
}
