package pdf

// PdfParser is an interface that describes the basic behaviour for each pdf parser
type PdfParser interface {
	OutputFile(fileName string) error
	DrawHeader(text string, size float64)
	DrawText(text string, style string)
}
