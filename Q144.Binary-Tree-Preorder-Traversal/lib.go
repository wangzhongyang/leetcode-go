package Q144_Binary_Tree_Preorder_Traversal

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	//return recursiveMain(root)
	return iteration(root)
}

type Stack []*TreeNode

func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Push(n *TreeNode) {
	*s = append(*s, n)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//迭代
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
//Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
func iteration(root *TreeNode) []int {
	s, res := new(Stack), make([]int, 0)
	for root != nil || !s.IsEmpty() {
		for root != nil {
			res = append(res, root.Val)
			s.Push(root)
			root = root.Left
		}
		root = s.Pop().Right
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
