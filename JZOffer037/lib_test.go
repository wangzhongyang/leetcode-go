package JZOffer037

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums   []int
	output []int
}{
	{
		nums:   []int{5, 10, -5},
		output: []int{5, 10},
	},
	{
		nums:   []int{8, -8},
		output: []int{},
	},
	{
		nums:   []int{10, 2, -5},
		output: []int{10},
	},
	{
		nums:   []int{-2, -1, 1, 2},
		output: []int{-2, -1, 1, 2},
	},
	{
		nums:   []int{-2, -2, 1, -2},
		output: []int{-2, -2, -2},
	},
}

func Test_asteroidCollision(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := asteroidCollision(item.nums)
		a.Equal(item.output, output)
	}
}
