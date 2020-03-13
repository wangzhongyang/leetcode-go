package Q3_ongest_Substring_Without_Repeating_Characters

import (
	"strings"
)

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/
// 虽然暴力，但是舒服
//Runtime: 32 ms, faster than 28.03% of Go online submissions for Longest Substring Without Repeating Characters.
//Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstring(s string) int {
	str := ""  // 最长
	begin := 0 // 起始
	end := 0   // 结束
	tempStr := ""
	for {
		if end == len(s) {
			break
		}
		temp := string(s[end])
		tempStr = s[begin:end]
		if strings.Contains(tempStr, temp) {
			begin += 1
			end = begin
			continue
		}
		end += 1
		tempStr = s[begin:end]
		if len(tempStr) > len(str) {
			str = tempStr
		}
	}
	return len(str)
}

// 想通过map保存临时字符串已有的字符，实现复杂且空间及时间消耗无法保证，放弃了
//func lengthOfLongestSubstring(s string) int {
//	str := ""  // 最长
//	begin := 0 // 起始
//	end := 0   // 结束
//	alreadyMap := make(map[string]int)
//	tempStr := ""
//	for {
//		for {
//			if end > len(s)-1 {
//				break
//			}
//			temp := string(s[end])
//			//pwwkew
//			fmt.Println(temp, alreadyMap)
//			if k, ok := alreadyMap[temp]; ok {
//				//tempStr = s[begin : end+1]
//				if k+1 == end {
//					begin = end
//				} else {
//					begin = k + 1
//				}
//				end = end + 1
//				if len(tempStr) < len(str) {
//					alreadyMap = make(map[string]int)
//				}
//				alreadyMap[temp] = k
//				break
//			}
//			alreadyMap[temp], tempStr = end, s[begin:end+1]
//			end = end + 1
//		}
//		if len(tempStr) > len(str) {
//			str = tempStr
//		}
//		if end > len(s)-1 {
//			break
//		}
//	}
//	fmt.Println("str:		", str)
//	//fmt.Println("test:", "0123456"[:3])
//	return len(str)
//}
