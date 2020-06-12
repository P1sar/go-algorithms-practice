package hackerrank

import (
	"fmt"
	"math"
	"strings"
)

func encryption(s string) string {
	L := len(s)
	sqrt := math.Sqrt(float64(L))
	rows := int32(math.Floor(sqrt))
	columns := int32(math.Ceil(sqrt))
	if !(rows*columns >= int32(L)) {
		rows += 1
	}
	str := ""
	fmt.Println(L, rows, columns)
	for j := int32(0);j<int32(math.Ceil(sqrt)); j++ {
		for i := int32(j); i < rows*columns; i += columns {
			fmt.Println(j,i)
			if i >= int32(L) {
				break
			}
			str += string(s[i])
		}
		str += " "
	}
	return strings.Trim(str, " ")
}

