package Q64_Minimum_Path_Sum

//https://leetcode.com/problems/minimum-path-sum
//Runtime: 8 ms, faster than 87.50% of Go online submissions for Minimum Path Sum.
//Memory Usage: 3.9 MB, less than 100.00% of Go online submissions for Minimum Path Sum.
func minPathSum(grid [][]int) int {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j == len(grid[i])-1 { // 最后一个点
				continue
			} else if i == len(grid)-1 { // 最后一行，只有往右
				grid[i][j] = grid[i][j] + grid[i][j+1]
			} else if j == len(grid[i])-1 { // 最后一列，只能往下
				grid[i][j] = grid[i][j] + grid[i+1][j]
			} else {
				grid[i][j] = min(grid[i][j]+grid[i][j+1], grid[i][j]+grid[i+1][j])
			}
		}
	}
	return grid[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
