package Tencent_interview

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcsQ1 = []struct {
	input  int
	output int
}{
	{
		input:  12,
		output: 11,
	},
	{
		input:  7,
		output: 6,
	},
	{
		input:  10,
		output: 9,
	},
	{
		input:  19,
		output: 17,
	},
}

func Test_Q1(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcsQ1 {
		output := Q1(item.input)
		fmt.Println(item.input, output)
		a.Equal(output, item.output)
	}
}

var tcsQ2 = []struct {
	nums   []int
	target int
	res    [][]int
}{
	{
		nums:   []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 0},
		target: 6,
		res:    [][]int{[]int{0, 5}, []int{1, 4}, []int{2, 3}, []int{6, 9}},
	},
}

func Test_Q2(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcsQ2 {
		output := Q2(item.nums, item.target)
		a.Equal(output, item.res)
	}
}

func Test_maxProfit(t *testing.T) {
	//a := assert.New(t)
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
