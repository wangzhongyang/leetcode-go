package Q33_Search_in_Rotated_Sorted_Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	Nums   []int
	Target int
	Out    int
}{
	{
		Nums:   []int{4, 5, 6, 7, 0, 1, 2},
		Target: 0,
		Out:    4,
	},
	{
		Nums:   []int{4, 5, 6, 7, 0, 1, 2},
		Target: 3,
		Out:    -1,
	},
	{
		Nums:   []int{1},
		Target: 0,
		Out:    -1,
	},
	{
		Nums:   []int{1},
		Target: 1,
		Out:    0,
	},
	{
		Nums:   []int{7, 8, 1, 2, 3, 4, 5, 6},
		Target: 2,
		Out:    3,
	},
	{
		Nums:   []int{3, 5, 1},
		Target: 3,
		Out:    0,
	},
	{
		Nums:   []int{4, 5, 6, 7, 0, 1, 2},
		Target: 5,
		Out:    1,
	},
	{
		Nums:   []int{5, 1, 2, 3, 4},
		Target: 1,
		Out:    1,
	},
}

func Test_search(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := search(item.Nums, item.Target)
		a.Equal(item.Out, output)
	}
}
