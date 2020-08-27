package xstack_test

import (
	"errors"
	"testing"

	"github.com/video-audio/xstack"
)

func TestLayout(t *testing.T) {
	tests := []struct {
		n int

		wlayout string
		werr    error
	}{
		{0, "", xstack.ErrZeroN},
		{1, "0_0", nil},
		{2, "0_0|0_h0", nil},
		{3, "0_0|0_h0|0_h0+h1", nil},
		{4, "0_0|w0_0|0_h0|w0_h0", nil},
		{8, "0_0|w0_0|0_h0|w0_h0|0_h0+h1|w0_h0+h1|0_h0+h1+h2|w0_h0+h1+h2", nil},

		// see https://trac.ffmpeg.org/wiki/Create%20a%20mosaic%20out%20of%20several%20input%20videos%20using%20xstack
		{9, "0_0|w0_0|w0+w1_0|0_h0|w0_h0|w0+w1_h0|0_h0+h1|w0_h0+h1|w0+w1_h0+h1", nil},

		{10, "0_0|w0_0|w0+w1_0|0_h0|w0_h0|w0+w1_h0|0_h0+h1|w0_h0+h1|w0+w1_h0+h1|0_h0+h1+h2", nil},

		{15,
			"0_0|w0_0|w0+w1_0|0_h0|w0_h0|w0+w1_h0|0_h0+h1" +
				"|w0_h0+h1|w0+w1_h0+h1|0_h0+h1+h2|w0_h0+h1+h2" +
				"|w0+w1_h0+h1+h2|0_h0+h1+h2+h3|w0_h0+h1+h2+h3" +
				"|w0+w1_h0+h1+h2+h3",
			nil},

		// see https://trac.ffmpeg.org/wiki/Create%20a%20mosaic%20out%20of%20several%20input%20videos%20using%20xstack
		{36,
			"0_0" +
				"|w0_0" +
				"|w0+w1_0" +
				"|w0+w1+w2_0" +
				"|w0+w1+w2+w3_0" +
				"|w0+w1+w2+w3+w4_0" +
				"|0_h0" +
				"|w0_h0" +
				"|w0+w1_h0" +
				"|w0+w1+w2_h0" +
				"|w0+w1+w2+w3_h0" +
				"|w0+w1+w2+w3+w4_h0" +
				"|0_h0+h1" +
				"|w0_h0+h1" +
				"|w0+w1_h0+h1" +
				"|w0+w1+w2_h0+h1" +
				"|w0+w1+w2+w3_h0+h1" +
				"|w0+w1+w2+w3+w4_h0+h1" +
				"|0_h0+h1+h2" +
				"|w0_h0+h1+h2" +
				"|w0+w1_h0+h1+h2" +
				"|w0+w1+w2_h0+h1+h2" +
				"|w0+w1+w2+w3_h0+h1+h2" +
				"|w0+w1+w2+w3+w4_h0+h1+h2" +
				"|0_h0+h1+h2+h3" +
				"|w0_h0+h1+h2+h3" +
				"|w0+w1_h0+h1+h2+h3" +
				"|w0+w1+w2_h0+h1+h2+h3" +
				"|w0+w1+w2+w3_h0+h1+h2+h3" +
				"|w0+w1+w2+w3+w4_h0+h1+h2+h3" +
				"|0_h0+h1+h2+h3+h4" +
				"|w0_h0+h1+h2+h3+h4" +
				"|w0+w1_h0+h1+h2+h3+h4" +
				"|w0+w1+w2_h0+h1+h2+h3+h4" +
				"|w0+w1+w2+w3_h0+h1+h2+h3+h4" +
				"|w0+w1+w2+w3+w4_h0+h1+h2+h3+h4",
			nil},

		{100,
			"0_0" +
				"|w0_0" +
				"|w0+w1_0" +
				"|w0+w1+w2_0" +
				"|w0+w1+w2+w3_0" +
				"|w0+w1+w2+w3+w4_0" +
				"|w0+w1+w2+w3+w4+w5_0" +
				"|w0+w1+w2+w3+w4+w5+w6_0" +
				"|w0+w1+w2+w3+w4+w5+w6+w7_0" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_0" +
				"|0_h0" +
				"|w0_h0" +
				"|w0+w1_h0" +
				"|w0+w1+w2_h0" +
				"|w0+w1+w2+w3_h0" +
				"|w0+w1+w2+w3+w4_h0" +
				"|w0+w1+w2+w3+w4+w5_h0" +
				"|w0+w1+w2+w3+w4+w5+w6_h0" +
				"|w0+w1+w2+w3+w4+w5+w6+w7_h0" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0" +
				"|0_h0+h1" +
				"|w0_h0+h1" +
				"|w0+w1_h0+h1" +
				"|w0+w1+w2_h0+h1" +
				"|w0+w1+w2+w3_h0+h1" +
				"|w0+w1+w2+w3+w4_h0+h1" +
				"|w0+w1+w2+w3+w4+w5_h0+h1" +
				"|w0+w1+w2+w3+w4+w5+w6_h0+h1" +
				"|w0+w1+w2+w3+w4+w5+w6+w7_h0+h1|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0+h1" +
				"|0_h0+h1+h2|w0_h0+h1+h2|w0+w1_h0+h1+h2|w0+w1+w2_h0+h1+h2" +
				"|w0+w1+w2+w3_h0+h1+h2|w0+w1+w2+w3+w4_h0+h1+h2|w0+w1+w2+w3+w4+w5_h0+h1+h2" +
				"|w0+w1+w2+w3+w4+w5+w6_h0+h1+h2|w0+w1+w2+w3+w4+w5+w6+w7_h0+h1+h2" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0+h1+h2|0_h0+h1+h2+h3|w0_h0+h1+h2+h3" +
				"|w0+w1_h0+h1+h2+h3|w0+w1+w2_h0+h1+h2+h3|w0+w1+w2+w3_h0+h1+h2+h3" +
				"|w0+w1+w2+w3+w4_h0+h1+h2+h3|w0+w1+w2+w3+w4+w5_h0+h1+h2+h3" +
				"|w0+w1+w2+w3+w4+w5+w6_h0+h1+h2+h3|w0+w1+w2+w3+w4+w5+w6+w7_h0+h1+h2+h3" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0+h1+h2+h3|0_h0+h1+h2+h3+h4" +
				"|w0_h0+h1+h2+h3+h4|w0+w1_h0+h1+h2+h3+h4|w0+w1+w2_h0+h1+h2+h3+h4" +
				"|w0+w1+w2+w3_h0+h1+h2+h3+h4|w0+w1+w2+w3+w4_h0+h1+h2+h3+h4" +
				"|w0+w1+w2+w3+w4+w5_h0+h1+h2+h3+h4|w0+w1+w2+w3+w4+w5+w6_h0+h1+h2+h3+h4" +
				"|w0+w1+w2+w3+w4+w5+w6+w7_h0+h1+h2+h3+h4" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0+h1+h2+h3+h4" +
				"|0_h0+h1+h2+h3+h4+h5|w0_h0+h1+h2+h3+h4+h5|w0+w1_h0+h1+h2+h3+h4+h5" +
				"|w0+w1+w2_h0+h1+h2+h3+h4+h5|w0+w1+w2+w3_h0+h1+h2+h3+h4+h5" +
				"|w0+w1+w2+w3+w4_h0+h1+h2+h3+h4+h5" +
				"|w0+w1+w2+w3+w4+w5_h0+h1+h2+h3+h4+h5" +
				"|w0+w1+w2+w3+w4+w5+w6_h0+h1+h2+h3+h4+h5" +
				"|w0+w1+w2+w3+w4+w5+w6+w7_h0+h1+h2+h3+h4+h5" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0+h1+h2+h3+h4+h5" +
				"|0_h0+h1+h2+h3+h4+h5+h6" +
				"|w0_h0+h1+h2+h3+h4+h5+h6" +
				"|w0+w1_h0+h1+h2+h3+h4+h5+h6" +
				"|w0+w1+w2_h0+h1+h2+h3+h4+h5+h6" +
				"|w0+w1+w2+w3_h0+h1+h2+h3+h4+h5+h6" +
				"|w0+w1+w2+w3+w4_h0+h1+h2+h3+h4+h5+h6" +
				"|w0+w1+w2+w3+w4+w5_h0+h1+h2+h3+h4+h5+h6" +
				"|w0+w1+w2+w3+w4+w5+w6_h0+h1+h2+h3+h4+h5+h6" +
				"|w0+w1+w2+w3+w4+w5+w6+w7_h0+h1+h2+h3+h4+h5+h6" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0+h1+h2+h3+h4+h5+h6" +
				"|0_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0+w1_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0+w1+w2_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0+w1+w2+w3_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0+w1+w2+w3+w4_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0+w1+w2+w3+w4+w5_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0+w1+w2+w3+w4+w5+w6_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0+w1+w2+w3+w4+w5+w6+w7_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0+h1+h2+h3+h4+h5+h6+h7" +
				"|0_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0+w1_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0+w1+w2_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0+w1+w2+w3_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0+w1+w2+w3+w4_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0+w1+w2+w3+w4+w5_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0+w1+w2+w3+w4+w5+w6_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0+w1+w2+w3+w4+w5+w6+w7_h0+h1+h2+h3+h4+h5+h6+h7+h8" +
				"|w0+w1+w2+w3+w4+w5+w6+w7+w8_h0+h1+h2+h3+h4+h5+h6+h7+h8", nil},
	}

	for i, tt := range tests {
		func() {
			layout, err := xstack.Layout(uint64(tt.n))

			if !errors.Is(err, tt.werr) {
				t.Errorf("#%d: err = %v, want %v", i, err, tt.werr)
			}

			if layout != tt.wlayout {
				t.Errorf("#%d: layout = %s, want %s", i, layout, tt.wlayout)
			}
		}()
	}
}
