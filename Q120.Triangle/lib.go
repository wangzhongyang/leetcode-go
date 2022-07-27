package Q120_Triangle

import "math"

//minimumTotal
//Runtime: 4 ms, faster than 86.62% of Go online submissions for Triangle.
//Memory Usage: 3.3 MB, less than 82.01% of Go online submissions for Triangle.
func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			min := math.MaxInt
			if j < len(triangle[i-1]) && triangle[i-1][j] < min {
				min = triangle[i-1][j]
			}
			if j-1 >= 0 && triangle[i-1][j-1] < min {
				min = triangle[i-1][j-1]
			}
			triangle[i][j] += min
		}
	}
	min := math.MaxInt
	for _, v := range triangle[len(triangle)-1] {
		if v < min {
			min = v
		}
	}
	return min
}
