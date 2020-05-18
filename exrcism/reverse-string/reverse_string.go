package reverse

func String(word string) string {

	rune_array := []rune(word)

	var reversed_word string

	for i := len(rune_array) - 1; i >= 0; i -- {
		reversed_word += string(rune_array[i])
	}
	return reversed_word
}