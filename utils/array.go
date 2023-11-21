package utils

func RemoveRep(slc []uint) []uint {
	if len(slc) < 1024 {
		return RemoveRepByLoop(slc)
	} else {
		return RemoveRepByMap(slc)
	}
}

func RemoveRepByMap(slc []uint) []uint {
	result := []uint{}
	tempMap := map[uint]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// 循环过滤
func RemoveRepByLoop(slc []uint) []uint {
	result := []uint{}
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}
