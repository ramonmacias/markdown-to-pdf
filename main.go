package main

import (
	"flag"

	"github.com/ramonmacias/markdown-to-pdf/parser"
)

// Flags for setup the program
var (
	inputFile      = flag.String("input-file", "README", "Use this flag to setup an entry file name")
	outputFile     = flag.String("output-file", "output", "Use this flag to setup an output file name")
	parserProvider = flag.String("parser-provider", "blackfriday", "Define which markdown parser provider want to use")
)

func main() {
	flag.Parse()

	err := parser.Parse(parser.Option{
		InputFileName:  *inputFile,
		OutputFileName: *outputFile,
		ParserProvider: *parserProvider,
	})

	if err != nil {
		panic(err)
	}
}
