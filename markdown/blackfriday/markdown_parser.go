package blackfriday

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/ramonmacias/markdown-to-pdf/pdf"
	"github.com/ramonmacias/markdown-to-pdf/pdf/gofpdf"
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

// GetSizeHeader give use the size depending of which kind of header we ask for
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

// Markdown is the struct that defines what is a markdown file
type Markdown struct {
	root *blackfriday.Node
}

// ParseFile from a filename we create a new Markdown object
func ParseFile(fileName string) (markdown Markdown, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return Parse(f)
}

// Parse from a reader interface we create a new Markdown object
func Parse(r io.Reader) (markdown Markdown, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return markdown, err
	}

	mkd := blackfriday.New()
	markdown.root = mkd.Parse(b)
	return markdown, nil
}

// ConvertToPDF go through all the tree with the information of this markdown file
// and generate a pdf file trying to keep the same information and format
func (m Markdown) ConvertToPDF() (pdfFile pdf.PdfParser, err error) {
	pdfFile = gofpdf.DefaultPDF()

	// We go through the tree that represents the markdown file
	m.root.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if node.Type == blackfriday.Heading {
			pdfFile.DrawHeader(string(node.FirstChild.Literal), GetSizeHeader(node.HeadingData.Level))
		} else if node.Type == blackfriday.List {
			pdfFile.DrawText(" - "+string(node.FirstChild.Literal), "")
		} else if node.Type == blackfriday.Text && node.Parent.Type != blackfriday.Heading && node.Parent.Type != blackfriday.List {
			if node.Parent.Type == blackfriday.Strong {
				pdfFile.DrawText(string(node.Literal), "B")
			} else if node.Parent.Type == blackfriday.Emph {
				pdfFile.DrawText(string(node.Literal), "I")
			} else {
				pdfFile.DrawText(string(node.Literal), "")
			}
		}
		return blackfriday.GoToNext
	})

	return pdfFile, err
}
