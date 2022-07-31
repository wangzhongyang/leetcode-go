package Q106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	postorder []int
	inorder   []int
	output    *TreeNode
}{
	{
		postorder: []int{9, 15, 7, 20, 3},
		inorder:   []int{9, 3, 15, 20, 7},
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
		postorder: []int{-1},
		inorder:   []int{-1},
		output:    &TreeNode{Val: -1},
	},
}

func Test_buildTree(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := buildTree(v.inorder, v.postorder)
		a.Equal(v.output, res)
	}
}
