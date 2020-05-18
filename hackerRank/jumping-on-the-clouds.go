package hackerrank

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	fmt.Println(c)
	var j int32 = 0
	for i:=0; i<len(c)-1; i++ {
		fmt.Println(i)
		if (len(c)-1)>=i+2 && c[i+2] == 0 {
			i+=1
			j += 1
			continue
		} else {
			j+=1
			continue
		}
	}
	return j
}