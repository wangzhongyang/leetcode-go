package Q110_Balanced_Binary_Tree

import "math"

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//isBalanced 自底向上递归求解。如果一个节点的为平衡的，返回高度；否则返回-1
//执行用时：4 ms, 在所有 Go 提交中击败了95.29%的用户
//内存消耗：5.5 MB, 在所有 Go 提交中击败了65.97%的用户
func isBalanced(root *TreeNode) bool {
	if height(root) < 0 {
		return false
	}
	return true
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h1 := height(root.Left)
	h2 := height(root.Right)
	if h1 < 0 || h2 < 0 || math.Abs(float64(h1)-float64(h2)) > 1 {
		return -1
	}
	if h1 < h2 {
		h1 = h2
	}
	return h1 + 1
}
