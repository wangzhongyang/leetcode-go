package Q655_Print_Binary_Tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *TreeNode
	output [][]string
}{
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		output: [][]string{
			[]string{"", "1", ""},
			[]string{"2", "", ""},
		},
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
		output: [][]string{
			[]string{"", "", "", "1", "", "", ""},
			[]string{"", "2", "", "", "", "3", ""},
			[]string{"", "", "5", "", "", "", ""},
		},
	},
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		output: [][]string{
			[]string{"", "", "", "", "", "", "", "1", "", "", "", "", "", "", ""},
			[]string{"", "", "", "2", "", "", "", "", "", "", "", "5", "", "", ""},
			[]string{"", "3", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			[]string{"4", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
		},
	},
}

func Test_printTree(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := printTree(item.input)
		fmt.Println(item.input, output)
		a.Equal(output, item.output)
	}
}
