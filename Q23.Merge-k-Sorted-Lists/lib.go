package Q23_Merge_k_Sorted_Lists

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := new(ListNode)
	last := head
	for !IsAllNil(lists) {
		minKey := WhoMin(lists)
		last.Next = lists[minKey]
		lists[minKey] = lists[minKey].Next
		last = last.Next
	}
	return head.Next
}

func IsAllNil(lists []*ListNode) bool {
	isAllNil := true
	for _, node := range lists {
		if node != nil {
			isAllNil = false
		}
	}
	return isAllNil
}

func WhoMin(lists []*ListNode) int {
	min, key := math.MaxInt, 0
	for k, node := range lists {
		if node != nil && node.Val < min {
			min, key = node.Val, k
		}
	}
	return key
}
