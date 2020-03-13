package Q16_3Sum_Closest

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/3sum-closest/
// 中间指针移动不能跳跃，会造成漏数据
//Runtime: 4 ms, faster than 97.71% of Go online submissions for 3Sum Closest.
//Memory Usage: 2.7 MB, less than 50.00% of Go online submissions for 3Sum Closest.
func threeSumClosest(nums []int, target int) int {
	min := 9223372036854775807
	res := 0
	sort.Ints(nums)
	if len(nums) < 3 {
		return 0
	}
	for k1, num1 := range nums {
		// 去重
		if k1 > 0 && nums[k1-1] == num1 {
			continue
		}
		k2, k3 := k1+1, len(nums)-1

		for k2 < k3 {
			temp := target - num1 - nums[k2] - nums[k3]
			tempAbs := int(math.Abs(float64(temp)))
			if tempAbs < min {
				min = tempAbs
				res = num1 + nums[k2] + nums[k3]
			}
			if temp < 0 {
				k3--
			} else {
				k2++
			}
		}
	}
	return res
}
