package Q11_Container_With_Most_Water

// https://leetcode-cn.com/problems/container-with-most-water/
//Runtime: 12 ms, faster than 92.86% of Go online submissions for Container With Most Water.
//Memory Usage: 5.8 MB, less than 46.67% of Go online submissions for Container With Most Water.
func maxArea(height []int) int {
	area, begin, end := 0, 0, len(height)-1
	for begin < end {
		area = max(area, min(height[begin], height[end])*(end-begin))
		if height[begin] < height[end] {
			begin++
		} else {
			end--
		}
	}
	return area
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
