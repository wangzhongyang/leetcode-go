package Q120_Triangle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  [][]int
	output int
}{
	{
		input:  [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
		output: 11,
	},
	{
		input:  [][]int{{2}, {3, 4}, {6, 5, 9}, {4, 4, 8, 0}},
		output: 14,
	},
	{
		input:  [][]int{{-1}, {9, -2}, {0, 4, 5}, {7, 4, -4, -5}, {9, 6, 0, 5, 7}, {9, 2, -9, -4, 5, -2}, {-4, -9, -5, -7, -5, -5, -2}, {-9, 5, -6, -4, 4, 1, 6, -4}, {-4, 3, 9, -2, 8, 6, -9, -2, -2}, {7, -6, 9, 8, -4, 2, -4, -2, -1, -2}, {0, 3, 2, 4, 0, -6, 7, 6, 7, -5, 2}, {9, 0, -8, 6, 4, 6, 2, 5, -9, 9, -1, -6}, {6, -3, -4, -5, 0, 3, 3, 4, -6, -4, -7, 7, 3}},
		output: -33,
	},
}

func Test_minimumTotal(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		Println(v.input)
		res := minimumTotal(v.input)
		Println(v.input)
		a.Equal(v.output, res)
	}
}

func Println(n [][]int) {
	for _, v := range n {
		for _, v2 := range v {
			fmt.Print(v2, "\t")
		}
		fmt.Println()
	}
}
