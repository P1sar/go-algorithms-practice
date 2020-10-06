package hackerrank

func minimumDistances(a []int32) int32 {
	var c int32 = -1
	for i := 0; i < len(a); i ++ {
		for j := i+1; j < len(a); j ++ {
			if a[i] == a[j] {
				tmp := int32(j - i)
				if tmp < c || c == -1 {
					c = tmp
				}
			}
		}
	}
	return c
}