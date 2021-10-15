package Q56_Merge_Intervals

import (
	"sort"
)

//merge 合并区间
//核心思路：对给定数组排序后，如果当前端点的右节点大于下一个端点的左节点，则两个端点可以合并
//Runtime: 8 ms, faster than 92.77% of Go online submissions for Merge Intervals.
//Memory Usage: 4.6 MB, less than 95.54% of Go online submissions for Merge Intervals.
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}
	sort.Sort(InterVal(intervals))
	min, max := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > max {
			res = append(res, []int{min, max})
			min, max = intervals[i][0], intervals[i][1]
		} else {
			if intervals[i][1] > max {
				max = intervals[i][1]
			}
		}
	}
	res = append(res, []int{min, max})
	return res
}

type InterVal [][]int

func (x InterVal) Len() int           { return len(x) }
func (x InterVal) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x InterVal) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
