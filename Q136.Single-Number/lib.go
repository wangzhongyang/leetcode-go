package Q136_Single_Number

// https://leetcode-cn.com/problems/single-number/
func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
