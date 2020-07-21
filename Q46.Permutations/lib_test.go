package Q46_Permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output [][]int
}{
	{
		input: []int{1, 2, 3},
		output: [][]int{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			[]int{2, 1, 3},
			[]int{2, 3, 1},
			[]int{3, 1, 2},
			[]int{3, 2, 1},
		},
	},
}

func Test_threeSum(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := permute(item.input)
		a.Equal(item.output, output)
	}
}
