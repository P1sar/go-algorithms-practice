package hackerrank

import (
	"fmt"
)

func acmTeam(topic []string) []int32 {
	var combinations = make(map[int32]int32)
	var max int32
	for i:= 0; i < len(topic); i++ {
		for j := i+1; j< len(topic); j++ {
			knowledge := symbolWiseOR(topic[i], topic[j])
			if knowledge >= max {
				max = knowledge
				if _, ok := combinations[knowledge]; ok {
					combinations[knowledge] += 1
				} else {
					combinations[knowledge] = 1
				}
			}
		}
	}
	fmt.Println(combinations)
	fmt.Println(max)
	return []int32{max, combinations[max]}
}

func symbolWiseOR(a, b string) int32 {
	var c int32
	for i := 0; i <len(a); i ++ {
		if string(a[i]) == "1" || string(b[i]) == "1" {
			c++
		}
	}
	return c
}

