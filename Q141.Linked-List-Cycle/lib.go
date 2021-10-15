package Q141_Linked_List_Cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

//hasCycle
//快慢指针或map
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}
