package Q33_Search_in_Rotated_Sorted_Array

//search
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
//Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Search in Rotated Sorted Array.
func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, len(nums)/2
	for l <= r {
		//fmt.Println(l, r, mid, target, nums)
		if nums[mid] == target {
			return mid
		}
		if nums[l] < nums[mid] && nums[mid] > target && nums[l] <= target { // 左边有序
			r = mid - 1
		} else if nums[l] > nums[mid] && (nums[l] <= target || nums[mid] > target) {
			r = mid - 1
		} else {
			l = mid + 1
		}
		mid = (r + l) / 2
	}
	return -1
}
