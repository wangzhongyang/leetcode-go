package Q5_Longest_Palindromic_Substring

//// https://leetcode-cn.com/problems/longest-palindromic-substring/
//func longestPalindrome(s string) string {
//	slice := make([]string, 0, 4)
//
//	for _, char := range s {
//		slice = append(slice, "#")
//		slice = append(slice, string(char))
//	}
//
//	slice = append(slice, "#")
//
//	maxR := 0     //记录最长的半径
//	maxIndex := 0 //记录最长的 index
//	sliceLen := len(slice)
//
//	for index, _ := range slice {
//
//		if index >= 1 {
//
//			r := 0
//			i := index - 1
//			j := index + 1
//
//			for { //每一个字符 计算最长的半径
//
//				if i < 0 || j >= sliceLen {
//					break
//				}
//
//				if slice[i] == slice[j] {
//					r++
//					i--
//					j++
//				} else {
//					break
//				}
//
//			}
//
//			if r > maxR {
//				maxR = r
//				maxIndex = index
//			}
//
//		}
//	}
//
//	res := ""
//
//	for index, str := range slice {
//
//		if index >= (maxIndex-maxR) && index <= (maxIndex+maxR) && str != "#" {
//			res += str
//		}
//	}
//
//	return res
//}

//Runtime: 16 ms, faster than 55.23% of Go online submissions for Longest Palindromic Substring.
//Memory Usage: 2.6 MB, less than 79.39% of Go online submissions for Longest Palindromic Substring.
func longestPalindrome(s string) string {
	left, right, res := 0, 0, ""
	for i := 0; i < len(s); i++ {
		left, right = i, i
		for {
			if left == right {
				for {
					if left-1 >= 0 && s[left-1] == s[i] {
						left = left - 1
					} else if right+1 < len(s) && s[right+1] == s[i] {
						right = right + 1
					} else {
						if right+1-left > len(res) {
							res = s[left : right+1]
						}
						left, right = left-1, right+1
						break
					}
				}
			}
			if left >= 0 && right < len(s) && s[left] == s[right] {
				if right+1-left > len(res) {
					res = s[left : right+1]
				}
				left, right = left-1, right+1
			} else {
				break
			}
		}
	}
	return res
}
