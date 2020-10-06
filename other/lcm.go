package other

// finding least common multiple using greatest common divisor of that numbers.
func lcm(a, b int64) int64 {
	return a*b/gcd(a,b)
}

