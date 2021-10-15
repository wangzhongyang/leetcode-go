package Q704_Binary_Search

//Runtime: 50 ms, faster than 25.94% of Go online submissions for Binary Search.
//Memory Usage: 7 MB, less than 35.56% of Go online submissions for Binary Search.
func search(nums []int, target int) int {
	for i, j := 0, len(nums)-1; i <= j; {
		m := i + (j-i)/2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}

	return -1
}
