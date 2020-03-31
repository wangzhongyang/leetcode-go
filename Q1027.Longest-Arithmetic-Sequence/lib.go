package Q1027_Longest_Arithmetic_Sequence

import "fmt"

// todo 本地可以运行，网上无法通过
func longestArithSeqLength(A []int) int {

	return violence(A)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 暴力求解
func violence(A []int) int {
	fmt.Println(A)
	if len(A) == 2 {
		return 2
	}
	aMap := make(map[int]int)
	for K, v := range A {
		if _, ok := aMap[v]; !ok {
			aMap[v] = K
		}
	}
	var x, y, z, max, maxTemp, yKey int
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			x, y, yKey, maxTemp = A[i], A[j], j, 2
			if x == y {
				continue
			}
			z = 2*y - x
			for {
				if zKey, ok := aMap[z]; ok && zKey > yKey {
					x, y, yKey = y, z, zKey
					z = 2*y - x
					maxTemp++
				} else {
					max = Max(max, maxTemp)
					break
				}
			}
		}
	}
	fmt.Println(max)
	return max
}
