package Q88_Merge_Sorted_Array

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m, 0; i < len(nums1); i++ {
		nums1[i] = nums2[j]
		j++
	}
	sort.Ints(nums1)
}
