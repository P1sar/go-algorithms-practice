package hackerrank

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	left := s
	var i int32 =  0
	for left >= m {
		if left < p {
			return i
		}
		left = left-p
		if p > m {
			p = p - d
		}
		if p < m {
			p = m
		}
		i++
	}
	return i
}
