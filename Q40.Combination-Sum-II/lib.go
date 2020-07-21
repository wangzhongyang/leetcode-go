package Q40_Combination_Sum_II

import (
	"sort"
)

//https://leetcode-cn.com/problems/combination-sum-ii/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum II.
//Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Combination Sum II.
func combinationSum2(candidates []int, target int) [][]int {
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
		// 同级去重
		if k+begin-1 >= begin && candidates[k+begin-1] == v {
			continue
		}
		path = append(path, v)
		dfs(candidates, path, k+begin+1, residue, res)
		path = path[:len(path)-1]
	}
}
