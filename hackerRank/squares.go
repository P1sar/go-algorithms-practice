package hackerrank

import "math"

func squares(a int32, b int32) int32 { //17 24 | 0
	var c int32 = 0
	var firstFactor int32
	for i:=a; i<=b; i++ {
		tmp := math.Sqrt(float64(i))
		tmpI := int32(tmp)
		if tmpI*tmpI == i {
			c=1
			firstFactor = tmpI
			break
		}
	}
	if firstFactor == 0 {
		return 0
	}
	for i:= firstFactor+1; (i*i)<=b; i++ {
		c++
	}

	return c
}
