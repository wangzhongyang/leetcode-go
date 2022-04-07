package Q23_Merge_k_Sorted_Lists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	Input  []*ListNode
	Output *ListNode
}{
	{
		Input: []*ListNode{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  6,
					Next: nil,
				},
			},
		},
		Output: &ListNode{
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
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  6,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func Test_mergeKLists(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := mergeKLists(item.Input)
		a.Equal(item.Output, output)
		Println(output)
	}
}

func Println(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, ",")
		head = head.Next
	}
}
