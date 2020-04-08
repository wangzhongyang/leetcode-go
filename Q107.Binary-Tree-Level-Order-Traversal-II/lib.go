package Q107_Binary_Tree_Level_Order_Traversal_II

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	recursive(root, 0, &res)
	// 翻转
	j := len(res) - 1
	for i := 0; i < len(res); i++ {
		if i >= j {
			break
		}
		res[i], res[j] = res[j], res[i]
		j--
	}
	return res
}

// 递归的实现
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
//Memory Usage: 3 MB, less than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
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
