package parser

import (
	"fmt"

	"github.com/ramonmacias/markdown-to-pdf/markdown"
	"github.com/ramonmacias/markdown-to-pdf/markdown/blackfriday"
	"github.com/ramonmacias/markdown-to-pdf/markdown/gomarkdown"
)

type Option struct {
	InputFileName  string
	OutputFileName string
	ParserProvider string
}

func Parse(option Option) (err error) {
	var md markdown.MarkdownParser

	switch option.ParserProvider {
	case "blackfriday":
		md, err = blackfriday.ParseFile(option.InputFileName + ".md")
		if err != nil {
			return err
		}
	case "gomarkdown":
		md, err = gomarkdown.ParseFile(option.InputFileName + ".md")
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Parser provider not allowed")
	}

	pdfFile, err := md.ConvertToPDF()
	if err != nil {
		return err
	}

	err = pdfFile.OutputFile(option.OutputFileName + ".pdf")
	if err != nil {
		return err
	}
	return nil
}
