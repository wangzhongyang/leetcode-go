package Q113_Path_Sum_II

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input     *TreeNode
	TargetSum int
	output    [][]int
}{
	{
		input: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 11,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 13,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		TargetSum: 22,
		output: [][]int{
			{5, 4, 11, 2},
			{5, 8, 4, 5},
		},
	},
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		TargetSum: 5,
		output:    [][]int{},
	},
	{
		input:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		TargetSum: 0,
		output:    [][]int{},
	},
}

func Test_pathSum(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		fmt.Println("----")
		output := pathSum(item.input, item.TargetSum)
		a.Equal(item.output, output)
	}
}
