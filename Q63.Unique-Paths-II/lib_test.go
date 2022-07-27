package Q63_Unique_Paths_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  [][]int
	output int
}{
	{
		input:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		output: 2,
	},
	{
		input:  [][]int{{0, 1}, {0, 0}},
		output: 1,
	},
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := uniquePathsWithObstacles(v.input)
		a.Equal(v.output, res)
	}
}
