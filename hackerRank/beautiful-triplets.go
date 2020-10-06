package hackerrank

// Naive (n*(n-1)*(n-2))... complexity implementation. Which is actually O(n) = n^3 (?)
func beautifulTripletsNaive(d int32, arr []int32) int32 {
	var c int32
	for i := 0; i<len(arr); i++ {
		for j := i+1; j<len(arr); j ++ {
			for k := j+1; k<len(arr); k++ {
				if arr[j] - arr[i]  == d  && arr[k] - arr[j] == d {
					c++
				}
			}
		}
	}
	return c
}

func beautifulTriplets(d int32, arr []int32) int32 {
	// first check with linear complexity to remove corner cases (TODO: Really not sure this is the right way?)
	var c int32
	for i := 0; i<len(arr); i++ {
		for j := i+1; j < len(arr) && arr[i] + d >= arr[j]; j++ {
			for k := j+1; k < len(arr) && arr[j] + d >= arr[k]; k++ {
				if arr[j] - arr[i]  == d  && arr[k] - arr[j] == d {
					c++
				}
			}
		}
	}
	return c
}
