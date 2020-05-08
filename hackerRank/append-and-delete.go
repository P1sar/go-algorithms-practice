package hackerrank

import "fmt"

// s to t, exactly k moves
func AppendAndDelete(s string, t string, k int32) string {
	hd := int32(hammingDistance(s, t))
	if len(s) > len(t) {
		// means we should firstly remove letter with amount == HD, so required steps is
		steps := int(hd)
		cuttedSLen := len(s) - int(hd)
		// how may we should append to s to receive k
		steps += len(t)-cuttedSLen
		fmt.Println(steps)
		if steps < int(k) {
			return "Yes"
		}
		if (steps-int(k)) % 2 == 0 {
			return "Yes"
		} else {
			return "No"
		}
	}
	if len(t) > len(s) {
		// means we should only add some letters to s
		if len(t) - int(hd) >= len(s) {
			if hd == k {
				return "Yes"
			}  else if hd > k {
				return "No"
			} else  { // hd < k
				if (k-hd)%2==0 {
					return "Yes"
				} else {
					// We actually can remove from s some first to make it fit in K, but let it be so... (:
					return "No"
				}
			}

		}
		// means we should remove some letters from s first
		if len(t) - int(hd) < len(s) {
			toRemove := len(s) - len(t) - int(hd)
			if toRemove*2 > int(k) {
				return "No"
			} else {
				return "Yes"
			}
		}
	}
	if len(t) == len(s) {
		// checking that strings are same
		if hd == 0 {
			// if k is even, we can perform even numbers in final leaving s same
			if k%2 == 0 {
				return "Yes"
			} else {
				// we can remove full string length and add it later, since removing from empty string reduces k but do not affect string
				if int(k)-len(t)*2 > 0 {
					return "Yes"
				}
				return "No"
			}
		} else if int(hd)==len(t) {
			if hd*2 <= k {
				return "Yes"
			} else {
				return "No"
			}
		} else {
			steps := hd * 2
			if steps > k {
				return "No"
			} else {
				if (k - steps)%2==0 {
					return "Yes"
				}else {
					return "No"
				}
			}
		}
	}
	return "NoTheFuck"
}

func hammingDistance(p, q string) int {
	distance := 0
	var longest, shortest string
	if len(p) > len(q) {
		longest = p
		shortest = q
	} else {
		longest = q
		shortest = p
	}
	for i:=0; i<len(longest); i++ {
		if i >= len(shortest) {
			distance = len(longest) - len(shortest)
			break
		}
		if longest[i] != shortest[i] {
			fmt.Println(i)
			distance = len(longest) - len(shortest) + len(longest) - i
			break
		}
	}
	fmt.Printf("HD %v \n", distance)
	return distance
}


//appendAndDelete("ashley", "ash", 2) // No