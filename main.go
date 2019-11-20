package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/ramonmacias/markdown-to-pdf/markdown"
	"github.com/ramonmacias/markdown-to-pdf/pdf"
	"gopkg.in/russross/blackfriday.v2"
)

func main() {

	fpdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	fpdf.AddPage()
	pdf := pdf.PDF{
		Fpdf: fpdf,
	}

	f, err := os.Open("samples/HEADERS.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fpdf.SetFont("times", "B", 10)
	fpdf.SetTextColor(50, 50, 50)
	pdf.MoveAbs(0, 100)

	mkd := blackfriday.New()
	root := mkd.Parse(b)
	root.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if node.Type == blackfriday.Text && node.Parent.Type != blackfriday.Heading {
			if node.Parent.Type == blackfriday.Strong {
				pdf.DrawText(string(node.Literal), "B")
			} else if node.Parent.Type == blackfriday.Emph {
				pdf.DrawText(string(node.Literal), "I")
			} else {
				log.Println(node.Parent.Type)
				pdf.DrawText(string(node.Literal), "")
			}
		} else if node.Type == blackfriday.Heading {
			pdf.DrawHeader(string(node.FirstChild.Literal), markdown.GetSizeHeader(node.HeadingData.Level))
		}
		return blackfriday.GoToNext
	})

	err = fpdf.OutputFileAndClose("sample.pdf")
	if err != nil {
		panic(err)
	}
}
