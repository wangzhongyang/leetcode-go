package Q64_Minimum_Path_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  [][]int
	output int
}{
	{
		input: [][]int{
			[]int{1, 3, 1},
			[]int{1, 5, 1},
			[]int{4, 2, 1},
		},
		output: 7,
	},
}

func TestMinPathSum(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := minPathSum(v.input)
		a.Equal(v.output, res)
	}
}
