package Q19_Remove_Nth_Node_From_End_of_List

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *ListNode
	n      int
	output *ListNode
}{
	{
		input: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		n: 2,
		output: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	},
	{
		input: &ListNode{
			Val:  1,
			Next: nil,
		},
		n:      1,
		output: nil,
	},
	{
		input: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
		n: 2,
		output: &ListNode{
			Val:  2,
			Next: nil,
		},
	},
}

func Test_removeNthFromEnd(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := removeNthFromEnd(item.input, item.n)
		a.Equal(item.output, output)
	}
}
