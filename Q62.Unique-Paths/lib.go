package Q62_Unique_Paths

//uniquePaths
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
//Memory Usage: 2.1 MB, less than 57.58% of Go online submissions for Unique Paths.
func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				res[i][j] = 1
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
	return res[m-1][n-1]
}
