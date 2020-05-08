package hackerrank

import (
	"bufio"
	"os"
	"strconv"
)

//
//var templates = [][][]int32{
//	{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
//	{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
//	{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
//	{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
//	{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
//	{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
//	{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
//	{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
//}
//
//func main() {
//	m := [][]int32{{4, 5, 8}, {2, 4, 1}, {1, 9, 7}}
//	pointsArr := make([]int32,0)
//	for _, v := range templates {
//		points := convertWithLowestPoints(m, v)
//		pointsArr = append(pointsArr, points)
//	}
//	low := pointsArr[0]
//	for _, p := range pointsArr {
//		if p < low {
//			low = p
//		}
//	}
//	fmt.Print(low)
//}
//
//func convertWithLowestPoints(from [][]int32, to [][]int32) int32 {
//	var points int32
//	for row:=0; row<3; row++ {
//		for i:=0; i<3; i++{
//			points += countPoints(from[row][i], to[row][i])
//		}
//
//	}
//	return points
//}
//
//func countPoints(from int32, to int32) int32 {
//	return Abs(from-to)
//}
//
//func Abs(x int32) int32 {
//	if x < 0 {
//		return -x
//	}
//	return x
//}

//
//func formingMagicSquare(s [][]int32) int32 {
//	variants := [][]int32{{8,6,4,2}, {9,7,3,1}, {8,6,4,2}, {9,7,3,1}, {5}, {9,7,3,1}, {8,6,4,2}, {9,7,3,1}, {8,6,4,2}}
//	//magicConst := int32(15)
//	//magicSum := int32(45)
//	c := []int32{1,2,3,4,5,6,7,8,9}
//	var points int32
//	variantsIndex := 0
//	for _, v := range s {
//		for _, i := range v {
//			allowed := variants[variantsIndex]
//			points += findPointsToClosetAllowedInt(i, &c, allowed)
//			fmt.Printf("Transforming %v with points %v \n", v, points)
//			variantsIndex += 1
//		}
//	}
//	return points
//}
//
//func Abs(x int32) int32 {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//
//func findPointsToClosetAllowedInt(x int32, arr *[]int32, allowed []int32) int32 {
//	var points int32 = 20
//	var index int32
//	for i, v := range *arr {
//		if v != int32(0) && in(v, allowed){
//			if points > Abs(x-v) {
//				index = int32(i)
//				points = Abs(x-v)
//			}
//		}
//	}
//	fmt.Printf("Transformed %v to %v\n", x, (*arr)[index])
//	(*arr)[index] = int32(0)
//	return points
//}
//
//func in(x int32, arr []int32) bool {
//	for _, v := range arr {
//		if x == v {
//			return true
//		}
//	}
//	return false
//}
//
//func arrSum(arr []int32) int32 {
//	var tmp int32
//	for _, v := range arr {
//		tmp += v
//	}
//	return tmp
//}
//
//func main() {
//	res := pickingNumbers([]int32{4,6,5,3,3,1})
//	fmt.Println(res)
//}
//
//func pickingNumbers(a []int32) int32 {
//	var biggestLen int32
//	for _, base := range a {
//		baseUp := base + 1
//		baseDown := base - 1
//		var pointsUp int32
//		var pointsDown int32
//		for _, v := range a {
//			if v == baseUp {
//				pointsUp +=1
//			} else if v == baseDown {
//				pointsDown += 1
//			} else if v == base {
//				pointsUp += 1
//				pointsDown += 1
//			}
//			fmt.Printf("PointsUp %v, PointsDown: %v, base: %v, v: %v \n", pointsUp, pointsDown, base, v)
//		}
//		if pointsUp >= pointsDown && pointsUp >= biggestLen {
//			biggestLen = pointsUp
//		} else if pointsDown >= pointsUp && pointsDown >= biggestLen {
//			biggestLen = pointsDown
//		}
//	}
//	return biggestLen
//}
//func pickingNumbers(a []int32) int32 {
//	var biggestLen int32
//	for _, base := range a {
//		var points int32 = 0
//		for _, v := range a {
//			if Abs(base-v)<=1 {
//				fmt.Printf("Points prev %v, base: %v, v: %v \n", points, base, v)
//				points += 1
//			}
//		}
//		if points > biggestLen {
//			biggestLen = points
//		}
//	}
//	return biggestLen
//}
// Complete the climbingLeaderboard function below.
//func climbingLeaderboard(scores []int32, alice []int32) []int32 {
//ranks := make([]int32, 0)
//var r int32 = 1
//var pS int32
//for _, aS := range alice {
//	pS = 0
//	r = 1
//	for i, tS := range scores {
//		if aS >= tS {
//			ranks = append(ranks, r)
//			pS = 0
//			r = 1
//			break
//		}
//		if pS != tS {
//			r +=1
//			pS = tS
//		}
//		if i == len(scores)-1 {
//			ranks = append(ranks, r)
//			pS = 0
//			r = 1
//		}
//	}
//}
//return ranks
//}


