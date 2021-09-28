package Q112_Path_Sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input     *TreeNode
	TargetSum int
	output    bool
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
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		TargetSum: 22,
		output:    true,
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
		output:    false,
	},
	{
		input:     nil,
		TargetSum: 0,
		output:    false,
	},
	{
		input:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		TargetSum: 1,
		output:    false,
	},
	{
		input:     &TreeNode{Val: 1},
		TargetSum: 1,
		output:    true,
	},
	{
		input: &TreeNode{
			Val:   -2,
			Right: &TreeNode{Val: -3},
		},
		TargetSum: -5,
		output:    true,
	},
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: -2,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -1,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val: -2,
				},
			},
		},
		TargetSum: -4,
		output:    true,
	},
}

func Test_HasPathSum(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		fmt.Println("----")
		output := hasPathSum(item.input, item.TargetSum)
		a.Equal(output, item.output)
	}
}
