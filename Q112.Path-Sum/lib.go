package Q112_Path_Sum

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum = 0

func hasPathSum(root *TreeNode, targetSum int) bool {
	return iterationHasPathSum(root, targetSum)
	//return recursiveHasPathSum(root, targetSum)
}

//iterationHasPathSum 迭代写法
//Runtime: 4 ms, faster than 89.72% of Go online submissions for Path Sum.
//Memory Usage: 6.7 MB, less than 7.71% of Go online submissions for Path Sum.
func iterationHasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 两个队列，用于储存节点和和路径
	q := []*TreeNode{root}
	qSum := []int{root.Val}
	for len(q) != 0 {
		// 出队
		n := q[0]
		p := qSum[0]
		q = q[1:]
		qSum = qSum[1:]
		// 判断:当遍历到叶子节点时,如果p等于给定的值就返回真
		if n.Left == nil && n.Right == nil {
			if p == targetSum {
				return true
			}
			continue
		}
		// 入队
		if n.Left != nil {
			q = append(q, n.Left)
			qSum = append(qSum, n.Left.Val+p)
		}
		if n.Right != nil {
			q = append(q, n.Right)
			qSum = append(qSum, n.Right.Val+p)
		}
	}
	return false
}

//recursiveHasPathSum 递归写法
//Runtime: 4 ms, faster than 89.72% of Go online submissions for Path Sum.
//Memory Usage: 4.7 MB, less than 38.97% of Go online submissions for Path Sum
func recursiveHasPathSum(root *TreeNode, targetSum int) bool {
	sum = targetSum
	return recursive(root, 0)
}

func recursive(root *TreeNode, pathSum int) bool {
	if root == nil {
		return false
	}
	pathSum += root.Val
	if pathSum > sum {
		return false
	}
	if root.Right == nil && root.Left == nil && pathSum == sum {
		return true
	}
	if recursive(root.Left, pathSum) || recursive(root.Right, pathSum) {
		return true
	}
	return false
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
