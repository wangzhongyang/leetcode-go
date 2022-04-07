package Q54_Spiral_Matrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	Input  [][]int
	Output []int
}{
	{
		Input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		Output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		Input: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12}},
		Output: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
	{
		Input:  [][]int{{3}, {2}},
		Output: []int{3, 2},
	},
	{
		Input:  [][]int{{6, 9, 7}},
		Output: []int{6, 9, 7},
	},
	{
		Input:  [][]int{{7}, {9}, {6}},
		Output: []int{7, 9, 6},
	},
}

func Test_spiralOrder(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := spiralOrder(item.Input)
		a.Equal(item.Output, output)
		fmt.Println(output)
	}
}
