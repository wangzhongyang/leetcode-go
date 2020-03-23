package Q94_Binary_Tree_Inorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	Recursive(root, &res)
	return res
}

// Recursive 递归
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
//Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
func Recursive(tree *TreeNode, nodes *[]int) {
	if tree != nil {
		if tree.Left != nil {
			Recursive(tree.Left, nodes)
		}
		*nodes = append(*nodes, tree.Val)
		if tree.Right != nil {
			Recursive(tree.Right, nodes)
		}
	}
}
