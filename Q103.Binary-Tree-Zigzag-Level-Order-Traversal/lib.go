package Q103_Binary_Tree_Zigzag_Level_Order_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
//Memory Usage: 2.5 MB, less than 81.17% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	list, isLeft := []*TreeNode{root}, true
	for len(list) != 0 {
		listTmp, resItem := make([]*TreeNode, 0), make([]int, len(list))
		for i := 0; i < len(list); i++ {
			index := i
			if !isLeft {
				index = len(list) - i - 1
			}
			resItem[index] = list[i].Val
			if list[i].Left != nil {
				listTmp = append(listTmp, list[i].Left)
			}
			if list[i].Right != nil {
				listTmp = append(listTmp, list[i].Right)
			}
		}

		if isLeft {
			isLeft = false
		} else {
			isLeft = true
		}
		res, list = append(res, resItem), listTmp
	}
	return res
}
