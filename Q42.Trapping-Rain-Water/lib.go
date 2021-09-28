package Q42_Trapping_Rain_Water

func trap(height []int) int {
	list, current, res := make([]int, 0), 0, 0
	for current < len(list) {
		for len(list) != 0 && height[current] > height[list[len(list)-1]] {
			last := height[list[len(list)-1]]
			list = list[:len(list)-1]
			if len(list) == 0 {
				break
			}
			distance := current - list[len(list)-1] - 1
			boundedHeight := min(height[current], last) - height[last]
			res += distance * boundedHeight
		}
		list = append(list, height[current])
		current++
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//int top = st.top();
//st.pop();
//if (st.empty())
//break;
//int distance = current - st.top() - 1;
//int bounded_height = min(height[current], height[st.top()]) - height[top];
//ans += distance * bounded_height;
//
//作者：LeetCode
//链接：https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
