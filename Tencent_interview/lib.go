package Tencent_interview

//实现f(n)函数，f(n）返回1-n之间整数不带7数字的个数；如17就是一个带7的数字，18是不带7的数字；
func Q1(n int) int {
	if n == 0 {
		return 0
	}
	digit, res, high, cur, low := 1, 0, n/10, n%10, 0
	for high != 0 || cur != 0 {
		if cur < 7 {
			res += high * digit
		} else if cur == 7 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return n - res
}

//求一个正整数数组，数组长度为10，输入一个sum，求出数组中元素相加之和等于sum的所有组合；（元素可能有重复；请输出数组元素的下标
func Q2(nums []int, target int) [][]int {
	res, map1 := make([][]int, 0), make(map[int][]int)
	//构建值与键的map
	for k, v := range nums {
		if arr, ok := map1[v]; ok {
			arr = append(arr, k)
			map1[v] = arr
		} else {
			map1[v] = []int{k}
		}
	}
	//查找需要的值及在本次组合中为使用的键
	for k, v := range nums {
		if arr, ok := map1[target-v]; ok {
			map2 := make(map[int]int)
			for _, key := range arr {
				if _, ok1 := map2[key]; !ok1 && key != k && key > k {
					res = append(res, []int{k, key})
					map2[key] = 1
				}
			}
		}
	}
	return res
}

func maxProfit(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max, i := 0, 0
	for {
		if i+1 >= len(arr) {
			break
		}
		// 比下一个大，
		if arr[i] >= arr[i+1] {
			i++
			continue
		}
		// 找到最大一个求差值
		m := 0
		for ii := i; ii < len(arr); ii++ {
			if arr[ii] > m {
				m = arr[ii]
			}
		}
		if temp := m - arr[i]; temp > max {
			max = temp
		}
		i++
	}
	return max
}
