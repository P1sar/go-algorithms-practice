package hackerrank

func nonDivisibleSubset(k int32, arr []int32) int32 {
	remainderArr := make([]int32,len(arr))
	remaindersMap := make(map[int32]int32)
	for i, v := range arr {
		r := v%k
		remainderArr[i] = r
		if _, ok := remaindersMap[r]; ok {
			remaindersMap[r] += 1
		} else {
			remaindersMap[r] = 1
		}
	}

	remindersScoresMap := make(map[int32]int32)
	for i:=int32(0); i<=k-1; i++ {
		if _, ok := remaindersMap[i]; ok {
			if i == 0 {
				if _, ok := remindersScoresMap[i]; ok {
					continue
				} else {
					remindersScoresMap[i] = 1
					continue
				}
			} else {
				// edge case is when reminder == opposit, then we should use it only once
				iOpponent := k - i
				if iOpponent == i {
					remindersScoresMap[i] = 1
					continue
				}
				iScore, ok := remaindersMap[i]
				if !ok {
					iScore = 0
				}
				kScore, ok := remaindersMap[iOpponent]
				if !ok {
					kScore = 0
				}
				if iScore > kScore {
					remindersScoresMap[i] = iScore
				} else {
					remindersScoresMap[iOpponent] = kScore
				}
			}
		}
	}
	var score int32 = 0
	for _, v := range remindersScoresMap {
		score += v
	}
	return score
}