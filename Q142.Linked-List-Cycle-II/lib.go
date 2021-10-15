package Q142_Linked_List_Cycle_II

type ListNode struct {
	Val  int
	Next *ListNode
}

//detectCycle
//快慢指针，获取相遇的节点；
//然后从头结点及相遇的节点同时向下读取相交的点为第一个入环的节点
//Runtime: 4 ms, faster than 96.05% of Go online submissions for Linked List Cycle II.
//Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for Linked List Cycle II.
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast, item := head, head, new(ListNode)
	for {
		if fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			return nil
		}
		if fast == slow {
			item = fast
			break
		}
	}
	tmp := head
	for tmp != item {
		tmp, item = tmp.Next, item.Next
	}
	return tmp
}
