package Q912_Sort_an_Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{5, 2, 3, 1},
		output: []int{1, 2, 3, 5},
	},
	{
		input:  []int{5, 1, 1, 2, 0, 0},
		output: []int{0, 0, 1, 1, 2, 5},
	},
	{
		input:  []int{1, 3, 2, 5, 7, 8},
		output: []int{1, 2, 3, 5, 7, 8},
	},
}

func Test_QuickSort(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := quickSort(item.input)
		a.Equal(item.output, output)
	}
}

func Test_HeapSort(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := heapSort(item.input)
		a.Equal(item.output, output)
	}
}
