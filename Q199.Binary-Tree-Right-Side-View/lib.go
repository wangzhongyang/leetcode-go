package Q199_Binary_Tree_Right_Side_View

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//rightSideView 使用的是层序遍历方式，取最右边的节点，专业的词叫广度优先搜索
//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2.1 MB, 在所有 Go 提交中击败了93.34%的用户
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	list := []*TreeNode{root}
	for len(list) != 0 {
		res = append(res, list[len(list)-1].Val)
		tmpList := make([]*TreeNode, 0)
		for _, v := range list {
			if v.Left != nil {
				tmpList = append(tmpList, v.Left)
			}
			if v.Right != nil {
				tmpList = append(tmpList, v.Right)
			}
		}
		list = tmpList
	}
	return res
}
