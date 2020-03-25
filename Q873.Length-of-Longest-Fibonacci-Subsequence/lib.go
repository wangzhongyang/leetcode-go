package Q873_Length_of_Longest_Fibonacci_Subsequence

import (
	"sort"
)

func lenLongestFibSubseq(A []int) int {
	return b(A)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 根据菲波齐纳数列性质
//Runtime: 240 ms, faster than 6.98% of Go online submissions for Length of Longest Fibonacci Subsequence.
//Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Length of Longest Fibonacci Subsequence.
//双重for循环加个二分搜索 时间复杂度：O(N^2 log n),空间复杂度为O(1)
func b(A []int) int {
	if len(A) <= 3 {
		return 0
	}
	max, tempMax := 0, 2
	for i := 0; i < len(A)-2; i++ {
		for j := i + 1; j < len(A)-1; j++ {
			x, y := A[i], A[j]
			for {
				z := sort.SearchInts(A, x+y)
				if z == len(A) || A[z] != x+y {
					break
				} else {
					x, y = y, x+y
					tempMax++
				}
			}
			max = Max(max, tempMax)
			tempMax = 2
		}
	}
	return max
}
