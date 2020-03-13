package Q15_3Sum

import (
	"sort"
)

// https://leetcode-cn.com/problems/3sum/
// 排序，以一个数为标准对后面的数使用双指针
//Runtime: 28 ms, faster than 99.72% of Go online submissions for 3Sum.
//Memory Usage: 6.9 MB, less than 100.00% of Go online submissions for 3Sum.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	for k1, num1 := range nums {
		// 去重
		if k1 > 0 && nums[k1-1] == num1 {
			continue
		}
		k2, k3 := k1+1, len(nums)-1
		for k2 < k3 {
			if num1+nums[k2]+nums[k3] == 0 {
				res = append(res, []int{num1, nums[k2], nums[k3]})
				// 去重
				{
					for k2 < k3 && nums[k2] == nums[k2+1] {
						k2++
					}
					for k2 < k3 && nums[k3] == nums[k3-1] {
						k3--
					}
					k2++
					k3--
				}
			} else if num1+nums[k2]+nums[k3] > 0 {
				k3--
			} else {
				k2++
			}
		}
	}
	return res
}
