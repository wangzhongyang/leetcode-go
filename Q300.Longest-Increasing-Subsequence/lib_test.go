package Q300_Longest_Increasing_Subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output int
}{
	{
		input:  []int{10, 9, 2, 5, 3, 7, 101, 18},
		output: 4,
	},
	{
		input:  []int{0, 8, 4, 12, 2},
		output: 3,
	},
	{
		input:  []int{4, 10, 4, 3, 8, 9},
		output: 3,
	},
}

func TestLengthOfLIS(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := lengthOfLIS(v.input)
		a.Equal(v.output, res)
	}
}
