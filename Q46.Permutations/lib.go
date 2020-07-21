package Q46_Permutations

//https://leetcode-cn.com/problems/permutations/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations.
//Memory Usage: 3 MB, less than 100.00% of Go online submissions for Permutations.
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res, path := make([][]int, 0), make([]int, 0)
	dfs(&res, nums, path)
	return res
}

func dfs(res *[][]int, nums, path []int) {
	if len(path) == len(nums) {
		pathTemp := make([]int, len(path))
		copy(pathTemp, path)
		*res = append(*res, pathTemp)
		return
	}
	begin := len(path)
	for k, v := range nums[begin:] {
		numsTemp := make([]int, len(nums))
		copy(numsTemp, nums)
		path = append(path, v)
		moveItem(numsTemp, begin, k+begin)
		dfs(res, numsTemp, path)
		path = path[:len(path)-1]
	}
}

//a := []int{0, 1, 2, 3, 4, 5, 6}
//change(a, 1, 5)
//[0 5 1 2 3 4 6]
func moveItem(nums []int, k1, k2 int) {
	for k2 > k1 {
		nums[k2], nums[k2-1] = nums[k2-1], nums[k2]
		k2--
	}
}
