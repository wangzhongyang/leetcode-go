package Q31_Next_Permutation

//nextPermutation
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Next Permutation.
//Memory Usage: 2.5 MB, less than 43.55% of Go online submissions for Next Permutation.
func nextPermutation(nums []int) {
	// 从右边开始的非递增第一个数x
	tmp, minIndex, maxIndex := nums[len(nums)-1], 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < tmp {
			minIndex = i
			break
		} else {
			tmp = nums[i]
		}
	}
	// 非递增第一个数右边比x大的第一个数
	for i := len(nums) - 1; i >= 0 && i > minIndex; i-- {
		if nums[i] > nums[minIndex] {
			maxIndex = i
			break
		}
	}
	if maxIndex == 0 { // 不存在，则该数组为倒序数组，翻转整个数组
		begin, over := 0, len(nums)-1
		for begin < over {
			nums[begin], nums[over] = nums[over], nums[begin]
			begin++
			over--
		}
		return
	}
	// 存在，交换min和max的位置，并翻转minIndex后的数组
	nums[minIndex], nums[maxIndex] = nums[maxIndex], nums[minIndex]
	begin, over := minIndex+1, len(nums)-1
	for begin < over {
		nums[begin], nums[over] = nums[over], nums[begin]
		begin++
		over--
	}
}
