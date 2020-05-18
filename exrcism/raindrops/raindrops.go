package raindrops

import "strconv"

func Convert(number int) string {

	raindropSpeak := ""


	if number % 3 == 0 {
		raindropSpeak += "Pling"
	}
	if number % 5 == 0  {
		raindropSpeak += "Plang"
	}
	if number % 7 == 0  {
		raindropSpeak += "Plong"
	}
	if raindropSpeak == "" {
		return strconv.Itoa(number)
	} else {
		return raindropSpeak
	}
}