package Q873_Length_of_Longest_Fibonacci_Subsequence

import (
	"sort"
)

func lenLongestFibSubseq(A []int) int {
	//return b(A)
	return dp(A)
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

// 动态规划 if A[i] + A[j] == A[k]{ longest[j, k] = longest[i, j] + 1 }
//Runtime: 60 ms, faster than 76.19% of Go online submissions for Length of Longest Fibonacci Subsequence.
//Memory Usage: 4.2 MB, less than 100.00% of Go online submissions for Length of Longest Fibonacci Subsequence.
func dp(A []int) int {
	if len(A) < 3 {
		return 0
	}
	n := len(A)
	index, longest, ans := make(map[int]int), make(map[int]int), 0
	for k, v := range A {
		index[v] = k
	}
	for k := 0; k < n; k++ {
		for j := 0; j < k; j++ {
			if i, ok := index[A[k]-A[j]]; ok && i < j {
				cand, ok := longest[i*n+j]
				if ok {
					cand += 1
				} else {
					cand = 3
				}
				longest[j*n+k] = cand
				ans = Max(cand, ans)
			}
		}
	}
	return ans
}
