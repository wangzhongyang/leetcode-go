package Q111_Minimum_Depth_of_Binary_Tree

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	//return recursive(root)
	return iteration(root)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type StackItem struct {
	Node  *TreeNode
	Depth int
}

type Stack []StackItem

func (s *Stack) Pop() StackItem {
	if s.IsEmpty() {
		return StackItem{}
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Push(n *TreeNode, depth int) {
	*s = append(*s, StackItem{
		Node:  n,
		Depth: depth,
	})
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// 迭代
//Runtime: 4 ms, faster than 97.23% of Go online submissions for Minimum Depth of Binary Tree.
//Memory Usage: 5.3 MB, less than 100.00% of Go online submissions for Minimum Depth of Binary Tree.
func iteration(root *TreeNode) int {
	min, s := math.MaxInt64, new(Stack)
	if root != nil {
		s.Push(root, 1)
	} else {
		return 0
	}
	for !s.IsEmpty() {
		temp := s.Pop()
		if temp.Node != nil {
			if temp.Node.Left == nil && temp.Node.Right == nil {
				min = Min(min, temp.Depth)
			}
			s.Push(temp.Node.Left, temp.Depth+1)
			s.Push(temp.Node.Right, temp.Depth+1)
		}
	}
	return min
}

//递归
//Runtime: 4 ms, faster than 97.13% of Go online submissions for Minimum Depth of Binary Tree.
//Memory Usage: 5.3 MB, less than 100.00% of Go online submissions for Minimum Depth of Binary Tree.
func recursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil {
		return recursive(root.Left) + 1
	}
	if root.Left == nil {
		return recursive(root.Right) + 1
	}
	return Min(recursive(root.Left), recursive(root.Right)) + 1
}