func readLines(path string) ([]int32, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int32
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		lines = append(lines, int32(i))
	}
	return lines, scanner.Err()
}

//"200""
//"5 5 6 14 19 20 23 25 29 29 30 30 32 37 38 38 38 41 41 44 45 45 47 59 59 62 63 65 67 69 70 72 72 76 79 82 83 90 91 92 93 98 98 100 100 102 103 105 106 107 109 112 115 118 118 121 122 122 123 125 125 125 127 128 131 131 133 134 139 140 141 143 144 144 144 144 147 150 152 155 156 160 164 164 165 165 166 168 169 170 171 172 173 174 174 180 184 187 187 188 194 197 197 197 198 201 202 202 207 208 211 212 212 214 217 219 219 220 220 223 225 227 228 229 229 233 235 235 236 242 242 245 246 252 253 253 257 257 260 261 266 266 268 269 271 271 275 276 281 282 283 284 285 287 289 289 295 296 298 300 300 301 304 306 308 309 310 316 318 318 324 326 329 329 329 330 330 332 337 337 341 341 349 351 351 354 356 357 366 369 377 379 380 382 391 391 394 396 396 400"
//func main() {
//	//res := climbingLeaderboard([]int32{100, 99, 88, 77, 66, 55, 44, 40, 10, 9,8,7,6,5,4,3,2,1}, []int32{25})
//	//res := climbingLeaderboard([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
//	//fmt.Println(res)
//	//6,4,2,1
//	dar, _ := readLines("scores")
//	alice, _ := readLines("alice")
//	fmt.Printf("%v", dar)
//	//res := climbingLeaderboard(dar, []int32{5,5,6,14,19,20,23,25,29,29,30,30,32,37,38,38,38,41,41,44,45,45})
//	res := climbingLeaderboard(dar, alice)
//	fmt.Printf("%v", res)
//
//
//}
//
//func timeTrack(start time.Time, name string) {
//	elapsed := time.Since(start)
//	log.Printf("%s took %s", name, elapsed)
//}
//
//func main() {
//	defer timeTrack(time.Now(), "climbing")
//	f, _ := os.Open("input07.txt")
//	reader := bufio.NewReaderSize(f, 8024 * 1024)
//
//	stdout, err := os.Create("output")
//	checkError(err)
//
//	defer stdout.Close()
//
//	writer := bufio.NewWriterSize(stdout, 8024 * 1024)
//
//	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//
//	scoresTemp := strings.Split(readLine(reader), " ")
//
//	var scores []int32
//
//	for i := 0; i < int(scoresCount); i++ {
//		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
//		checkError(err)
//		scoresItem := int32(scoresItemTemp)
//		scores = append(scores, scoresItem)
//	}
//
//	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//
//	aliceTemp := strings.Split(readLine(reader), " ")
//
//	var alice []int32
//
//	for i := 0; i < int(aliceCount); i++ {
//		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
//		checkError(err)
//		aliceItem := int32(aliceItemTemp)
//		alice = append(alice, aliceItem)
//	}
//
//	result := climbingLeaderboard(scores, alice)
//
//	for i, resultItem := range result {
//		fmt.Fprintf(writer, "%d", resultItem)
//
//		if i != len(result) - 1 {
//			fmt.Fprintf(writer, "\n")
//		}
//	}
//
//	fmt.Fprintf(writer, "\n")
//
//	writer.Flush()
//}
//
//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//
//
//func climbingLeaderboard(scores []int32, alice []int32) []int32 {
//	// Remove duplicates O(n)
//	var prev int32
//	noDupScores := make([]int32,0)
//	for _, v := range scores {
//		if prev != v {
//			prev = v
//			noDupScores = append(noDupScores, v)
//		}
//		continue
//	}
//	scores = noDupScores
//	lenScores := len(scores)
//	var prevAlice int32 = -1
//	var prevRate int32 = -1
//	ss := make([]int32, len(alice))
//	var median int32
//	if lenScores % 2 == 0 {
//		median = int32(lenScores/2)
//	} else {
//		median = int32((lenScores-1)/2)
//	}
//	for i, v := range alice {
//		var rate int32
//		if v == prevAlice {
//			rate = prevRate
//		} else {
//			rate = searchIndexBinaryWay(&scores, median, v)
//		}
//		prevAlice = v
//		prevRate = rate
//		if prevRate >= int32(lenScores) {
//			median = prevRate - 1
//		} else {
//			median = prevRate
//		}
//		ss[i]= rate+1
//		continue
//	}
//	return ss
//}
//
//func searchIndexBinaryWay(scores *[]int32, median int32, alice int32) int32 {
//	index := median
//	lenScores := len(*scores)
//	score := (*scores)[index]
//	prevLowerIndex := int32(len(*(scores))-1)
//	prevBiggerIndex := int32(0)
//	for {
//		//fmt.Printf("Index: %v, PrevLowerIndex: %v PrevBiggerIndex: %v Score: %v ||", index, prevLowerIndex, prevBiggerIndex, score)
//		switch {
//		case score < alice:
//			fmt.Printf("%v < %v \n", score, alice)
//			if index-1 < 0 {
//				return 0
//			}
//			prevScore := (*scores)[index-1]
//			if prevScore > alice {
//				//fmt.Printf("prev %v > %v alice \n", prevScore, alice)
//				return index
//			}
//			index, prevLowerIndex = defineNewIndex(prevBiggerIndex, index), index
//			score = (*scores)[index]
//			continue
//		case score > alice:
//			//fmt.Printf("%v > %v \n", score, alice)
//			if index+1 > int32(lenScores-1) {
//				return int32(lenScores)
//			}
//			nextScore := (*scores)[index+1]
//			if nextScore < alice {
//				//fmt.Printf("Next %v < %v alice with index %v \n", nextScore, alice, index)
//				return index+1
//			}
//			if prevBiggerIndex < index {
//				prevBiggerIndex = int32(len(*scores) - 1)
//			}
//			index,  prevBiggerIndex = defineNewIndex(index+1, prevLowerIndex), index
//			score = (*scores)[index]
//			continue
//		case score == alice:
//			//fmt.Printf("%v %v == %v \n", index, score, alice)
//			return index
//		}
//	}
//}
//
//func defineNewIndex(from, to int32) int32 {
//	if from - to == 1 {
//		return from
//	}
//	total := from + to
//	if total % 2 == 0 {
//		return int32(total/2)
//	} else if total == 1 {
//		return int32(0)
//	} else {
//		return int32((total+1)/2)
//	}
//}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	lenA := int32(len(a))
	k = k%lenA
	for i:=int32(0); i<k; i++ {
		var tmp int32
		for j, v := range a {
			if j == 0 {
				a[j], tmp = a[lenA-1], v
			} else {
				a[j], tmp = tmp, v
			}
		}
	}
	res := make([]int32, len(queries))
	for i, v := range queries {
		res[i] = a[v]
	}
	return res
}

func circularArrayRotationS(a []int32, k int32, queries []int32) []int32 {
	lenA := int32(len(a))
	k = k%lenA
	res := make([]int32, len(queries))
	for i, v := range queries {
		res[i] = a[moveIndexOffset(v, k, lenA)]
	}
	return res
}

func moveIndexOffset(index, moveTimes, lenA int32) int32{
	newPos := index - moveTimes
	if newPos < 0 {
		newPos = lenA - Abs(newPos)
	}
	return newPos%lenA
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

//func main() {
//	res := circularArrayRotationS([]int32{1,2,3}, 2, []int32{0,1,2})
//	fmt.Println(res)
//}


//
//for _, i := range arr {
//	if i == "0" {
//		continue
//	}
//	ii, _ := strconv.Atoi(i)
//	if n%int32(ii) == 0 {
//		c += 1
//	}
//}
