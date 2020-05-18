package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	r, _ := regexp.Compile("[a-zA-Z]+")
	strs := r.FindAllString(s, -1)
	acronym := ""
	for _, str := range strs {
		elm := str[:1]
		acronym = acronym + strings.ToUpper(elm)
	}
	return acronym
}
