package Q114_Flatten_Binary_Tree_to_Linked_List

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//flatten 使用的常规方法：前序遍历得到节点列表，然后组装。LeetCode中的方法3确实快
func flatten(root *TreeNode) {
	list := dp(root)
	if len(list) == 1 {
		return
	}
	for i := 1; i < len(list); i++ {
		list[i-1].Left = nil
		list[i-1].Right = list[i]
	}
}

func dp(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	if root == nil {
		return nil
	}
	res = append(res, root)
	res = append(res, dp(root.Left)...)
	res = append(res, dp(root.Right)...)
	return res
}
