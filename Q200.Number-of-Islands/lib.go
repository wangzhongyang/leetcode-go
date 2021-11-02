package Q200_Number_of_Islands

func dfs(i, j int, grid [][]byte) {

	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
		if grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i+1, j, grid)
			dfs(i-1, j, grid)
			dfs(i, j+1, grid)
			dfs(i, j-1, grid)
		}
	}
}

//Runtime: 4 ms, faster than 76.48% of Go online submissions for Number of Islands.
//Memory Usage: 3.8 MB, less than 73.16% of Go online submissions for Number of Islands.
func numIslands(grid [][]byte) int {
	var res int

	for i, row := range grid {
		for j, cell := range row {
			if cell == '1' {
				res++
				dfs(i, j, grid)
			}
		}
	}

	return res
}
