package Q76_Minimum_Window_Substring

// minWindow 大佬的写法
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Window Substring.
//Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Minimum Window Substring.
func minWindow(s string, t string) string {
	n, m := len(s), len(t)

	if n < m {
		return ""
	}

	var need [58]int
	valid := 0

	for i := 0; i < m; i++ {
		if need[t[i]-'A'] == 0 {
			valid++
		}
		need[t[i]-'A']++
	}

	var freq [58]int
	curr := 0
	l, r := 0, 0
	res := ""

	for l < n {
		if r < n && valid != curr {
			freq[s[r]-'A']++
			if freq[s[r]-'A'] == need[s[r]-'A'] {
				curr++
			}
			r++
		} else { // young注：只在需要验证个数与已验证的字母个数相等时，进行窗口缩小
			freq[s[l]-'A']--
			if freq[s[l]-'A'] < need[s[l]-'A'] {
				curr--
			}
			l++
		}

		if curr == valid && (res == "" || len(res) >= r+1-l) {
			res = s[l:r]
		}
	}

	return res
}

// minWindow2
//Runtime: 304 ms, faster than 5.30% of Go online submissions for Minimum Window Substring.
//Memory Usage: 3 MB, less than 58.90% of Go online submissions for Minimum Window Substring.
func minWindow2(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	l, r, m, mt, res := 0, 0, make(map[int32]int), make(map[int32]int), ""
	for _, v := range t {
		m[v] = 0
		mt[v]++
	}
	for {
		if (Status(m, mt) == 2 || r == 0 || r-l <= len(t)) && r < len(s) {
			add(int32(s[r]), m)
			r++
			if Status(m, mt) == 0 && (len(res) == 0 || len(res) > r-l) {
				res = s[l:r]
			}

		} else if Status(m, mt) == 0 || Status(m, mt) == 1 || l < r {
			del(int32(s[l]), m)
			l++
			if Status(m, mt) == 0 && (len(res) == 0 || len(res) > r-l) {
				res = s[l:r]
			}

		} else {
			break
		}
	}
	return res
}

//Status 滑动窗口状态
// 0：刚好相等，需要缩小
// 1：已经过大，需要缩小
// 2：数量不够，需要扩大
func Status(m, mt map[int32]int) uint8 {
	hasOverFlow := true
	for k, v := range m {
		if v < mt[k] {
			return 2
		} else if v == mt[k] {
			hasOverFlow = false
		}
	}
	if hasOverFlow {
		return 1
	}
	return 0
}

func add(b int32, m map[int32]int) {
	if _, ok := m[b]; ok {
		m[b] += 1
	}
}

func del(b int32, m map[int32]int) {
	if _, ok := m[b]; ok {
		m[b] -= 1
	}
}
