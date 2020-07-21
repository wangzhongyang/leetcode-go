package Q39_Combination_Sum

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
		candidates: []int{2, 3, 6, 7},
		target:     7,
		output: [][]int{
			[]int{2, 2, 3},
			[]int{7},
		},
	},
	{
		candidates: []int{2, 3, 5},
		target:     8,
		output: [][]int{
			[]int{2, 2, 2, 2},
			[]int{2, 3, 3},
			[]int{3, 5},
		},
	},
}

func Test_threeSum(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := combinationSum(item.candidates, item.target)
		a.Equal(item.output, output)
	}
}
