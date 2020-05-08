package sorts

//[]int{4,7,23,1,23,77,3,1,77,3,2,6,7}
func MyInsertionSort(a []int) []int {
	for i := 1; i<len(a); i++ {
		j := i
		for j > 0 {
			if a[j] > a[j-1] {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j--
		}
	}
	return a
}

func SomeInsertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}