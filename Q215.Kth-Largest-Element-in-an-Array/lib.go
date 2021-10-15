package Q215_Kth_Largest_Element_in_an_Array

//findKthLargest
//Runtime: 4 ms, faster than 99.05% of Go online submissions for Kth Largest Element in an Array.
//Memory Usage: 3.5 MB, less than 100.00% of Go online submissions for Kth Largest Element in an Array.
func findKthLargest(nums []int, k int) int {
	// 根据nums[:k]构建一个长度为K的最小堆
	for i := k / 2; i >= 0; i-- {
		adjust(nums[:k], i, k)
	}
	// 如果nums[i]大于堆顶，则替换堆顶重新交换堆
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			swap(nums, i, 0)
			adjust(nums[:k], 0, k)
		}
	}
	return nums[0]
}

func adjust(nums []int, i, length int) {
	l := 2*i + 1
	tmpI := i
	if l < length && l+1 < length && nums[l] > nums[l+1] {
		l++
	}
	if l < length && nums[i] > nums[l] {
		tmpI = l
	}
	if tmpI != i {
		swap(nums, tmpI, i)
		adjust(nums, tmpI, length)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// heapSort 堆排序，此处为最小堆最后得到逆序
// 步骤：
// 1. 创建一个堆 H[0……n-1],并且把最小值放到最上面；
// 2. 把堆首（最大值）和堆尾互换；
// 3. 把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
// 4. 重复步骤 2，直到堆的尺寸为 1。
func heapSort(arr []int) []int {
	length := len(arr)
	for i := length / 2; i >= 0; i-- {
		heapSortAdjust(arr, i, length)
	}
	for i := length - 1; i >= 0; i-- {
		swap(arr, i, 0)
		heapSortAdjust(arr, 0, i)
	}
	return arr
}

func heapSortAdjust(arr []int, i, length int) {
	l := 2*i + 1
	tmpI := i
	if l < length && l+1 < length && arr[l] > arr[l+1] {
		l++
	}
	if l < length && arr[i] > arr[l] {
		tmpI = l
	}
	if tmpI != i {
		swap(arr, tmpI, i)
		heapSortAdjust(arr, tmpI, length)
	}
}
