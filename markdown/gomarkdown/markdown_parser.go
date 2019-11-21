package gomarkdown

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/ramonmacias/markdown-to-pdf/pdf"
	"github.com/ramonmacias/markdown-to-pdf/pdf/pdfshift"
)

// Markdown is the struct that defines what is a markdown file
type Markdown struct {
	mdToHTML []byte
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
func Parse(r io.Reader) (md Markdown, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return md, err
	}

	md.mdToHTML = markdown.ToHTML(b, nil, nil)
	return md, nil
}

// ConvertToPDF just create a new pdf parser with the information need it to build
// a new PDF on that case using an html file
func (m Markdown) ConvertToPDF() (pdfFile pdf.PdfParser, err error) {
	pdfFile = pdfshift.PDF{Html: m.mdToHTML}
	return pdfFile, err
}
