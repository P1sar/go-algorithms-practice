package hamming

import "errors"


func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1,errors.New("wrong length")
	}
	dif := 0
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			dif += 1
		}
	}
	return dif, nil
}
