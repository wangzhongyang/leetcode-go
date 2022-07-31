package Q105_Construct_Binary_Tree_From_Preorder_and_Inorder_Traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	preorder []int
	inorder  []int
	output   *TreeNode
}{
	{
		preorder: []int{3, 9, 20, 15, 7},
		inorder:  []int{9, 3, 15, 20, 7},
		output: &TreeNode{
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
	},
	{
		preorder: []int{-1},
		inorder:  []int{-1},
		output:   &TreeNode{Val: -1},
	},
}

func Test_buildTree(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := buildTree(v.preorder, v.inorder)
		a.Equal(v.output, res)
	}
}
