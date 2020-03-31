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
		input:  []int{9, 4, 7, 2, 10},
		output: 3,
	},
	{
		input:  []int{20, 1, 15, 3, 10, 5, 8},
		output: 4,
	},
	{
		input:  []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24},
		output: 2,
	},
	{
		input:  []int{24, 13, 1, 100, 0, 94, 3, 0, 3},
		output: 2,
	},
	{
		input:  []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24},
		output: 2,
	},
	{
		input: []int{
			44, 46, 22, 68, 45, 66, 43, 9, 37, 30, 50, 67, 32, 47, 44, 11, 15, 4, 11, 6, 20, 64, 54, 54, 61, 63, 23, 43,
			3, 12, 51, 61, 16, 57, 14, 12, 55, 17, 18, 25, 19, 28, 45, 56, 29, 39, 52, 8, 1, 21, 17, 21, 23, 70, 51, 61,
			21, 52, 25, 28},
		output: 6,
	},
}

func TestLongestArithSeqLength(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := longestArithSeqLength(v.input)
		a.Equal(v.output, res)
	}
	//fmt.Println(longestArithSeqLength([]int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}))
}
