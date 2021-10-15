package Q82_Remove_Duplicates_from_Sorted_List_II

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
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
		output: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	},
	{
		input: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
		output: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	},
	{
		input:  nil,
		output: nil,
	},
	{
		input: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
		output: &ListNode{
			Val: 1,
		},
	},
}

func Test_deleteDuplicates(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := deleteDuplicates(item.input)
		a.Equal(item.output, output)
	}
}
