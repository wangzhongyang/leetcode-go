package Q145_Binary_Tree_Postorder_Traversal

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
		output: []int{4, 7, 8, 5, 2, 6, 3, 1},
	},
}

func Test_PostorderTraversal(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := postorderTraversal(item.input)
		a.Equal(output, item.output)
	}
}
