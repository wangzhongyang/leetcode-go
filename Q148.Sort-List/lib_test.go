package Q148_Sort_List

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *ListNode
	output *ListNode
}{
	{
		input: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		output: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
	},
}

func Test_sortList(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := sortList(item.input)
		a.Equal(item.output, output)
	}
}
