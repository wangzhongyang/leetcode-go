package Q3_ongest_Substring_Without_Repeating_Characters

import "strings"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/
// 虽然暴力，但是舒服
//Runtime: 32 ms, faster than 28.03% of Go online submissions for Longest Substring Without Repeating Characters.
//Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
//func lengthOfLongestSubstring(s string) int {
//	str := ""  // 最长
//	begin := 0 // 起始
//	end := 0   // 结束
//	tempStr := ""
//	for {
//		if end == len(s) {
//			break
//		}
//		temp := string(s[end])
//		tempStr = s[begin:end]
//		if strings.Contains(tempStr, temp) {
//			begin += 1
//			end = begin
//			continue
//		}
//		end += 1
//		tempStr = s[begin:end]
//		if len(tempStr) > len(str) {
//			str = tempStr
//		}
//	}
//	return len(str)
//}

// lengthOfLongestSubstring
// Runtime: 6 ms, faster than 82.25% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	n1, maxLen := 0, 1
	for i := 1; i < len(s); i++ {
		for strings.Contains(s[n1:i], string(s[i])) {
			n1++
			if n1 == i {
				break
			}
		}
		if i-n1+1 > maxLen {
			maxLen = i - n1 + 1
		}
	}
	return maxLen
}
