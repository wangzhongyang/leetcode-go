package Q103_Binary_Tree_Zigzag_Level_Order_Traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *TreeNode
	output [][]int
}{
	{
		input: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		output: [][]int{
			{3},
			{20, 9},
			{15, 7},
		},
	},
}

func Test_zigzagLevelOrder(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := zigzagLevelOrder(item.input)
		a.Equal(item.output, output)
	}
}
