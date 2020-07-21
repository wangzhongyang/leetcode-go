package Q47_Permutations_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output [][]int
}{
	{
		input: []int{1, 1, 2},
		output: [][]int{
			[]int{1, 1, 2},
			[]int{1, 2, 1},
			[]int{2, 1, 1},
		},
	},
}

func Test_threeSum(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := permuteUnique(item.input)
		a.Equal(item.output, output)
	}
}
