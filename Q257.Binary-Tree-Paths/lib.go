package Q257_Binary_Tree_Paths

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	return iteration(root)
	//return recursive(root)
}

// 迭代
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Paths.
//Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Binary Tree Paths.
func iteration(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res, stackNode, stackStr := make([]string, 0), make([]*TreeNode, 0), make([]string, 0)
	stackNode, stackStr = append(stackNode, root), append(stackStr, fmt.Sprint(root.Val))
	for len(stackNode) != 0 {
		node, nodeStr := stackNode[len(stackNode)-1], stackStr[len(stackStr)-1]
		stackNode, stackStr = stackNode[:len(stackNode)-1], stackStr[:len(stackStr)-1]
		if node.Left == nil && node.Right == nil {
			res = append(res, nodeStr)
		}
		if node.Right != nil {
			stackNode = append(stackNode, node.Right)
			stackStr = append(stackStr, fmt.Sprintf("%s->%d", nodeStr, node.Right.Val))
		}
		if node.Left != nil {
			stackNode = append(stackNode, node.Left)
			stackStr = append(stackStr, fmt.Sprintf("%s->%d", nodeStr, node.Left.Val))
		}
	}
	return res
}

//递归
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Paths.
//Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Binary Tree Paths.
func recursive(root *TreeNode) []string {
	list := make([]string, 0)
	if root == nil {
		return list
	}
	if root.Right == nil && root.Left == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}
	listLeft := recursive(root.Left)
	for k, v := range listLeft {
		listLeft[k] = fmt.Sprintf("%d->%s", root.Val, v)
	}
	listRight := recursive(root.Right)
	for k, v := range listRight {
		listRight[k] = fmt.Sprintf("%d->%s", root.Val, v)
	}
	return append(listLeft, listRight...)
}
