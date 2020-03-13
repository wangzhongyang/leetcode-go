package Q19_Remove_Nth_Node_From_End_of_List

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
//Memory Usage: 2.2 MB, less than 28.57% of Go online submissions for Remove Nth Node From End of List.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headI, headJ := head, head
	for i := 0; i < n+1; i++ {
		if headI == nil {
			return head.Next
		}
		headI = headI.Next
	}
	for headI != nil {
		headI = headI.Next
		headJ = headJ.Next
	}
	headJ.Next = headJ.Next.Next
	//PrintList(head)
	return head
}

func PrintList(head *ListNode) {
	var temp *ListNode
	temp = head
	for {
		fmt.Print(temp.Val)
		if temp.Next != nil {
			fmt.Print("->")
			temp = temp.Next
		} else {
			fmt.Println("")
			break
		}
	}
}
