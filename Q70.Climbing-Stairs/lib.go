package Q70_Climbing_Stairs

//climbStairs 动态规划 f(n)=f(n-1)+f(n-2)
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
//Memory Usage: 1.9 MB, less than 77.52% of Go online submissions for Climbing Stairs.
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n1, n2, n3 := 1, 2, 0
	for i := 2; i < n; i++ {
		n3 = n1 + n2
		n1, n2 = n2, n3
	}
	return n2
}
