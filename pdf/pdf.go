package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

// PDF is a struct that saves information about pdf and his cursor position
type PDF struct {
	Fpdf *gofpdf.Fpdf
	X, Y float64
}

// Move moves the currrent cursor to a specific position from the current one
func (p *PDF) Move(xDelta, yDelta float64) {
	p.X, p.Y = p.X+xDelta, p.Y+yDelta
	p.Fpdf.MoveTo(p.X, p.Y)
}

// MoveAbs moves the current pdf cursor into a specific position
func (p *PDF) MoveAbs(x, y float64) {
	p.X, p.Y = x, y
	p.Fpdf.MoveTo(p.X, p.Y)
}

// DrawHeader draw a text in header format with a specific size, this size comes
// from the standar definition of the six different headers in markdown
func (p *PDF) DrawHeader(text string, size float64) {
	p.Fpdf.SetFont("times", "B", size)
	_, lineHt := p.Fpdf.GetFontSize()
	p.Fpdf.WriteAligned(0, lineHt, text, gofpdf.AlignLeft)
	p.Move(0, lineHt*2.0)
}

// DrawText draw a text and move to the next line, text variable is the value of
// the text and the style can be "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out)
// or any combination. The default value (specified with an empty string) is
// regular. Bold and italic styles do not apply to Symbol and ZapfDingbats.
func (p *PDF) DrawText(text string, style string) {
	p.Fpdf.SetFont("times", style, 10)
	_, lineHt := p.Fpdf.GetFontSize()
	p.Fpdf.WriteAligned(0, lineHt, text, gofpdf.AlignLeft)
	p.Move(0, lineHt*2.0)
}

// OutputFile just call the specific call from the library for output and create the pdf file
func (p *PDF) OutputFile(fileName string) error {
	return p.Fpdf.OutputFileAndClose(fileName)
}
