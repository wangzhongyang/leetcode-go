package Q144_Binary_Tree_Preorder_Traversal

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func preorderTraversal(root *TreeNode) []int {
	//return recursiveMain(root)
	return iteration(root)
}

//iteration 迭代
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
//Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
func iteration(root *TreeNode) []int {
	res := make([]int, 0)
	list, index := make([]*TreeNode, 100), -1
	for root != nil || index != -1 {
		if root != nil {
			res = append(res, root.Val)
			index++
			list[index] = root
			root = root.Left
			continue
		}
		root = list[index].Right
		list[index] = nil
		index--
	}
	return res
}

// 递归
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
//Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
func recursiveMain(root *TreeNode) []int {
	res := make([]int, 0)
	recursive(root, &res)
	return res
}

func recursive(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	recursive(root.Left, res)
	recursive(root.Right, res)
}
