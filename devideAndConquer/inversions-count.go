package devideAndConquer

import (
	"algs/sorts"
)

func CountInversions(arr []int64) ([]int64, int64) {
	if len(arr) == 1 {
		return arr, 0
	}
	sortedLeft, inversionsLeft := CountInversions(arr[:len(arr)/2])
	sortedRight, inversionsRight := CountInversions(arr[len(arr)/2:])
	sortedArr, totalInversions := mergeAndCountInversions(sortedLeft, sortedRight)


	return sortedArr, totalInversions+inversionsLeft+inversionsRight
}

func mergeAndCountInversions(b, c []int64) ([]int64, int64) {
	b = sorts.MergeSort(b)
	c = sorts.MergeSort(c)
	res := make([]int64, len(c)+len(b))
	k := 0
	j := 0
	inversions := 0
	for i:=0; i<len(res); i++ {
		if k >= len(b) {
			res[i] = c[j]
			j++
			continue
		}
		if j >= len(c) {
			res[i] = b[k]
			k++
			continue
		}
		if c[j] < b[k] {
			res[i] = c[j]
			j++
			inversions += len(b) -k
		} else {
			res[i] = b[k]
			k++
		}
	}
	return res, int64(inversions)
}
