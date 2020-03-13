package Q169_Majority_Element

// https://leetcode-cn.com/problems/majority-element/
func majorityElement(nums []int) int {
	count, maj := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if maj == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				maj = nums[i+1]
			}
		}
	}
	return maj
}
