package Q9_Palindrome_Number

// https://leetcode-cn.com/problems/palindrome-number/
// Runtime: 16 ms, faster than 68.95% of Go online submissions for Palindrome Number.
// Memory Usage: 5.2 MB, less than 62.50% of Go online submissions for Palindrome Number.
// 与Q7重复
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp, n := x, 0
	for temp != 0 {
		n = n*10 + temp%10
		temp = temp / 10
	}
	if n == x {
		return true
	}
	return false
}
