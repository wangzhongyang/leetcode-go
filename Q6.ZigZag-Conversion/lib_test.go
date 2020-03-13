package Q6_ZigZag_Conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input   string
	numRows int
	output  string
}{
	{
		input:   "PAYPALISHIRING",
		numRows: 3,
		output:  "PAHNAPLSIIGYIR",
	},
	{
		input:   "LEETCODEISHIRING",
		numRows: 3,
		output:  "LCIRETOESIIGEDHN",
	},
	{
		input:   "PAYPALISHIRING",
		numRows: 4,
		output:  "PINALSIGYAHRPI",
	},
	{
		input:   "A",
		numRows: 1,
		output:  "A",
	},
}

func Test_convert(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := convert(item.input, item.numRows)
		a.Equal(item.output, output)
	}
}
