package xstack

// Align sets alignment of grid.
type Align uint8

const (
	// AlignVertical is for vertical grid emission.
	// For 6 inputs will produce 2x3 grid. e.g:
	// | | |
	// | | |
	// | | |
	AlignVertical Align = iota

	// AlignHorizontal is for horizontal grid emission.
	// For 6 inputs will produce 3x2 grid. e.g:
	// | | | |
	// | | | |
	AlignHorizontal
)
