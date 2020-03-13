package Q4_Median_of_Two_Sorted_Arrays

import "sort"

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num := append(nums1, nums2...)
	sort.Ints(num)
	temp := len(num) / 2
	if median := len(num) % 2; median == 0 {
		return float64(num[temp-1]+num[temp]) / 2
	} else {
		return float64(num[temp])
	}
}
