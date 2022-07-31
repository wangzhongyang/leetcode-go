package Q101_Symmetric_Tree

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//isSymmetric
//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2.7 MB, 在所有 Go 提交中击败了99.84%的用户
func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if (node1 == nil && node2 != nil) || node2 == nil && node1 != nil {
		return false
	}
	return node1.Val == node2.Val && check(node1.Left, node2.Right) && check(node1.Right, node2.Left)
}
