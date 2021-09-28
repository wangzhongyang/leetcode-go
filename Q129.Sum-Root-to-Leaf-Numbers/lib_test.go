package Q129_Sum_Root_to_Leaf_Numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *TreeNode
	output int
}{
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		output: 25,
	},
	{
		input: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
		output: 1026,
	},
}

func Test_SumNumbers(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := sumNumbers(item.input)
		a.Equal(output, item.output)
	}
}
