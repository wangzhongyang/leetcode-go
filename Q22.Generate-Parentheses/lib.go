package Q22_Generate_Parentheses

import (
	"fmt"
	"sort"
)

// todo 题目意思理解错了
func generateParenthesis(n int) []string {
	left, right := 1, 0
	res := make([]string, 0)
	if n == 0 {
		return res
	} else {
		res = []string{"("}
	}
	for i := 1; i <= 2*n; i++ {
		if left-right == 0 { // 已加入成对的(),每个字符串加(
			for k, v := range res {
				res[k] = fmt.Sprintf("%s%s", v, "(")
				left++
			}
		} else { // 每个字符串分别添加"("")"
			for k, v := range res {
				if left <= n {
					res = append(res, fmt.Sprintf("%s%s", v, "("))
					left++
				}
				if left > right {
					res[k] = fmt.Sprintf("%s%s", v, ")")
					right++
				}
			}
		}
	}
	//a := []string{
	//	"((()))",
	//	"(()())",
	//	"(())()",
	//	"()(())",
	//	"()()()",
	//}
	sort.Strings(res)
	fmt.Println(res)
	return res
}
