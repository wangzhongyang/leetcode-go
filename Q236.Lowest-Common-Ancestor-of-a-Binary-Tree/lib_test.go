package Q236_Lowest_Common_Ancestor_of_a_Binary_Tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	p1 = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	q1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}
	o1 = &TreeNode{
		Val:   3,
		Left:  p1,
		Right: q1,
	}
)

var tcs = []struct {
	input  *TreeNode
	p      *TreeNode
	q      *TreeNode
	output *TreeNode
}{
	{
		input:  o1,
		p:      p1,
		q:      q1,
		output: o1,
	},
}

func Test_zigzagLevelOrder(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := lowestCommonAncestor(item.input, item.p, item.q)
		a.Equal(item.output, output)
	}
}
