package devideAndConquer

import (
	"math"
)

func MaxSubArray(arr []int) ([]int, int) {
	if len(arr) == 1 {
		return arr, arr[0]
	}

	maxLeftSubArray, leftSum := MaxSubArray(arr[:len(arr)/2])
	maxRightSubArray, rightSum := MaxSubArray(arr[len(arr)/2:])
	maxCrossSubArr, crossSum := maxCrossArr(arr)

	//fmt.Println(leftSum, rightSum, crossSum)

	if leftSum >= rightSum && leftSum >= crossSum {
		return maxLeftSubArray, leftSum
	} else if rightSum > leftSum && rightSum >= crossSum {
		return maxRightSubArray, rightSum
	} else {
		return maxCrossSubArr, crossSum
	}
}

func maxCrossArr(arr []int) ([]int, int) {
	if len(arr) == 2 {
		return arr, arr[0] + arr[1]
	}
	sumLeft := int64(math.MinInt64)
	sum := int64(0)
	maxLeft := 0
	for i := len(arr)/2; i>= 0; i-- {
		sum += int64(arr[i])
		if sum > sumLeft {
			sumLeft = sum
			maxLeft = i
		}
	}

	sumRight := int64(math.MinInt64)
	sum = int64(0)
	maxRight := 0
	for j := len(arr)/2+1; j < len(arr); j++ {
		sum += int64(arr[j])
		if sum > sumRight {
			maxRight = j
			sumRight = sum
		}
	}
	return arr[maxLeft:maxRight+1], int(sumLeft+sumRight)
}

//a= []int{13,-3,-25,20,-3,-16,-23,18,20,-7,12,-5,-22,15,-4,7}