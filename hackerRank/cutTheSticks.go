package hackerrank

import "fmt"

func CutTheSticks(arr []int32) []int32 {
	sortedArr := mergeSort(arr)
	first := sortedArr[0]
	last := sortedArr[len(sortedArr)-1]
	resArr := make([]int32, 0)
	resArr = append(resArr, int32(len(arr)))
	for first != last {
		c := int32(0)
		lowest := sortedArr[0]
		newArr := make([]int32, 0)
		for i := 0; i < len(sortedArr); i++ {
			if sortedArr[i] > lowest {
				newArr = append(newArr, sortedArr[i]-lowest)
				c++
			}
		}
		resArr = append(resArr, c)
		sortedArr = newArr
		first = sortedArr[0]
		last = sortedArr[len(sortedArr)-1]
		fmt.Println(first, last)
	}
	return resArr
}


func mergeSort(arr []int32)  []int32 {
	if len(arr) <= 1 {
		return arr
	}
	c := mergeSort(arr[:len(arr)/2])
	d := mergeSort(arr[len(arr)/2:])
	res := make([]int32,len(c)+len(d))
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