package Q40_Combination_Sum_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	candidates []int
	target     int
	output     [][]int
}{
	{
		candidates: []int{10, 1, 2, 7, 6, 1, 5},
		target:     8,
		output: [][]int{
			[]int{1, 7},
			[]int{1, 2, 5},
			[]int{2, 6},
			[]int{1, 1, 6},
		},
	},
	{
		candidates: []int{2, 5, 2, 1, 2},
		target:     5,
		output: [][]int{
			[]int{1, 2, 2},
			[]int{5},
		},
	},
}

func Test_threeSum(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := combinationSum2(item.candidates, item.target)
		a.Equal(item.output, output)
	}
}
