package markdown

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
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
	root *blackfriday.Node
}

func ParseFile(fileName string) (markdown Markdown, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return Parse(f)
}

func Parse(r io.Reader) (markdown Markdown, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return markdown, err
	}

	mkd := blackfriday.New()
	markdown.root = mkd.Parse(b)
	return markdown, nil
}

func (m *Markdown) ConvertToPDF() (pdfFile pdf.PDF, err error) {
	fpdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	fpdf.AddPage()
	pdfFile = pdf.PDF{
		Fpdf: fpdf,
	}

	fpdf.SetFont("times", "B", 10)
	fpdf.SetTextColor(50, 50, 50)
	pdfFile.MoveAbs(0, 100)

	m.root.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if node.Type == blackfriday.Text && node.Parent.Type != blackfriday.Heading {
			if node.Parent.Type == blackfriday.Strong {
				pdfFile.DrawText(string(node.Literal), "B")
			} else if node.Parent.Type == blackfriday.Emph {
				pdfFile.DrawText(string(node.Literal), "I")
			} else {
				log.Println(node.Parent.Type)
				pdfFile.DrawText(string(node.Literal), "")
			}
		} else if node.Type == blackfriday.Heading {
			pdfFile.DrawHeader(string(node.FirstChild.Literal), GetSizeHeader(node.HeadingData.Level))
		}
		return blackfriday.GoToNext
	})

	return pdfFile, err
}
