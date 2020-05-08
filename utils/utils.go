package utils


// TODO rewrite with 'generics'
func RemoveElementFromArr(slice []string, e string) {
	for i, v := range slice  {
		if v == e {
			// Moving last arr element on position of found element
			slice[i] = slice[len(slice)-1]
			// Remove last element
			slice = slice[:len(slice)-1]
		}
	}
}

