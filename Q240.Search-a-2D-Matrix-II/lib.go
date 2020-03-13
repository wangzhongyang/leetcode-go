package Q240_Search_a_2D_Matrix_II

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	h := len(matrix)
	if matrix[h-1][0] == target {
		return true
	}
	if matrix[h-1][0] > target {
		return searchMatrix(matrix[:h-1], target)
	}
	if matrix[h-1][0] < target {
		for k, v := range matrix {
			matrix[k] = v[1:]
		}
		return searchMatrix(matrix, target)
	}
	return false
}
