package sorts

import (
	"math/rand"
	"time"
)


func RSelect(arr []int64, i int64) int64 {
	rand.Seed(time.Now().UnixNano())
	lenArr := int64(len(arr))
	if lenArr <= 1 { return arr[0] }

	pivotIndex := rand.Int63n(lenArr)

	pivotIndex = rearrangeAroundPivot(arr, pivotIndex)

	if pivotIndex == i {
		return arr[i]
	}
	if pivotIndex > i {
		return RSelect(arr[:pivotIndex], i)
	} else {
		return RSelect(arr[pivotIndex:], i-pivotIndex)
	}
}
