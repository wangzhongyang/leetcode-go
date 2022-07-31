package Q101_Symmetric_Tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *TreeNode
	output bool
}{
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		},
		output: true,
	},
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		},
		output: false,
	},
}

func Test_isSymmetric(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := isSymmetric(item.input)
		a.Equal(output, item.output)
	}
}
