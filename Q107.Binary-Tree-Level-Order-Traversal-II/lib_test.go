package Q107_Binary_Tree_Level_Order_Traversal_II

import (
	"fmt"
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
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		output: [][]int{
			{15, 7},
			{9, 20},
			{3},
		},
	},
}

func TestLevelOrderBottom(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := levelOrderBottom(v.input)
		fmt.Println(res)
		a.Equal(v.output, res)
	}
}
