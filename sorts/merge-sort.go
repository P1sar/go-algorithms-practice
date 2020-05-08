package sorts

func MergeSort(arr []int64)  []int64 {
	if len(arr) <= 1 {
		return arr
	}
	c := MergeSort(arr[:len(arr)/2])
	d := MergeSort(arr[len(arr)/2:])
	res := make([]int64,len(c)+len(d))
	k := 0
	j := 0
	for i:=0; i<len(res); i++ {
		if k >= len(d) {
			res[i] = c[j]
			j++
			continue
		}
		if j >= len(c) {
			res[i] = d[k]
			k++
			continue
		}
		if c[j] > d[k] {
			res[i] = d[k]
			k++
		} else {
			res[i] = c[j]
			j++
		}

	}
	return res
}

