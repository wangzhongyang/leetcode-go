package Q111_Minimum_Depth_of_Binary_Tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	//return recursive(root)
	return iteration(root)
}

type StackItem struct {
	Node  *TreeNode
	Depth int
}

type Stack []StackItem

func (s *Stack) Pop() StackItem {
	if s.IsEmpty() {
		return StackItem{}
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Push(n *TreeNode, depth int) {
	*s = append(*s, StackItem{
		Node:  n,
		Depth: depth,
	})
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// 迭代
func iteration(root *TreeNode) int {
	min, s := 0, new(Stack)
	for root != nil || !s.IsEmpty() {
		for root != nil {

			if s.IsEmpty() {

			}
		}
	}
	return min
}

//递归
//Runtime: 4 ms, faster than 97.13% of Go online submissions for Minimum Depth of Binary Tree.
//Memory Usage: 5.3 MB, less than 100.00% of Go online submissions for Minimum Depth of Binary Tree.
func recursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil {
		return recursive(root.Left) + 1
	}
	if root.Left == nil {
		return recursive(root.Right) + 1
	}
	return Min(recursive(root.Left), recursive(root.Right)) + 1
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
