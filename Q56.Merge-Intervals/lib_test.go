package Q56_Merge_Intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  [][]int
	output [][]int
}{
	{
		input:  [][]int{{8, 10}, {1, 3}, {2, 6}, {15, 18}},
		output: [][]int{{1, 6}, {8, 10}, {15, 18}},
	},
}

func Test_Merge(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := merge(item.input)
		a.Equal(item.output, output)
	}
}
