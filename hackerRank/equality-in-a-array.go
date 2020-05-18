package hackerrank


func equalizeArray(arr []int32) int32 {
	valsMap := make(map[int32]int32)
	for _, v := range arr {
		if _, ok:=valsMap[v]; ok {
			valsMap[v] += 1
		} else {
			valsMap[v] = 1
		}
	}
	var maxVal int32 = 0
	for _,v := range valsMap {
		if v > maxVal {
			maxVal = v
		}
	}
	return int32(len(arr)) - maxVal
}
