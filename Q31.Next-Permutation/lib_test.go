package Q31_Next_Permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{1, 2, 3},
		output: []int{1, 3, 2},
	},
	{
		input:  []int{3, 2, 1},
		output: []int{1, 2, 3},
	},
	{
		input:  []int{1, 1, 5},
		output: []int{1, 5, 1},
	},
	{
		input:  []int{1},
		output: []int{1},
	},
}

func Test_nextPermutation(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		nextPermutation(item.input)
		a.Equal(item.output, item.input)
	}
}
