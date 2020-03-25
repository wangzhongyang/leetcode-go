package Q25_reverse_nodes_in_k_group

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.com/problems/reverse-nodes-in-k-group/
//Runtime: 4 ms, faster than 93.79% of Go online submissions for Reverse Nodes in k-Group.
//Memory Usage: 3.6 MB, less than 50.00% of Go online submissions for Reverse Nodes in k-Group.
func reverseKGroup(head *ListNode, k int) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: head,
	}
	begin, end := res, res
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start, next := begin.Next, end.Next
		end.Next = nil
		begin.Next = flip(start)
		start.Next = next
		begin = start
		end = begin
	}
	return res.Next
}

// 翻转
func flip(head *ListNode) *ListNode {
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
