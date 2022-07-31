package Q106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//buildTree
//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：4 MB, 在所有 Go 提交中击败了77.17%的用户
func buildTree(inorder []int, postorder []int) *TreeNode {
	return dp(inorder, postorder)
}

func dp(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	var index int
	for k, v := range inorder {
		if v == root.Val {
			index = k
			break
		}
	}
	// 左子树
	if index > 0 {
		root.Left = dp(inorder[:index], postorder[:index])
	}
	if index < len(inorder)-1 {
		root.Right = dp(inorder[index+1:], postorder[index:len(postorder)-1])
	}
	return root
}
