package parser_test

import (
	"testing"

	"github.com/ramonmacias/markdown-to-pdf/parser"
)

// TEST FOR USING THE FIRST OPTION blackfriday + fgopdf

func TestMarkdownSamplesBlackFriday(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_SAMPLES",
		OutputFileName: "testdata/pdf/blackfriday/MARKDOWN_SAMPLES_CONVERTED",
		ParserProvider: "blackfriday",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownCodeBlackFriday(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_CODE",
		OutputFileName: "testdata/pdf/blackfriday/MARKDOWN_CODE_CONVERTED",
		ParserProvider: "blackfriday",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownImagesBlackFriday(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_IMAGE",
		OutputFileName: "testdata/pdf/blackfriday/MARKDOWN_IMAGE_CONVERTED",
		ParserProvider: "blackfriday",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownLinksBlackFriday(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_LINKS",
		OutputFileName: "testdata/pdf/blackfriday/MARKDOWN_LINKS_CONVERTED",
		ParserProvider: "blackfriday",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownListBlackFriday(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_LIST",
		OutputFileName: "testdata/pdf/blackfriday/MARKDOWN_LIST_CONVERTED",
		ParserProvider: "blackfriday",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownTablesBlackFriday(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_TABLES",
		OutputFileName: "testdata/pdf/blackfriday/MARKDOWN_TABLES_CONVERTED",
		ParserProvider: "blackfriday",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

// TEST FOR USING THE SECOND OPTION gomarkdown + pdfshift
func TestMarkdownSamplesGomarkdown(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_SAMPLES",
		OutputFileName: "testdata/pdf/gomarkdown/MARKDOWN_SAMPLES_CONVERTED",
		ParserProvider: "gomarkdown",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownCodeGomarkdown(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_CODE",
		OutputFileName: "testdata/pdf/gomarkdown/MARKDOWN_CODE_CONVERTED",
		ParserProvider: "gomarkdown",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownImagesGomarkdown(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_IMAGE",
		OutputFileName: "testdata/pdf/gomarkdown/MARKDOWN_IMAGE_CONVERTED",
		ParserProvider: "gomarkdown",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownLinksGomarkdown(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_LINKS",
		OutputFileName: "testdata/pdf/gomarkdown/MARKDOWN_LINKS_CONVERTED",
		ParserProvider: "gomarkdown",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownListGomarkdown(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_LIST",
		OutputFileName: "testdata/pdf/gomarkdown/MARKDOWN_LIST_CONVERTED",
		ParserProvider: "gomarkdown",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}

func TestMarkdownTablesGomarkdown(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_TABLES",
		OutputFileName: "testdata/pdf/gomarkdown/MARKDOWN_TABLES_CONVERTED",
		ParserProvider: "gomarkdown",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}
