package twofer

import (
	"fmt"
	"strings"
)

func ShareWith(name string) string {
	name = strings.Trim(name, " \t\n\r")

	if name == "" {
		return "One for you, one for me."
	} else {
		return fmt.Sprintf("One for %v, one for me.", name)
	}
}
