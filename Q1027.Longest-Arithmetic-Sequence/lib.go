package Q1027_Longest_Arithmetic_Sequence

// todo 本地可以运行，网上无法通过
func longestArithSeqLength(A []int) int {
	return dp(A)
	//return violence(A)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 时间复杂度：O(n^2)
//Runtime: 488 ms, faster than 25.93% of Go online submissions for Longest Arithmetic Sequence.
//Memory Usage: 104.8 MB, less than 100.00% of Go online submissions for Longest Arithmetic Sequence.
func dp(A []int) int {
	if len(A) == 2 {
		return 2
	}
	mem, res := make([]map[int]int, len(A)), 1
	mem[0] = make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if mem[j] == nil {
				mem[j] = make(map[int]int)
			}
			v := A[j] - A[i]
			mem[j][v] = Max(mem[i][v]+1, mem[j][v])
			res = Max(res, mem[j][v])
		}
	}
	return res + 1
}

// 暴力求解 O(n^3) 超时
func violence(A []int) int {
	if len(A) == 2 {
		return 2
	}
	res := 2
	for i := 0; i < len(A)-2; i++ {
		for j := i + 1; j < len(A)-1; j++ {
			pre, temp, d := j, 2, A[j]-A[i]
			for k := j + 1; k < len(A); k++ {
				if A[k]-A[pre] == d {
					temp++
					pre = k
				}
			}
			res = Max(res, temp)
		}
	}
	return res
}
