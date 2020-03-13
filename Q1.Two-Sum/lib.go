package Q1_Two_Sum

// https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		temp := target - v
		if k1, ok := numMap[temp]; ok && k1 != k {
			return []int{k1, k}
		}
		numMap[v] = k
	}
	return nil
}
