package hackerrank

import "fmt"

func timeInWords(h int32, m int32) string {
	var w string
	keyWords := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"quarter",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
		"twenty",
		"twenty one",
		"twenty two",
		"twenty three",
		"twenty four",
		"twenty five",
		"twenty six",
		"twenty seven",
		"twenty eight",
		"twenty nine",
		"half",
	}
	if m == 0 {
		return fmt.Sprintf("%s o' clock", keyWords[h])
	}
	if m > 30 {
		m = 60 - m
		w = "to"
		h = (h+1)%12
		if h == 0 {
			h = 12
		}

	} else {
		w = "past"
	}
	o := " minutes"
	if m==1 {
		o = " minute"
	} else if  m%15 == 0 {
		o = ""

	}
	mins := keyWords[m]
	hours := keyWords[h]

	return fmt.Sprintf("%s%s %s %s", mins,o,w,hours)
}