package Q912_Sort_an_Array

import "fmt"

func quickSort(nums []int) []int {
	quickSortRecursio(nums, 0, len(nums)-1)
	return nums
}

func swap(nums []int, l, r int) {
	nums[l], nums[r] = nums[r], nums[l]
}

func quickSortRecursio(nums []int, left, right int) {
	if left < right {
		index := partition(nums, left, right)
		//fmt.Println(nums, left, index, right)
		quickSortRecursio(nums, left, index-1)
		quickSortRecursio(nums, index+1, right)
	}
}

//partition 分区
// 返回index的值为位置确定的元素，后续递归中不对该位置进行处理
func partition(nums []int, left, right int) int {
	fmt.Println("partition", nums, left, right)
	pivot, index := left, left+1
	for index <= right {
		if nums[pivot] < nums[index] {
			swap(nums, index, right)
			right--
		} else {
			index++
		}
	}
	swap(nums, pivot, index-1)
	return index - 1
}
