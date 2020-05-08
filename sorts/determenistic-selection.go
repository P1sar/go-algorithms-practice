package sorts

import (
	"math/rand"
	"time"
)

func DSelect(arr []int64, i int64) int64 {
	rand.Seed(time.Now().UnixNano())
	lenArr := int64(len(arr))

	if lenArr <= 1 { return arr[0] }

	if lenArr == 5 {
		sortedArr := MergeSort(arr)
		return sortedArr[len(sortedArr)/2]
	}

	arr1, arr2, arr3, arr4, arr5 := breakToFivePartsAndMergeSort(arr, lenArr)
	medianArr := []int64{arr1[len(arr1)/2], arr2[len(arr2)/2], arr3[len(arr3)/2], arr4[len(arr4)/2], arr5[len(arr5)/2]}
	medianOfMedian := DSelect(medianArr, int64(len(medianArr)/2))

	var medianIndex int64

	for i := int64(0); i < lenArr; i++ {
		if i == medianOfMedian {
			medianIndex = i
		}
	}

	pivotIndex := rearrangeAroundPivot(arr, medianIndex)


	if pivotIndex == i {
		return arr[i]
	}
	if pivotIndex > i {
		return RSelect(arr[:pivotIndex], i)
	} else {
		return RSelect(arr[pivotIndex:], i-pivotIndex)
	}
}

func breakToFivePartsAndMergeSort(arr []int64, l int64) ([]int64, []int64, []int64, []int64, []int64) {
	partSize := l/5
	partModulo := l%5
	return MergeSort(arr[:partSize]), MergeSort(arr[partSize:2*partSize]), MergeSort(arr[2*partSize:3*partSize]), MergeSort(arr[3*partSize:4*partSize]), MergeSort(arr[4*partSize:5*partSize+partModulo])
}
