package Q151_Reverse_Words_in_a_String

import (
	"strings"
)

func reverseWords(s string) string {
	return f2(s)
}

// 与f1方法的区别主要在于字符串的拼接造成的内存消耗不一致
//Runtime: 4 ms, faster than 41.38% of Go online submissions for Reverse Words in a String.
//Memory Usage: 8.2 MB, less than 7.59% of Go online submissions for Reverse Words in a String.
func f2(s string) string {
	arr, buf := strings.Split(s, " "), ""
	for i := len(arr) - 1; i >= 0; i-- {
		if strings.TrimSpace(arr[i]) == "" {
			continue
		}
		buf += arr[i] + " "
	}
	return buf[:len(buf)-1]
}

//Runtime: 4 ms, faster than 41.38% of Go online submissions for Reverse Words in a String.
//Memory Usage: 3.5 MB, less than 48.28% of Go online submissions for Reverse Words in a String.
func f1(s string) string {
	arr, buf := make([]string, 0), new(strings.Builder)
	for _, v := range s {
		if v == 32 {
			if buf.Len() != 0 {
				arr = append(arr, buf.String())
				buf.Reset()
			}
			continue
		}
		buf.WriteString(string(v))
	}
	if buf.Len() != 0 {
		arr = append(arr, buf.String())
		buf.Reset()
	}
	for i := len(arr) - 1; i >= 0; i-- {
		buf.WriteString(arr[i])
		buf.WriteString(" ")
	}
	return buf.String()[:buf.Len()-1]
}
