package markdown

import (
	"io"

	"github.com/ramonmacias/markdown-to-pdf/pdf"
	"gopkg.in/russross/blackfriday.v2"
)

// All the constants related with the size of the headers in markdown
const (
	headerH1 = 24
	headerH2 = 18
	headerH3 = 14
	headerH4 = 12
	headerH5 = 10
	headerH6 = 8
)

func GetSizeHeader(level int) float64 {
	switch level {
	case 1:
		return headerH1
	case 2:
		return headerH2
	case 3:
		return headerH3
	case 4:
		return headerH4
	case 5:
		return headerH5
	case 6:
		return headerH6
	default:
		panic("No header allowed")
	}
}

type Markdown struct {
	md blackfriday.Markdown
}

func Parse(r io.Reader) Markdown {
	mkd := blackfriday.New()
	return mkd.Parse(r)
}

func (m *Markdown) ConvertToPDF() (pdf pdf.PDF, err error) {
	return pdf, err
}
