package Q17_Letter_Combinations_of_a_Phone_Number

import "fmt"

var dictionary = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 没其他方法，暴力for循环
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
//Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	if len(digits) == 1 {
		return dictionary[digits]
	}
	res := dictionary[string(digits[0])]
	if len(digits) > 1 {
		for i := 1; i < len(digits); i++ {
			temp := make([]string, 0)
			for _, v1 := range res {
				for _, v2 := range dictionary[string(digits[i])] {
					temp = append(temp, fmt.Sprintf("%s%s", v1, v2))
				}
			}
			res = temp
		}
	}
	return res
}
