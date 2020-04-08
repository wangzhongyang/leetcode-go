package Q98_Validate_Binary_Search_Tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *TreeNode
	output bool
}{
	{
		input: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		output: true,
	},
	{
		input: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
		},
		output: false,
	},
	{
		input: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  nil,
					Right: nil,
				},
			},
		},
		output: false,
	},
}

func TestIsValidBST(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		fmt.Println("root:\t", v.input.Val)
		res := isValidBST(v.input)
		fmt.Println("res:\t", res)
		a.Equal(v.output, res)
	}
}
