package Q704_Binary_Search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	target int
	output int
}{
	{
		input:  []int{-1, 0, 3, 5, 9, 12},
		target: 9,
		output: 4,
	},
	{
		input:  []int{-1, 0, 3, 5, 9, 12},
		target: 2,
		output: -1,
	},
}

func Test_search(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := search(item.input, item.target)
		a.Equal(item.output, output)
	}
}
