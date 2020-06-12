package hackerrank

//- n: an integer, the number of rows and columns in the board
//- k: an integer, the number of obstacles on the board
//- r_q: integer, the row number of the queen's position
//- c_q: integer, the column number of the queen's position
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// queen with position q[r][c] attacks:
	// - its own row [r]
	// - its own column [c]
	// - 4 diagonals:
	//   - left down  c-1...n, r-1...n
	//   - left up    c-1...n, r+1...n
	//   - right down c+1...n, r-1...n
	//   - right up   c+1,     r+1...n

	var mostLeftObstacleCol, rightMostObstacleCol, upperMostObstacleRow, downMostObstacleRow int32
	var leftUpDiagonal, leftDownDiagonal, rightUpDiagonal, rightDownDiagonal int32 = -1,-1,-1,-1
	for _, o := range obstacles {
		// obstacle row same as queen
		if o[0] == r_q {
			// obstacle is to the right from the queen
			if o[1] > c_q {
				if o[1] < rightMostObstacleCol || rightMostObstacleCol == 0 {
					rightMostObstacleCol = o[1]-1
				}
			} else {
				if o[1] > mostLeftObstacleCol {
					mostLeftObstacleCol = o[1]
				}
			}
		} else if o[1] == c_q { // obstacle col same as queen col
			// obstacle upper then the queen
			if o[0] > r_q {
				if upperMostObstacleRow == 0 || o[0] < upperMostObstacleRow {
					upperMostObstacleRow = o[0] - 1
				}
			} else {
				// obstacle lower then the queen
				if o[0] > downMostObstacleRow {
					downMostObstacleRow = o[0]
				}
			}
		} else if abs(o[0] - r_q) == abs(o[1]-c_q) { // check whatever this obstacle is on the same diagonal
			if r_q - o[0] < 0 {
				// means diagonal obstacle row is bigger so it is upper
				if c_q - o[1] < 0 {
					var tmp int32
					// means diagonal obstacle col bigger, so diagonal obstacle to the right-up
					if (n-r_q) > (n-c_q) {
						tmp = o[1]- c_q - 1
					} else {
						tmp = o[0]- r_q - 1
					}
					if tmp < rightUpDiagonal || rightUpDiagonal == -1 {
						rightUpDiagonal = tmp
					}

				} else {
					// means diagonal obstacle col smaller, so diagonal obstacle ti the left-up
					//leftUpDiagonal
					var tmp int32
					if (n-r_q) > (c_q - 1) {
						tmp = c_q - 1 - o[1]
					} else {
						tmp = o[0]-r_q - 1
					}
					if tmp < leftUpDiagonal || leftUpDiagonal == -1 {
						leftUpDiagonal = tmp
					}
				}
			} else {
				// diagonal obstacle is lower
				if c_q - o[1] < 0 {
					var tmp int32
					// means diagonal obstacle col bigger, so diagonal obstacle to the right-down
					if (r_q-1) > (n-c_q) {
						tmp = o[1]- c_q - 1
					} else {
						tmp = r_q - o[0] - 1
					}
					if tmp < rightDownDiagonal || rightDownDiagonal == -1 {
						rightDownDiagonal = tmp
					}
				} else {
					// means diagonal obstacle col smaller, so diagonal obstacle ti the left-down
					var tmp int32
					if (r_q-1) > (c_q-1) {
						tmp = c_q - o[1] - 1 // limited by obstacle col
					} else {
						tmp = r_q-o[0]-1//limited by obstacle row
					}
					if tmp < leftDownDiagonal || leftDownDiagonal == -1{
						leftDownDiagonal = tmp
					}
				}
			}
		}
	}

	// row left could be limited if some obstacle (row_r, col_c)  has row_r on [1...r_q)
	rowLeftToAttack := c_q - 1 - mostLeftObstacleCol

	// row right could be limited if some obstacle (row_r, col_c) has row_r on (r_q...n]
	//rightMostO is table border or obstacle
	if rightMostObstacleCol == 0 {
		rightMostObstacleCol = n
	}
	rowRightToAttack := rightMostObstacleCol - c_q

	// col up could be limited if some obstacle (row_r, col_c) has col on [n...c_q)
	if upperMostObstacleRow == 0 {
		upperMostObstacleRow = n
	}
	colUpToAttack := upperMostObstacleRow-r_q
	// col down could be limited if some obstacle (row_r, col_c) has col on (c_q...1]

	colDownToAttack := r_q - 1 - downMostObstacleRow

	// row limits upper diagonal bound to [1...(n-row)] and down diagonal limit to [1...(row-1)]
	// columns limits left diagonal bounds to [1...(col-1)] and right diagonal bound [1...(n-col)]

	//leftUpDiagonal
	if leftUpDiagonal < 0 {
		if (n - r_q) > (c_q - 1) {
			leftUpDiagonal = c_q - 1
		} else {
			leftUpDiagonal = n - r_q
		}
	}
	//leftDownDiagonal
	if leftDownDiagonal < 0 {
		if (r_q - 1) > (c_q - 1) {
			leftDownDiagonal = c_q - 1
		} else {
			leftDownDiagonal = r_q - 1
		}
	}
	//rightUpDiagonal
	if rightUpDiagonal < 0 {
		if (n - r_q) > (n - c_q) {
			rightUpDiagonal = n - c_q
		} else {
			rightUpDiagonal = n - r_q
		}
	}
	//rightDownDiagonal
	if rightDownDiagonal < 0 {
		if (r_q - 1) > (n - c_q) {
			rightDownDiagonal = n - c_q
		} else {
			rightDownDiagonal = r_q - 1
		}
	}
	p := rowLeftToAttack+rowRightToAttack+colUpToAttack+colDownToAttack+leftUpDiagonal+leftDownDiagonal+rightUpDiagonal+rightDownDiagonal
	//fmt.Println(rowLeftToAttack, rowRightToAttack, colUpToAttack, colDownToAttack, leftUpDiagonal, leftDownDiagonal, rightUpDiagonal, rightDownDiagonal)
	return p
}

func abs(i int32) int32 {
	if i < 0 {
		return -i
	}
	return i
}