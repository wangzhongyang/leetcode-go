package Q21_Merge_Two_Sorted_Lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input1 *ListNode
	input2 *ListNode
	output *ListNode
}{
	{
		input1: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		input2: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		output: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
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
