package Q82_Remove_Duplicates_from_Sorted_List_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//return newNode(head)
	return appendNode(head)
}

//appendNode
//Runtime: 4 ms, faster than 82.80% of Go online submissions for Remove Duplicates from Sorted List II.
//Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List II.
func appendNode(head *ListNode) *ListNode {
	root := &ListNode{Val: 0}
	p, now, before, shouldAppend := root, head.Next, head, true
	for now != nil {
		if now.Val != before.Val {
			if shouldAppend {
				p.Next = before
				p = p.Next
			}
			before, shouldAppend = now, true
		} else {
			shouldAppend = false
		}
		now = now.Next
	}
	if shouldAppend {
		p.Next = before
	} else {
		p.Next = nil
	}
	return root.Next
}

// newNode
// 思路：通过新增节点进行append节点
//Runtime: 4 ms, faster than 82.80% of Go online submissions for Remove Duplicates from Sorted List II.
//Memory Usage: 3.1 MB, less than 43.55% of Go online submissions for Remove Duplicates from Sorted List II.
func newNode(head *ListNode) *ListNode {
	root := &ListNode{Val: 0}
	p, now, beforeVal, shouldAppend := root, head.Next, head.Val, true
	for now != nil {
		if now.Val != beforeVal {
			if shouldAppend {
				p.Next = &ListNode{
					Val: beforeVal,
				}
				p = p.Next
			}

			beforeVal, shouldAppend = now.Val, true

		} else {
			shouldAppend = false
		}
		now = now.Next
	}
	if shouldAppend {
		p.Next = &ListNode{
			Val: beforeVal,
		}
	}
	return root.Next
}
