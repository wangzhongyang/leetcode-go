package Q300_Longest_Increasing_Subsequence

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	//return dp(nums)
	return greedy(nums)
}

// 贪心算法 时间复杂度O(n*log n)
// 需要使用 搜索算法的case:  [4, 10, 4, 3, 8, 9]
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Increasing Subsequence.
//Memory Usage: 2.4 MB, less than 40.00% of Go online submissions for Longest Increasing Subsequence.
func greedy(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := []int{nums[0]}
	for _, v := range nums {
		l := len(res) - 1
		if v > res[l] {
			res = append(res, v)
		} else if k := sort.SearchInts(res, v); res[k] != v { // 此处为二分搜索
			res[k] = v
		}
	}
	return len(res)
}

// 动态规划 时间复杂度O(n^2)
//Runtime: 8 ms, faster than 64.32% of Go online submissions for Longest Increasing Subsequence.
//Memory Usage: 2.4 MB, less than 60.00% of Go online submissions for Longest Increasing Subsequence
func dp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, res := make([]int, len(nums)), 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && max < dp[j]+1 {
				max = dp[j] + 1
			}
		}
		if res < max {
			res = max
		}
		dp[i] = max
	}
	return res
}

func Print(nums []int) {
	for _, v := range nums {
		fmt.Print(fmt.Sprintf("%d\t", v))
	}
	fmt.Println("")
}
