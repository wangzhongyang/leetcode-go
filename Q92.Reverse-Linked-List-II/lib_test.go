package Q92_Reverse_Linked_List_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input       *ListNode
	left, right int
	output      *ListNode
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
							Val: 5,
						},
					},
				},
			},
		},
		left:  2,
		right: 4,
		output: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
	},

	//{
	//	input: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val: 5,
	//		},
	//	},
	//	left:  1,
	//	right: 2,
	//	output: &ListNode{
	//		Val: 5,
	//		Next: &ListNode{
	//			Val: 3,
	//		},
	//	},
	//},
}

func Test_reverseBetween(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := reverseBetween(item.input, item.left, item.right)
		a.Equal(item.output, output)
	}
}
