package Q43_Multiply_Strings

//使用数组临时存放临时结果
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Multiply Strings.
//Memory Usage: 3.1 MB, less than 30.77% of Go online submissions for Multiply Strings.
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	resArr := make([]int, len(num1)+len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		n2 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := int(num1[j] - '0')
			sum := n2*n1 + resArr[i+j+1]
			resArr[i+j+1] = sum % 10
			resArr[i+j] += sum / 10
		}
	}

	res := ""
	for k, v := range resArr {
		if k == 0 && v == 0 {
			continue
		}
		res += string(v + '0')
	}

	return res
}

// 直接根据乘法运算方式，调用addString 得到结果
//Runtime: 1488 ms, faster than 5.64% of Go online submissions for Multiply Strings.
//Memory Usage: 6.7 MB, less than 7.69% of Go online submissions for Multiply Strings.
//var zeroString = "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
//
//func multiply(num1 string, num2 string) string {
//	res := "0"
//	for i := len(num1) - 1; i >= 0; i-- {
//		tempi, tempj := "0", ""
//		ni, _ := strconv.Atoi(string(num1[i]))
//		for j := len(num2) - 1; j >= 0; j-- {
//			nj, _ := strconv.Atoi(string(num2[j]))
//			//tempFloat += nj * ni * int(math.Pow10(len(num2)-1-j)) // 会造成超过int64范围
//			tempj = addZero(len(num2)-1-j, fmt.Sprint(nj*ni))
//			tempi = addStrings(tempi, tempj)
//		}
//		//temp = strconv.Itoa(tempFloat * int(math.Pow10(len(num1)-1-i))) // 会造成超过int64范围
//		tempi = addZero(len(num1)-1-i, tempi)
//		res = addStrings(res, tempi)
//	}
//	return res
//}
//
//// 使用截取字符串的方式，可以在长大字符串相乘节约一半的时间
//func addZero(n int, s string) string {
//	return fmt.Sprintf("%s%s", s, zeroString[:n])
//}
//
//func addStrings(num1 string, num2 string) string {
//	i, j, carry, res := len(num1)-1, len(num2)-1, 0, make([]string, 0)
//	for i >= 0 || j >= 0 {
//		ni, nj := 0, 0
//		if i >= 0 {
//			ni, _ = strconv.Atoi(string(num1[i]))
//		}
//		if j >= 0 {
//			nj, _ = strconv.Atoi(string(num2[j]))
//		}
//
//		temp := ni + nj + carry
//		carry = temp / 10
//		res = append(res, fmt.Sprint(temp%10))
//		i--
//		j--
//	}
//	if carry == 1 {
//		res = append(res, fmt.Sprint(1))
//	}
//	i, j = 0, len(res)-1
//	for i < j {
//		res[i], res[j] = res[j], res[i]
//		i++
//		j--
//	}
//
//	if s := strings.TrimLeft(strings.Join(res, ""), "0"); len(s) == 0 { // 去除左边的0
//		return "0"
//	} else {
//		return s
//	}
//}
