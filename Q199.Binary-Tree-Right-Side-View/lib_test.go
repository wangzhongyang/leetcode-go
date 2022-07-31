package Q199_Binary_Tree_Right_Side_View

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
			Right: &TreeNode{
				Val: 3,
			},
		},
		output: []int{1, 3},
	},
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		output: []int{1, 3, 4},
	},
	{
		input:  nil,
		output: []int{},
	},
}

func Test_rightSideView(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := rightSideView(item.input)
		a.Equal(output, item.output)
	}
}
