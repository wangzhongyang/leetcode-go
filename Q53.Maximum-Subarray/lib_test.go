package Q53_Maximum_Subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output int
}{
	{
		input:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		output: 6,
	},
	{
		input:  []int{1},
		output: 1,
	},
}

func Test_maxSubArray(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := maxSubArray(item.input)
		a.Equal(item.output, output)
	}
}
