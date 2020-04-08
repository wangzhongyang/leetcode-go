package Q144_Binary_Tree_Preorder_Traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *TreeNode
	output []int
}{
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		output: []int{1, 2, 4, 5, 7, 8, 3, 6},
	},
}

func Test_preorderTraversal(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := preorderTraversal(item.input)
		a.Equal(output, item.output)
	}
}
