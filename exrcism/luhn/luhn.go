package luhn

import (
	"regexp"
	"strconv"
	"unicode"
)

func Valid(input string) bool {
	r, _ := regexp.Compile("[0-9]{2,}")

	match := r.MatchString(input)


	if !match {
		return false
	}

	odd_input := len(input) % 2 == 0

	var summed_input int32

	internal_iterator := 0

	for _, d := range input {

		_, ok := strconv.ParseInt(string(d), 10, 32)

		if ok != nil {
			if unicode.IsSpace(d) {
				continue
			} else {
				return false
			}
		}
		d = d - 48

		if odd_input {
			if internal_iterator%2 == 0 {
				d = double_and_substract(d)
			}
		} else {
			if internal_iterator%2 != 0 {
				d = double_and_substract(d)
			}
		}

		summed_input += d
		internal_iterator ++
	}


	if summed_input % 10 == 0 {
		return true
	} else {
		return false
	}
}

func double_and_substract(d int32) int32 {
	d = d * 2
	if d > 9 {
		d = d - 9
	}

	return d
}