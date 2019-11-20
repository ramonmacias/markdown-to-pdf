package main

import (
	"flag"

	"github.com/ramonmacias/markdown-to-pdf/markdown"
	"github.com/ramonmacias/markdown-to-pdf/markdown/blackfriday"
)

// Flags for setup the program
var (
	inputFile      = flag.String("input-file", "samples/MARKDOWN_SAMPLES", "Use this flag to setup an entry file name")
	outputFile     = flag.String("output-file", "output", "Use this flag to setup an output file name")
	parserProvider = flag.String("parser-provider", "blackfriday", "Define which markdown parser provider want to use")
)

func main() {
	flag.Parse()
	var md markdown.MarkdownParser
	var err error

	switch *parserProvider {
	case "blackfriday":
		md, err = blackfriday.ParseFile(*inputFile + ".md")
		if err != nil {
			panic(err)
		}
	default:
		panic("Parser provider not allowed")
	}

	pdfFile, err := md.ConvertToPDF()
	if err != nil {
		panic(err)
	}

	err = pdfFile.OutputFile(*outputFile + ".pdf")
	if err != nil {
		panic(err)
	}
}
