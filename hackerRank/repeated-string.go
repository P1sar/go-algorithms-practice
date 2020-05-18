package hackerrank

func repeatedString(s string, n int64) int64 {
	var a int64 = 97
	var count = countLetterInPattern(s, a)

	integerPatterns := n/int64(len(s))

	reminderOfPattern := n%int64(len(s))

	leftInPatternReminder := countLetterInPattern(s[:reminderOfPattern], a)

	count = count * integerPatterns + int64(leftInPatternReminder)


	return count
}

func countLetterInPattern(s string, l int64) int64 {
	var count int64 = 0
	for _, v := range s {
		if int64(v) == l {
			count += 1
		}
	}
	return count
}
