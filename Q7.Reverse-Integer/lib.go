package Q7_Reverse_Integer

// https://leetcode-cn.com/problems/reverse-integer/
//Runtime: 4 ms, faster than 42.15% of Go online submissions for Reverse Integer.
//Memory Usage: 2.2 MB, less than 80.00% of Go online submissions for Reverse Integer.
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
//Memory Usage: 2.2 MB, less than 80.00% of Go online submissions for Reverse Integer.
func reverse(x int) int {
	var n int
	for x != 0 {
		n = n*10 + x%10
		x = x / 10
	}
	// 使用这种if 会每次求值
	//Runtime: 4 ms, faster than 42.15% of Go online submissions for Reverse Integer.
	//Memory Usage: 2.2 MB, less than 80.00% of Go online submissions for Reverse Integer.
	//if n > math.MaxInt32 || n < math.MinInt32
	if n > 2147483647 || n < -2147483648 {
		return 0
	}
	return n
}
