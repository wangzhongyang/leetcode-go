package Q873_Length_of_Longest_Fibonacci_Subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output int
}{
	{
		input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		output: 5,
	},
	{
		input:  []int{1, 3, 7, 11, 12, 14, 18},
		output: 3,
	},
}

func TestLongestArithSeqLength(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := lenLongestFibSubseq(v.input)
		a.Equal(v.output, res)
	}
}
