package Q93_Restore_IP_Addresses

import (
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/restore-ip-addresses/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Restore IP Addresses.
//Memory Usage: 2.1 MB, less than 66.67% of Go online submissions for Restore IP Addresses
func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}
	cur, res := make([]string, 0), make([]string, 0)
	dfs(s, 0, cur, &res)
	return res
}

func dfs(s string, index int, cur []string, res *[]string) {
	if len(cur) == 4 {
		if index == len(s) {
			*res = append(*res, strings.Join(cur, "."))
		}
	}
	for i := 1; i <= 3; i++ {
		tempIndex := index + i
		if tempIndex > len(s) {
			break
		}
		ii, err := strconv.Atoi(s[index:tempIndex])
		if err != nil || ii > 255 || (len(s[index:tempIndex]) > 1 && s[index:tempIndex][0] == '0') {
			break
		}
		if len(s)-tempIndex-1 > 3*(4-len(cur)-1) {
			continue
		}
		cur = append(cur, s[index:tempIndex])
		dfs(s, tempIndex, cur, res)
		cur = cur[:len(cur)-1]
	}
}
