package Q21_Merge_Two_Sorted_Lists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/merge-two-sorted-lists/
// for 循环 在中国的LeetCode中，递归需要4ms，且需要的内存更多，for 循环为0ms
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
//Memory Usage: 2.5 MB, less than 71.43% of Go online submissions for Merge Two Sorted Lists.
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := res
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next, l1 = l1, l1.Next
		} else {
			temp.Next, l2 = l2, l2.Next
		}
		temp = temp.Next
	}
	if l1 == nil {
		temp.Next = l2
	}
	if l2 == nil {
		temp.Next = l1
	}
	return res.Next
}

// 递归
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
//Memory Usage: 2.6 MB, less than 7.14% of Go online submissions for Merge Two Sorted Lists.
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	var res *ListNode
//	if l1.Val >= l2.Val {
//		res = l2
//		res.Next = mergeTwoLists(l1, l2.Next)
//	} else {
//		res = l1
//		res.Next = mergeTwoLists(l1.Next, l2)
//	}
//	return res
//}

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
