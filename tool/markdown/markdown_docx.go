package lark_markdown

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/elliotchance/orderedmap"
	"github.com/olekukonko/tablewriter"
	"github.com/zzzzer91/gopkg/urlx"
	lark_docx "github.com/zzzzer91/lark-go/service/docx/v1"
)

type Parser struct {
	ImgTokens []string // 所有图片的 token
	blockMap  *orderedmap.OrderedMap
}

func NewParser() *Parser {
	return &Parser{blockMap: orderedmap.NewOrderedMap()}
}

var stringBuilderPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func getStringBuilder() *strings.Builder {
	return stringBuilderPool.Get().(*strings.Builder)
}

func putStringBuilder(sb *strings.Builder) {
	sb.Reset()
	stringBuilderPool.Put(sb)
}

func (p *Parser) ParseDocxContent(blocks []*lark_docx.Block) string {
	// block map
	// - Table cell block needs block map to collect children blocks
	// - ParseDocxContent needs block map to avoid duplicate rendering
	for _, block := range blocks {
		p.blockMap.Set(block.BlockId, block)
	}

	buf := getStringBuilder()
	defer putStringBuilder(buf)
	// buf.WriteString(p.ParseDocxDocument(doc))
	// buf.WriteString("\n")
	for _, v := range blocks {
		s := p.ParseDocxBlock(v)
		if s != "" {
			buf.WriteString(s)
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

func (p *Parser) ParseDocxBlock(b *lark_docx.Block) string {
	if _, ok := p.blockMap.Get(b.BlockId); p.blockMap != nil && !ok {
		// ignore rendered children block
		return ""
	}

	buf := getStringBuilder()
	defer putStringBuilder(buf)
	switch b.BlockType {
	case lark_docx.DocxBlockTypePage:
		buf.WriteString("# ")
		buf.WriteString(p.ParseDocxBlockText(b.Page))
	case lark_docx.DocxBlockTypeText:
		return p.ParseDocxBlockText(b.Text)
	case lark_docx.DocxBlockTypeHeading1:
		buf.WriteString("## ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading1))
	case lark_docx.DocxBlockTypeHeading2:
		buf.WriteString("### ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading2))
	case lark_docx.DocxBlockTypeHeading3:
		buf.WriteString("#### ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading3))
	case lark_docx.DocxBlockTypeHeading4:
		buf.WriteString("##### ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading4))
	case lark_docx.DocxBlockTypeHeading5:
		buf.WriteString("###### ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading5))
	case lark_docx.DocxBlockTypeHeading6:
		buf.WriteString("####### ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading6))
	case lark_docx.DocxBlockTypeHeading7:
		buf.WriteString("######## ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading7))
	case lark_docx.DocxBlockTypeHeading8:
		buf.WriteString("######### ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading8))
	case lark_docx.DocxBlockTypeHeading9:
		buf.WriteString("########## ")
		buf.WriteString(p.ParseDocxBlockText(b.Heading9))
	case lark_docx.DocxBlockTypeBullet:
		// calculate indent level
		indentLevel := 1
		parent := p.blockMap.GetOrDefault(b.ParentId, nil)
		for {
			if parent == nil || (parent.(*lark_docx.Block).BlockType != lark_docx.DocxBlockTypeBullet &&
				parent.(*lark_docx.Block).BlockType != lark_docx.DocxBlockTypeOrdered) {
				break
			}
			indentLevel += 1
			parent = p.blockMap.GetOrDefault(parent.(*lark_docx.Block).ParentId, nil)
		}
		buf.WriteString(strings.Repeat("    ", indentLevel-1))
		buf.WriteString("- ")
		buf.WriteString(p.ParseDocxBlockText(b.Bullet))
	case lark_docx.DocxBlockTypeOrdered:
		// calculate indent level
		indentLevel := 1
		parent := p.blockMap.GetOrDefault(b.ParentId, nil)
		for {
			if parent == nil || (parent.(*lark_docx.Block).BlockType != lark_docx.DocxBlockTypeBullet &&
				parent.(*lark_docx.Block).BlockType != lark_docx.DocxBlockTypeOrdered) {
				break
			}
			indentLevel += 1
			parent = p.blockMap.GetOrDefault(parent.(*lark_docx.Block).ParentId, nil)
		}
		buf.WriteString(strings.Repeat("    ", indentLevel-1))
		buf.WriteString("1. ")
		buf.WriteString(p.ParseDocxBlockText(b.Ordered))
	case lark_docx.DocxBlockTypeCode:
		buf.WriteString("```" + b.Code.Style.Language.String() + "\n")
		buf.WriteString(strings.TrimSpace(p.ParseDocxBlockText(b.Code)))
		buf.WriteString("\n```")
		buf.WriteString("\n")
	case lark_docx.DocxBlockTypeQuote:
		buf.WriteString("> ")
		buf.WriteString(p.ParseDocxBlockText(b.Quote))
	case lark_docx.DocxBlockTypeEquation:
		buf.WriteString("$$\n")
		buf.WriteString(p.ParseDocxBlockText(b.Equation))
		buf.WriteString("\n$$")
	case lark_docx.DocxBlockTypeTodo:
		if *b.Todo.Style.Done {
			buf.WriteString("- [x] ")
		} else {
			buf.WriteString("- [ ] ")
		}
		buf.WriteString(p.ParseDocxBlockText(b.Todo))
	case lark_docx.DocxBlockTypeImage:
		buf.WriteString(p.ParseDocxBlockImage(b.Image))
	case lark_docx.DocxBlockTypeTableCell:
		buf.WriteString(p.ParseDocxBlockTableCell(b.BlockId))
	case lark_docx.DocxBlockTypeTable:
		buf.WriteString(p.ParseDocxBlockTable(b.ParentId, b.Table))
	case lark_docx.DocxBlockTypeQuoteContainer:
		buf.WriteString(p.ParseDocxBlockQuoteContainer(b.BlockId, b.QuoteContainer))
	default:
		return ""
	}
	return buf.String()
}

func (p *Parser) ParseDocxBlockText(b *lark_docx.Text) string {
	buf := getStringBuilder()
	defer putStringBuilder(buf)
	for _, e := range b.Elements {
		buf.WriteString(p.ParseDocxTextElement(e))
	}
	buf.WriteString("\n")
	return buf.String()
}

func (p *Parser) ParseDocxTextElement(e *lark_docx.TextElement) string {
	buf := getStringBuilder()
	defer putStringBuilder(buf)
	if e.TextRun != nil {
		buf.WriteString(p.ParseDocxTextElementTextRun(e.TextRun))
	}
	if e.MentionUser != nil {
		buf.WriteString(*e.MentionUser.UserId)
	}
	if e.MentionDoc != nil {
		buf.WriteString(fmt.Sprintf("[%s](%s)", *e.MentionDoc.Title, urlx.UnescapeURL(*e.MentionDoc.Url)))
	}
	if e.Equation != nil {
		buf.WriteString("$$" + strings.TrimSuffix(*e.Equation.Content, "\n") + "$$")
	}
	return buf.String()
}

func (p *Parser) ParseDocxTextElementTextRun(tr *lark_docx.TextRun) string {
	buf := getStringBuilder()
	defer putStringBuilder(buf)
	postWrite := ""
	if style := tr.TextElementStyle; style != nil {
		if style.Bold {
			buf.WriteString("**")
			postWrite = "**"
		} else if style.Italic {
			buf.WriteString("*")
			postWrite = "*"
		} else if style.Strikethrough {
			buf.WriteString("~~")
			postWrite = "~~"
		} else if style.Underline {
			// ignore underline
		} else if style.InlineCode {
			buf.WriteString("`")
			postWrite = "`"
		} else if link := style.Link; link != nil {
			buf.WriteString("[")
			postWrite = fmt.Sprintf("](%s)", urlx.UnescapeURL(*link.Url))
		}
	}
	buf.WriteString(tr.Content)
	buf.WriteString(postWrite)
	return buf.String()
}

func (p *Parser) ParseDocxBlockImage(img *lark_docx.Image) string {
	buf := getStringBuilder()
	defer putStringBuilder(buf)
	buf.WriteString(fmt.Sprintf("![](%s)", img.Token))
	buf.WriteString("\n")
	p.ImgTokens = append(p.ImgTokens, img.Token)
	return buf.String()
}

func (p *Parser) ParseDocxBlockTableCell(blockId string) string {
	var contents string
	for _, key := range p.blockMap.Keys() {
		value, ok := p.blockMap.Get(key)
		if !ok {
			continue
		}
		block := value.(*lark_docx.Block)
		if block.ParentId != blockId {
			continue
		}

		content := p.ParseDocxBlock(block)
		if content == "" {
			continue
		}
		contents += content
		// remove table cell children block from map
		p.blockMap.Delete(block.BlockId)
	}
	return strings.TrimSpace(contents)
}

func (p *Parser) ParseDocxBlockTable(documentId string, t *lark_docx.Table) string {
	// - First row as header
	// - Ignore cell merging
	var rows [][]string
	for i, blockId := range t.Cells {
		block, ok := p.blockMap.Get(blockId)
		if !ok {
			log.Printf("got invalid block cell '%s', document: %s\n", blockId, documentId)
			continue
		}

		content := p.ParseDocxBlock(block.(*lark_docx.Block))
		rowIndex := i / t.Property.ColumnSize
		if len(rows) < int(rowIndex)+1 {
			rows = append(rows, []string{})
		}
		rows[rowIndex] = append(rows[rowIndex], content)
		// remove table cell block from map
		p.blockMap.Delete(blockId)
	}

	buf := getStringBuilder()
	defer putStringBuilder(buf)
	buf.WriteString(renderMarkdownTable(rows))
	return buf.String()
}

func (p *Parser) ParseDocxBlockQuoteContainer(blockId string, q *lark_docx.QuoteContainer) string {
	contents := "> "
	for _, key := range p.blockMap.Keys() {
		value, ok := p.blockMap.Get(key)
		if !ok {
			continue
		}
		block := value.(*lark_docx.Block)
		if block.ParentId != blockId {
			continue
		}

		content := p.ParseDocxBlock(block)
		if content == "" {
			continue
		}
		contents += content
		// remove quote container children block from map
		p.blockMap.Delete(block.BlockId)
	}
	return contents
}

func renderMarkdownTable(data [][]string) string {
	builder := &strings.Builder{}
	table := tablewriter.NewWriter(builder)
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetAutoMergeCells(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetHeader(data[0])
	table.AppendBulk(data[1:])
	table.Render()
	return builder.String()
}
