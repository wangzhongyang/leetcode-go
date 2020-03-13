package Q20_Valid_Parentheses

//https://leetcode-cn.com/problems/valid-parentheses/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
//Memory Usage: 2.1 MB, less than 40.00% of Go online submissions for Valid Parentheses.
func isValid(s string) bool {
	var node *Node
	valid := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for _, v := range s {
		if val, ok := valid[string(v)]; ok {
			if node != nil && node.value == val {
				node = node.next
				continue
			} else {
				return false
			}
		} else {
			node = &Node{
				value: string(v),
				next:  node,
			}
		}
	}
	if node == nil {
		return true
	} else {
		return false
	}
}

type Node struct {
	value string
	next  *Node
}
