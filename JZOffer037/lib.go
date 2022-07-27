package JZOffer037

func asteroidCollision(asteroids []int) []int {
	res, resCap := make([]int, len(asteroids)), 0
	for _, v := range asteroids {
		resCap = p(res, resCap, v)
	}
	return res[:resCap]
}

func p(res []int, resCap, v int) int {
	for resCap >= 0 {
		// 初始 或 方向相同 或 反方向且相离
		if resCap == 0 || res[resCap-1]*v > 0 || res[resCap-1] < 0 && v > 0 {
			res[resCap] = v
			resCap++
			break
		}

		// 反方向且相撞
		if res[resCap-1] > 0 && v < 0 {
			//新的更大 减去res中元素
			if res[resCap-1] < v*-1 {
				res[resCap-1] = 0
				resCap--
				continue
			}
			//两数相同，减去res中元素
			if res[resCap-1] == v*-1 {
				res[resCap-1] = 0
				resCap--
				break
			}
			break
		}
	}
	return resCap
}
