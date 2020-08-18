package hackerrank

import (
	"fmt"
	"strconv"
)

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p int32, q int32){
	var c bool
	for i := p; i<=q; i++ {
		strI := strconv.Itoa(int(i))
		d := len(strI)
		sqr := int64(i)*int64(i)
		lensqr := len(strconv.Itoa(int(sqr)))
		fmt.Println(sqr)
		var m1,m2 string
		if lensqr == d*2 {
			m1,m2 = strconv.Itoa(int(sqr))[:d], strconv.Itoa(int(sqr))[d:]
		} else if lensqr == d*2 - 1 {
			m1,m2 = strconv.Itoa(int(sqr))[:d-1], strconv.Itoa(int(sqr))[d-1:]
		}
		fmt.Println(m1,m2)
		if lensqr == d*2 || lensqr == d*2 - 1 {
			m1Int, _ := strconv.ParseInt(m1, 10, 64)
			m2Int, _ := strconv.ParseInt(m2, 10, 64)
			if m2Int == 0 {
				continue
			}
			if m1Int + m2Int == int64(i) {
				fmt.Printf("%v ",i)
				c = true
			}
		}
	}
	if !c {
		fmt.Printf("%v","INVALID RANGE")
	}
}