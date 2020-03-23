package Q94_Binary_Tree_Inorder_Traversal

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
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		output: []int{1, 3, 2},
	},
}

func TestRecursive(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := inorderTraversal(v.input)
		a.Equal(v.output, res)
	}
}
