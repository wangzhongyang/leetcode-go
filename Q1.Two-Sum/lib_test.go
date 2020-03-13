package Q1_Two_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums   []int
	target int
	output []int
}{
	{
		nums:   []int{2, 7, 11, 15},
		target: 9,
		output: []int{0, 1},
	},
	{
		nums:   []int{5, 4, 6},
		target: 10,
		output: []int{1, 2},
	},
}

func Test_twoSum(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := twoSum(item.nums, item.target)
		a.Equal(item.output, output)
	}
}
