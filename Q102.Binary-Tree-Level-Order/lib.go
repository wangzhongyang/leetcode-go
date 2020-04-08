package Q102_Binary_Tree_Level_Order

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	recursive(root, 0, &res)
	return res
}

// 递归的实现
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
//Memory Usage: 3 MB, less than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
func recursive(root *TreeNode, n int, res *[][]int) {
	if root == nil {
		return
	} else {
		if len(*res)-1 < n {
			*res = append(*res, make([]int, 0))
		}
		(*res)[n] = append((*res)[n], root.Val)
	}
	n += 1
	recursive(root.Left, n, res)
	recursive(root.Right, n, res)
}
