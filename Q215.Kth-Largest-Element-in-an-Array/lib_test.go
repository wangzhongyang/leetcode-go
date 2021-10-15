package Q215_Kth_Largest_Element_in_an_Array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{10, 7, 5, 2, 1},
		output: []int{10, 7, 5, 2, 1},
	},
}

func Test_heapSort(t *testing.T) {
	fmt.Println(heapSort(tcs[0].input))
}

var tcs2 = []struct {
	arr    []int
	k      int
	output int
}{
	{
		arr:    []int{3, 2, 1, 5, 6, 4},
		k:      2,
		output: 5,
	},
	{
		arr:    []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		k:      4,
		output: 4,
	},
	{
		arr:    []int{2, 1},
		k:      2,
		output: 1,
	},
	{
		arr:    []int{-1, 2, 0},
		k:      1,
		output: 2,
	},
}

func Test_findKthLargest(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs2 {
		output := findKthLargest(item.arr, item.k)
		a.Equal(item.output, output)
	}
}
