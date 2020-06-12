package hackerrank


func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	naiveTotal := int64(b)*int64(bc) + int64(w)*int64(wc)
	if int64(wc)+int64(z) < int64(bc) {
		// cheaper to buy white and change to black
		return int64(wc)*int64((b+w)) + int64(z)*int64(b)
	}
	if int64(bc)+int64(z) < int64(wc) {
		//cheaper to buy black and change to white
		return int64(bc)*int64((b+w)) + int64(z)*int64(w)
	}
	return naiveTotal
}
