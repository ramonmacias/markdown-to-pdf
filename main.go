package main

import (
	"github.com/ramonmacias/markdown-to-pdf/markdown"
)

func main() {

	md, err := markdown.ParseFile("samples/HEADERS.md")
	if err != nil {
		panic(err)
	}

	pdfFile, err := md.ConvertToPDF()
	if err != nil {
		panic(err)
	}

	err = pdfFile.OutputFile("sample.pdf")
	if err != nil {
		panic(err)
	}
}
