package Q102_Binary_Tree_Level_Order

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
			{3},
			{9, 20},
			{15, 7},
		},
	},
}

func TestLevelOrder(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := levelOrder(v.input)
		fmt.Println(res)
		a.Equal(v.output, res)
	}
	//fmt.Println(longestArithSeqLength([]int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}))
}
