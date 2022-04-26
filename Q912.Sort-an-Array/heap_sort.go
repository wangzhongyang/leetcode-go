package Q912_Sort_an_Array

func heapSort(list []int) []int {
	// 生成堆
	heapify(list)
	// 将堆顶放在最后，下沉新堆顶
	for i := len(list) - 1; i >= 0; i-- {
		swap(list, 0, i)
		down(list, 0, i-1)
	}
	return list
}

// heapify 生成堆
func heapify(list []int) {
	l := len(list)
	for i := (l - 1) / 2; i >= 0; i-- {
		down(list, i, l-1)
	}
}

// down 当前位置下沉
// i：当前位置
// end：结束位置
func down(list []int, i, end int) {
	for 2*i+1 <= end {
		j := 2*i + 1
		if j+1 <= end && list[j+1] > list[j] { // 选出两个子节点中较大的一个
			j = j + 1
		}
		if list[i] < list[j] {
			swap(list, i, j)
		} else {
			break
		}
		i = j
	}
}
