package Q94_Binary_Tree_Inorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	//res := make([]int, 0)
	//Recursive(root, &res)
	//return res
	return IterationMain(root)
}

// Recursive 递归
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
//Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
func Recursive(tree *TreeNode, nodes *[]int) {
	if tree != nil {
		if tree.Left != nil {
			Recursive(tree.Left, nodes)
		}
		*nodes = append(*nodes, tree.Val)
		if tree.Right != nil {
			Recursive(tree.Right, nodes)
		}
	}
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

// 迭代
func IterationMain(root *TreeNode) []int {
	s, res := new(Stack), make([]int, 0)
	for root != nil || !s.IsEmpty() {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
		if temp := s.Pop(); temp != nil {
			res = append(res, temp.Val)
			root = temp.Right
		}

	}
	return res
}
