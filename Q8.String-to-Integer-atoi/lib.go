package Q8_String_to_Integer_atoi

import (
	"regexp"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/string-to-integer-atoi/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).
//Memory Usage: 4.6 MB, less than 12.50% of Go online submissions for String to Integer (atoi).
func myAtoi(str string) int {
	reg, _ := regexp.Compile(`^[\+\-]?\d+`)
	n, _ := strconv.Atoi(reg.FindString(strings.TrimLeft(str, " ")))
	if n > 2147483647 {
		return 2147483647
	}
	if n < -2147483648 {
		return -2147483648
	}
	return n
}
