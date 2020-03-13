package Q6_ZigZag_Conversion

import (
	"strings"
)
// https://leetcode-cn.com/problems/zigzag-conversion/
//Runtime: 8 ms, faster than 72.81% of Go online submissions for ZigZag Conversion.
//Memory Usage: 6 MB, less than 71.43% of Go online submissions for ZigZag Conversion.
func convert(s string, numRows int) string {
	sss := make([]string, 0)
	number := 2 * (numRows - 1) // 每轮元素个数
	if number == 0 {
		number = 1
	}
	for i := 0; i < numRows; i++ {
		// 第一行
		if i == 0 {
			temp := 0
			for {
				if temp*number >= len(s) {
					break
				}
				sss = append(sss, string(s[temp*number]))
				temp++
			}
			continue
		}
		// 最后一行
		if i == numRows-1 {
			temp := 0
			for {
				key := temp*number + numRows - 1
				if key >= len(s) {
					break
				}
				sss = append(sss, string(s[key]))
				temp++
			}
		} else {
			// 中间行
			temp := 0
			for {
				key := temp*number + i
				if key >= len(s) {
					break
				}
				sss = append(sss, string(s[key]))
				key = (temp+1)*number - i
				if key >= len(s) {
					break
				}
				sss = append(sss, string(s[key]))
				temp++
			}
		}
	}
	return strings.Join(sss, "")
}
