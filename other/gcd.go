package other


// Euclidian algorithm gcd(a,b) = gcd(b,rem(a,b)


func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	 return gcd(b, a%b)
}
