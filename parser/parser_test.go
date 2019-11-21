package parser_test

import (
	"testing"

	"github.com/ramonmacias/markdown-to-pdf/parser"
)

func TestMarkdownSamplesBlackFriday(t *testing.T) {
	err := parser.Parse(parser.Option{
		InputFileName:  "testdata/markdown/MARKDOWN_SAMPLES",
		OutputFileName: "testdata/pdf/MARKDOWN_SAMPLES_CONVERTED",
		ParserProvider: "blackfriday",
	})

	if err != nil {
		t.Errorf("Shouldn't be an error but got %v", err)
	}
}
