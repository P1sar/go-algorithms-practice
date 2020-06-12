package hackerrank

// Complete the organizingContainers function below.
func organizingContainers(container [][]int32) string {
	//idea is that number of balls in container is intact since its +1/-1 swaps
	totalBallsInBasket := make(map[int32]int32)
	totalBallTypeAmounts := make(map[int32]int32)
	for boxNumber:=int32(0); boxNumber<int32(len(container)); boxNumber++ {
		for ballType:=int32(0);ballType<int32(len(container[boxNumber]));ballType++ {
				ballTypeAmount := container[boxNumber][ballType]
				totalBallTypeAmounts[ballType] += ballTypeAmount
				totalBallsInBasket[boxNumber] += container[boxNumber][ballType]
		}
	}
	for _, ballsInBasket := range totalBallsInBasket {
		var tmp int32 = -1
		for ballType, ballTypeAmount := range totalBallTypeAmounts {
			if ballsInBasket == ballTypeAmount {
				tmp = ballType
			}
		}
		if tmp == -1 {
			return "Impossible"
		}
		delete(totalBallTypeAmounts, tmp)
		tmp = -1
	}
	return "Possible"
}


