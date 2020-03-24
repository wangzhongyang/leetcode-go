package Q206_reverse_linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/reverse-linked-list
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
//Memory Usage: 2.6 MB, less than 83.33% of Go online submissions for Reverse Linked List.
func reverseList(head *ListNode) *ListNode {
	var tempHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = tempHead
		tempHead = head
		head = next
	}
	return tempHead
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(fmt.Sprintf("%d->", head.Val))
		head = head.Next
	}
	fmt.Println("null")
}
