package Q567_Permutation_in_String

//https://leetcode-cn.com/problems/permutation-in-string/
//Runtime: 4 ms, faster than 71.15% of Go online submissions for Permutation in String.
//Memory Usage: 2.7 MB, less than 92.86% of Go online submissions for Permutation in String.
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Array, s2Array := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		s1Array[s1[i]-'a']++
		s2Array[s2[i]-'a']++
	}
	for i := 0; i < (len(s2) - len(s1)); i++ {
		if matches(s1Array, s2Array) {
			return true
		}
		// 移动窗口
		s2Array[s2[i+len(s1)]-'a']++
		s2Array[s2[i]-'a']--
	}
	return matches(s1Array, s2Array)
}

func matches(s1Array, s2Array [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1Array[i] != s2Array[i] {
			return false
		}
	}
	return true
}
