package markdown

// All the constants related with the size of the headers in markdown
const (
	headerH1 = 24
	headerH2 = 18
	headerH3 = 14
	headerH4 = 12
	headerH5 = 10
	headerH6 = 8
)

func GetSizeHeader(level int) float64 {
	switch level {
	case 1:
		return headerH1
	case 2:
		return headerH2
	case 3:
		return headerH3
	case 4:
		return headerH4
	case 5:
		return headerH5
	case 6:
		return headerH6
	default:
		panic("No header allowed")
	}
}
