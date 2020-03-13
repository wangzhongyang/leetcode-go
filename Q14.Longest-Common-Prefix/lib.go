package Q14_Longest_Common_Prefix

import "strings"

// https://leetcode-cn.com/problems/longest-common-prefix/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
//Memory Usage: 2.4 MB, less than 25.00% of Go online submissions for Longest Common Prefix.
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res, i, isBreak, firstChar := make([]string, 0), 0, false, uint8(0)
	for {
		for k, v := range strs {
			if i == len(v) {
				isBreak = true
				break
			}
			if k == 0 {
				firstChar = v[i]
			} else {
				if firstChar != v[i] {
					isBreak = true
					break
				}
			}
			if k == len(strs)-1 {
				res = append(res, string(firstChar))
			}
		}
		if isBreak {
			break
		} else {
			i++
		}
	}
	return strings.Join(res, "")
}
