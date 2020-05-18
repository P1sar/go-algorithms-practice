package backtracking

import "fmt"



func nQueensBacktracking(perm []int32, n int32) []int32{
	if int32(len(perm)) == n {
		fmt.Println(perm)
		return perm
	}

	for i:=int32(0); i<n; i++ {
		if !inArr(perm, i) {
			perm = append(perm, i)
			if canBeExtendedToSolution(perm) {
				fmt.Println(perm)
				nQueensBacktracking(perm, n)
			}
			perm = perm[:len(perm)-1]
		}
	}
	return perm
}

func canBeExtendedToSolution(perm []int32) bool {
	i := int32(len(perm) - 1)
	for j := int32(0); j<i; j++ {
		if i-j == abs(perm[i] - perm[j]) {
			return false
		}
	}
	return true
}

func abs(i int32) int32 {
	if i <0 {
		return -i
	}
	return i
}

func inArr(arr []int32, n int32) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}



//Branch and bound solution to N queens