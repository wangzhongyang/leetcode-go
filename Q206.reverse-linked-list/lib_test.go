package Q206_reverse_linked_list

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
		output: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	},
}

func Test_mergeTwoLists(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := reverseList(item.input)
		a.Equal(output, item.output)
	}
}
