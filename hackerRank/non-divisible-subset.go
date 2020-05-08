package hackerrank

import "fmt"

//{278,576,496,727,410,124,338,149,209,702,282,718,771,575,436}
func NonDivisibleSubset(k int32, arr []int32) int32 {
	remainders := make(map[int32]int32)
	for _, v := range arr {
		r := v%k
		_, ok := remainders[r]
		if ok {
			remainders[r] += 1
		} else {
			remainders[r] = 1
		}
	}
	var max int32 = 0
	for k1, _ := range remainders {
		items := make([]int32, 0)
		for k2, _ := range remainders {
			items = append(items, k1)
			if k1 == k2 {
				continue
			}
			d := false
			for _, v := range items {
				if (k2 + v)%k == 0 {
					d = true
				}
			}
			if !d {
				items = append(items, k2)
			}
		}
		var tmp int32 = 0
		fmt.Println(items)
		for _, divisible := range items {
			tmp += remainders[divisible]
		}
		if tmp > max {
			max = tmp
		}
	}
	fmt.Println(remainders)
	return max
}



//t := make([]int32, len(arr))
//for i := range t {
//	t[i] = 1
//}
//
//max := int32(0)
//for j:=1; j<len(arr); j++ {
//	tmp := int32(0)
//	for i:=0; i<j; i++ {
//		if (arr[i] + arr[j])%k != 0 {
//			fmt.Printf("%v + %v NOT divisible by %v\n", arr[i], arr[j], k)
//			tmp += 1
//		} else {
//			fmt.Printf("%v + %v divisible by %v\n", arr[i], arr[j], k)
//			continue
//		}
//	}
//	if tmp > t[j] {
//		t[j] = tmp
//		max = tmp
//	}
//}
//fmt.Println(t)
//return max