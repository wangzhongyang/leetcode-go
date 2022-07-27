package Q63_Unique_Paths_II

//uniquePathsWithObstacles
//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2.4 MB, 在所有 Go 提交中击败了77.23%的用户
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	res := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		res[i] = make([]int, len(obstacleGrid[i]))
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if i == 0 && j == 0 && obstacleGrid[0][0] != 1 {
				res[0][0] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				continue
			}
			top, left := 0, 0
			if i-1 >= 0 {
				top = res[i-1][j]
			}
			if j-1 >= 0 {
				left = res[i][j-1]
			}
			res[i][j] = top + left
		}
	}
	return res[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
