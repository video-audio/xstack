package xstack

import (
	"errors"
	"io"
	"math"
	"strconv"
	"strings"
)

const (
	// special case: zero offset
	// with no byte prefix
	offsetZeroSign = '0'

	// sum separator between offsets on same axis
	offsetSumSign = '+'

	// separator between two total offsets for
	// different axis
	offsetsSep = '_'

	// separator between two cells
	cellsSep = '|'

	// widht char prefix. width for x axis.
	wChar = 'w'

	// htight char prefix. height for y axis.
	hChar = 'h'
)

// ErrZeroN indicates that passed n to Layout function is 0
// and cann't be processed.
var ErrZeroN = errors.New("layout for 0 elements couldn't be calculated")

// writeOffsetTo writes offset to w.
// Produces data like this:
//   0
//   w0
//   h0+h1+h2+h3
func writeOffsetTo(w io.Writer, c byte, pos uint64) error {
	if pos == 0 {
		_, err := w.Write([]byte{offsetZeroSign})
		return err
	}

	for i := uint64(0); i < pos; i++ {
		if _, err := w.Write([]byte{c}); err != nil {
			return err
		}

		if _, err := w.Write([]byte(strconv.FormatUint(i, 10))); err != nil {
			return err
		}

		// if not last element
		if i != pos-1 {
			if _, err := w.Write([]byte{offsetSumSign}); err != nil {
				return err
			}
		}
	}

	return nil
}

// LayoutTo writes xstack layout description to w.
// Where n is total number of elements in grid.
func LayoutTo(w io.Writer, n uint64) error {
	if n == 0 {
		return ErrZeroN
	}

	cols := uint64(math.Trunc(math.Sqrt(float64(n))))
	rows := uint64(math.Ceil(float64(n) / float64(cols)))

	// remaining items to place in grid
	remaining := n

	for row := uint64(0); row < rows; row++ {
		for col := uint64(0); col < cols && remaining != 0; col++ {
			if err := writeOffsetTo(w, wChar, col); err != nil {
				return err
			}

			if _, err := w.Write([]byte{offsetsSep}); err != nil {
				return err
			}

			if err := writeOffsetTo(w, hChar, row); err != nil {
				return err
			}

			remaining--

			if remaining != 0 {
				if _, err := w.Write([]byte{cellsSep}); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Layout returns xstack layout description as string.
// Is's just helper function - wrapper around LayoutTo.
func Layout(n uint64) (string, error) {
	var b strings.Builder

	if err := LayoutTo(&b, n); err != nil {
		return "", err
	}

	return b.String(), nil
}
