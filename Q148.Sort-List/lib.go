package Q148_Sort_List

type ListNode struct {
	Val  int
	Next *ListNode
}

//sortList
// 通过快慢指针进行二分，然后合并。即：归并算法
//Runtime: 28 ms, faster than 92.55% of Go online submissions for Sort List.
//Memory Usage: 7.3 MB, less than 79.26% of Go online submissions for Sort List.
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	var preSlow *ListNode
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	preSlow.Next = nil
	l := sortList(head)
	r := sortList(slow)
	return mergeList(l, r)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return dummy.Next
}
