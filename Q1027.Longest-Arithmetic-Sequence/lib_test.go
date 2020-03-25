package Q1027_Longest_Arithmetic_Sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output int
}{
	{
		input:  []int{3, 6, 9, 12},
		output: 4,
	},
	{
		input:  []int{3, 4, 7, 2, 10},
		output: 3,
	},
	{
		input:  []int{20, 1, 15, 3, 10, 5, 8},
		output: 4,
	},
}

func TestLongestArithSeqLength(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := longestArithSeqLength(v.input)
		a.Equal(v.output, res)
	}
}
