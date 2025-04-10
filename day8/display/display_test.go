package display_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day8/display"
	"github.com/stretchr/testify/assert"
)

//revive:disable:line-length-limit

func TestCalibrate(t *testing.T) {
	tests := []struct {
		inChunks []string
		code     []string
		expected int
	}{
		{
			inChunks: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			code:     []string{"cdfeb"},
			expected: 5,
		},
		{
			inChunks: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			code:     []string{"acedgfb"},
			expected: 8,
		},
		{
			inChunks: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			code:     []string{"gcdfa"},
			expected: 2,
		},
		{
			inChunks: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			code:     []string{"cdfeb", "acedgfb"},
			expected: 58,
		},
		{
			inChunks: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			code:     []string{"acedgfb", "gcdfa"},
			expected: 82,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			disp := display.New()
			disp.Calibrate(tt.inChunks)
			result := disp.Decode(tt.code)
			assert.Equal(t, tt.expected, result)
		})
	}
}

//revive:enable:line-length-limit
