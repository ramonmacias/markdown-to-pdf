package markdown

import (
	"github.com/ramonmacias/markdown-to-pdf/pdf"
)

// MarkdownParser define the behaviour of every different markdown parser
type MarkdownParser interface {
	ConvertToPDF() (pdfFile pdf.PDF, err error)
}
