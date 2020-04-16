package Q257_Binary_Tree_Paths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *TreeNode
	output []string
}{
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		output: []string{"1->2"},
	},
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		output: []string{"1->2->5", "1->3"},
	},
}

func Test_binaryTreePaths(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := binaryTreePaths(item.input)
		fmt.Println(item.input, output)
		a.Equal(output, item.output)
	}
}
