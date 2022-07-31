package Q105_Construct_Binary_Tree_From_Preorder_and_Inorder_Traversal

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//buildTree
//执行用时：4 ms, 在所有 Go 提交中击败了93.07%的用户
//内存消耗：3.7 MB, 在所有 Go 提交中击败了96.43%的用户
//前序遍历：{[根节点]，[左子树前序遍历]，[右子树谦虚遍历]}
//中序遍历：{[左子树中序遍历]，[跟节点]，[右子树中序遍历]}
func buildTree(preorder []int, inorder []int) *TreeNode {
	return dp(preorder, inorder)
}

func dp(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var index int
	for k, v := range inorder {
		if v == preorder[0] {
			index = k
			break
		}
	}
	// 左子树
	if index > 0 {
		root.Left = dp(preorder[1:1+index], inorder[:index])
	}
	if index < len(inorder) {
		root.Right = dp(preorder[1+index:], inorder[index+1:])
	}
	return root
}
