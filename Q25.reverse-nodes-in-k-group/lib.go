package Q25_reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}

// 翻转
func flip(head *ListNode) (begin, end *ListNode) {
	begin, end = head, head
	for head.Next != nil {
		for head != nil {
			next := head.Next
			head.Next = begin
			begin = head
			head = next
		}
	}
	return begin, end
}
