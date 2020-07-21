package Q39_Combination_Sum

import (
	"sort"
)

//https://leetcode-cn.com/problems/combination-sum/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum.
//Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Combination Sum.
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	path, res := make([]int, 0), make([][]int, 0)
	dfs(candidates, path, 0, target, &res)
	return res
}

func dfs(candidates, path []int, begin, target int, res *[][]int) {
	if target == 0 {
		//slice 为引用类型，不copy后面的值会影响前面的结果
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for k, v := range candidates[begin:] {
		residue := target - v
		if residue < 0 {
			break
		}
		path = append(path, v)
		dfs(candidates, path, k+begin, residue, res)
		path = path[:len(path)-1]
	}
}
