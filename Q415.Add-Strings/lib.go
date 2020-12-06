package Q415_Add_Strings

import (
	"fmt"
	"strconv"
	"strings"
)

//https://leetcode.com/problems/add-strings/
//Runtime: 4 ms, faster than 62.09% of Go online submissions for Add Strings.
//Memory Usage: 3.9 MB, less than 46.67% of Go online submissions for Add Strings.
func addStrings(num1 string, num2 string) string {
	i, j, carry, res := len(num1)-1, len(num2)-1, 0, make([]string, 0)
	for i >= 0 || j >= 0 {
		ni, nj := 0, 0
		if i >= 0 {
			ni, _ = strconv.Atoi(string(num1[i]))
		}
		if j >= 0 {
			nj, _ = strconv.Atoi(string(num2[j]))
		}

		temp := ni + nj + carry
		carry = temp / 10
		res = append(res, fmt.Sprint(temp%10))
		i--
		j--
	}
	if carry == 1 {
		res = append(res, fmt.Sprint(1))
	}
	i, j = 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return strings.Join(res, "")
}
