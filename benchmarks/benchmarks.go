package main

import "testing"

var code = []int32{0, 10, 100, 100, 0, 10, 0, 10, 100, 14, 1000, 100, 1000, 0, 0, 10, 100, 1000, 10, 0, 1000, 12}
var mapCode = map[int32]int32{
	0:    1,
	10:   2,
	100:  3,
	1000: 4,
}

var sliceCode = []int32{
	0:    1,
	10:   2,
	100:  3,
	1000: 4,
}

func BenchmarkSlice(b *testing.B) {
	success := int32(0)
	fail := int32(0)
	for n := 0; n < b.N; n++ {
		// for each value in code array, do a specific action
		for _, v := range code {
			c := sliceCode[v]
			if c == 0 {
				fail++
			} else {
				success += c
			}
		}
	}
}

func BenchmarkMap(b *testing.B) {
	success := int32(0)
	fail := int32(0)
	for n := 0; n < b.N; n++ {
		// for each value in code array, do a specific action
		for _, v := range code {
			c, ok := mapCode[v]
			if !ok {
				fail++
			} else {
				success += c
			}
		}
	}
}

func BenchmarkSwitch(b *testing.B) {
	success := int32(0)
	fail := int32(0)
	for n := 0; n < b.N; n++ {
		// for each value in code array, do a specific action
		for _, v := range code {
			switch v {
			case 0:
				success++
			case 10:
				success += 2
			case 100:
				success += 3
			case 1000:
				success += 4
			default:
				fail++
			}
		}
	}
}

