package bob

import (
	"strings"
)

func Hey(remark string) string {
	var yell, ask = false, false

	remark = strings.Trim(remark, " \t\n\r")

	if remark == "" {
		return "Fine. Be that way!"
	}

	if strings.ToUpper(remark) == remark  && strings.ToLower(remark) != remark {
		yell = true
	}

	if remark[len(remark)-1:] == "?" {
		ask = true
	}

	if yell && ask {
		return "Calm down, I know what I'm doing!"
	} else if yell {
		return "Whoa, chill out!"
	} else if ask {
		return "Sure."
	} else {
		return "Whatever."
	}

	return ""
}
