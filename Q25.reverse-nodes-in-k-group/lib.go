package Q25_reverse_nodes_in_k_group

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//这方法思路清晰，也不会超时
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 遍历
	nodeArr, i := make([]*ListNode, 0), 0
	for head != nil {
		tmpHead := head
		head = head.Next
		tmpHead.Next = nil
		nodeArr = append(nodeArr, tmpHead)
		i++
	}
	// 交换
	for n := 1; n*k-1 < i; n++ {
		l, r := (n-1)*k, n*k-1
		for l < r {
			nodeArr[l], nodeArr[r] = nodeArr[r], nodeArr[l]
			l++
			r--
		}
	}
	// 组装
	res := new(ListNode)
	next := res
	for _, v := range nodeArr {
		next.Next = v
		next = next.Next
	}
	return res.Next
}

func flip(head *ListNode) (*ListNode, *ListNode) {
	res, last := new(ListNode), new(ListNode)
	for head != nil {
		next := res.Next
		res.Next = head
		res.Next.Next = next
		last = head
		head = head.Next
	}
	return res.Next, last
}

//https://leetcode.com/problems/reverse-nodes-in-k-group/
//Runtime: 4 ms, faster than 93.79% of Go online submissions for Reverse Nodes in k-Group.
//Memory Usage: 3.6 MB, less than 50.00% of Go online submissions for Reverse Nodes in k-Group.
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	res := &ListNode{
//		Val:  0,
//		Next: head,
//	}
//	begin, end := res, res
//	for end.Next != nil {
//		for i := 0; i < k && end != nil; i++ {
//			end = end.Next
//		}
//		if end == nil {
//			break
//		}
//		start, next := begin.Next, end.Next
//		end.Next = nil
//		begin.Next = flip(start)
//		start.Next = next
//		begin = start
//		end = begin
//	}
//	return res.Next
//}
//
//// 翻转
//func flip(head *ListNode) *ListNode {
//	var tempHead *ListNode
//	for head != nil {
//		next := head.Next
//		head.Next = tempHead
//		tempHead = head
//		head = next
//	}
//	return tempHead
//}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(fmt.Sprintf("%d->", head.Val))
		head = head.Next
	}
	fmt.Println("null")
}
