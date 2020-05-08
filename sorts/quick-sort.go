package sorts

func QuickSort(arr []int64, tc int64) ([]int64, int64) {
	if len(arr) <= 1 {
		return arr, tc
	}
	pivotIndex := selectPivot(arr)

	tc += int64(len(arr))-1
	i := rearrangeAroundPivot(arr, pivotIndex)
	_, tc1 := QuickSort(arr[:i], tc)
	_, tc2 := QuickSort(arr[i+1:], tc1)

	return arr, tc2
}


func selectPivot(arr []int64) int64 {
	//var pivots = make([]int64,3)
	//if len(arr)%2!=0 {
	//	pivots[1] = arr[int64(len(arr))/2]
	//} else {
	//	pivots[1] = arr[int64(len(arr))/2 - 1]
	//}
	//pivots[0] = arr[0]
	//pivots[2] = arr[int64(len(arr))-1]
	//pivotsSorted := MergeSort(pivots[:])
	////fmt.Println(arr)
	////fmt.Println(pivots)
	////fmt.Println(pivotsSorted)
	//for i, v := range arr {
	//	if v == pivotsSorted[1] {
	//		//fmt.Println(i)
	//		return int64(i)
	//	}
	//}

	return int64(len(arr))-1
	//return 0
}

func rearrangeAroundPivot(arr []int64, pI int64) int64 {
	arr[0], arr[pI] = arr[pI], arr[0]
	pI = 0
	p := arr[pI]
	i := pI+1
	for j := pI+1; j < int64(len(arr)); j++ {
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i-1], arr[pI] = arr[pI], arr[i-1]
	return i-1
}
