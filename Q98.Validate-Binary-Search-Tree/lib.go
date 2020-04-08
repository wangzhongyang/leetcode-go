package Q98_Validate_Binary_Search_Tree

import (
	"math"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Runtime: 12 ms, faster than 5.43% of Go online submissions for Validate Binary Search Tree.
//Memory Usage: 5.6 MB, less than 33.33% of Go online submissions for Validate Binary Search Tree.
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MaxInt64, math.MinInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	val := root.Val
	if lower != math.MaxInt64 && val >= lower {
		return false
	}
	if upper != math.MinInt64 && val <= upper {
		return false
	}
	if !helper(root.Left, val, upper) {
		return false
	}
	if !helper(root.Right, lower, val) {
		return false
	}
	return true
}

type Stack []*TreeNode

func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Push(n *TreeNode) {
	*s = append(*s, n)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
