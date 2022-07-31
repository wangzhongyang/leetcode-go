package Q113_Path_Sum_II

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	recursivePathSum(root, targetSum, []int{}, &res)
	return res
}

//recursivePathSum 递归，遍历到所有子节点，中间记录路径
//执行用时：4 ms, 在所有 Go 提交中击败了85.39%的用户
//内存消耗：7.5 MB, 在所有 Go 提交中击败了6.21%的用户
func recursivePathSum(root *TreeNode, targetSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	targetSum -= root.Val
	if targetSum < 0 {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		*res = append(*res, path)
	}
	tmpPath := make([]int, len(path))
	copy(tmpPath, path)
	recursivePathSum(root.Left, targetSum, path, res)
	recursivePathSum(root.Right, targetSum, tmpPath, res)
}
