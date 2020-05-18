package isogram

import (
	"fmt"
	"strings"
)

func  IsIsogram(words string) bool  {

	verbs_count := make(map[int32]int)

	r := strings.NewReplacer("-", "", " ", "")

	words = r.Replace(words)
	fmt.Println(words)

	for _, v := range strings.ToUpper(words) {
		_, exists := verbs_count[v]
		if exists {
			return false
		} else {
			verbs_count[v] = 1
		}
	}
	return true
}