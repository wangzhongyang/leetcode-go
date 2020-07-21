package Q94_Binary_Tree_Inorder_Traversal

import (
	"fmt"
	"testing"
)

var tcs = []struct {
	input  *TreeNode
	output []int
}{
	//{
	//	input: &TreeNode{
	//		Val:  1,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val: 2,
	//			Left: &TreeNode{
	//				Val:   3,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//	},
	//	output: []int{1, 3, 2},
	//},
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
		output: []int{4, 2, 7, 5, 8, 1, 3, 6},
	},
}

func TestRecursive(t *testing.T) {

	a := []int{1, 2, 3, 4, 5, 6}
	b := a[:4]
	c := b[:2]
	a[1] = 0
	fmt.Println(a, b, c)
	return
	//a := assert.New(t)
	//for _, v := range tcs {
	//	res := inorderTraversal(v.input)
	//	a.Equal(v.output, res)
	//}
}
