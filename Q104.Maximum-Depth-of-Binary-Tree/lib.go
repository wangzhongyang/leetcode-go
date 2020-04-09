package Q104_Maximum_Depth_of_Binary_Tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return iteration(root)
	//return recursiveMain(root)
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
//Runtime: 4 ms, faster than 90.52% of Go online submissions for Maximum Depth of Binary Tree.
//Memory Usage: 4.3 MB, less than 100.00% of Go online submissions for Maximum Depth of Binary Tree.
func iteration(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s, max := new(Stack), 0
	s.Push(root, 1)
	for !s.IsEmpty() {
		item := s.Pop()
		if item.Node != nil {
			max = Max(item.Depth, max)
			s.Push(item.Node.Left, item.Depth+1)
			s.Push(item.Node.Right, item.Depth+1)
		}
	}
	return max
}

//递归
//Runtime: 4 ms, faster than 90.52% of Go online submissions for Maximum Depth of Binary Tree.
//Memory Usage: 4.4 MB, less than 83.33% of Go online submissions for Maximum Depth of Binary Tree.
func recursiveMain(root *TreeNode) int {
	max := 0
	recursive(root, &max, 0)
	return max
}

func recursive(root *TreeNode, max *int, n int) {
	if root == nil {
		return
	}
	n++
	*max = Max(*max, n)
	recursive(root.Left, max, n)
	recursive(root.Right, max, n)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
