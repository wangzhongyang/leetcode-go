package Q25_reverse_nodes_in_k_group

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  *ListNode
	output *ListNode
	n      int
}{
	{
		n: 2,
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
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
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
	},
	{
		n: 3,
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
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
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
	},
}

func Test_mergeTwoLists(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := mergeTwoLists(item.input1, item.input2)
		PrintList(output)
		a.Equal(item.output, output)
	}
}
