package xstack

type Align uint8

const (
	// For 6 inputs will produce 2x3 grid. e.g:
	// | | |
	// | | |
	// | | |
	AlignVertical Align = iota

	// For 6 inputs will produce 3x2 grid. e.g:
	// | | | |
	// | | | |
	AlignHorizontal
)
