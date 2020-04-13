package Q111_Minimum_Depth_of_Binary_Tree

import (
	"fmt"
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
		},
		output: 2,
	},
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
		output: 2,
	},
}

func Test_maxDepth(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := minDepth(item.input)
		fmt.Println(item.input, output)
		a.Equal(output, item.output)
	}
}
