package Q53_Maximum_Subarray

import "math"

func maxSubArray(nums []int) int {
	max, sum := -math.MaxInt, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		max = maxFunc(max, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

//dynamicProgramming 动态规划
//以下标i结束时的最大的和为f(i),则有f(i)=max(f(i-1)+nums[i],f(i))
//Runtime: 96 ms, faster than 35.40% of Go online submissions for Maximum Subarray.
//Memory Usage: 9.5 MB, less than 7.77% of Go online submissions for Maximum Subarray.
func dynamicProgramming(nums []int) int {
	max := nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

//greedy 贪心算法
func greedy(nums []int) int {
	max, sum := -math.MaxInt, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		max = maxFunc(max, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
