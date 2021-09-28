package Q129_Sum_Root_to_Leaf_Numbers

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sumNumbers 根节点到叶节点数字之和
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum Root to Leaf Numbers.
//Memory Usage: 2.3 MB, less than 42.94% of Go online submissions for Sum Root to Leaf Numbers.
func sumNumbers(root *TreeNode) int {
	path, res := make([]int, 0), 0
	arr := recursive(root, path)
	for _, v := range arr {
		res += v
	}
	return res
}

func recursive(root *TreeNode, path []int) []int {
	res := make([]int, 0)
	if root == nil {
		return nil
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		res = []int{genInt(path)}
		path = path[:len(path)-1]
	}
	res = append(res, recursive(root.Left, path)...)
	res = append(res, recursive(root.Right, path)...)
	return res
}

func genInt(arr []int) int {
	res, l := 0, len(arr)
	for i := 0; i < l; i++ {
		res += int(math.Pow(10, float64(l-1-i))) * arr[i]
	}
	return res
}
