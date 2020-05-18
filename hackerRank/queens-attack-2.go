package hackerrank

import "fmt"


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

	var leftMostO, rightMostO, upperMostO, downMostO int32
	for _, o := range obstacles {
		
		// obstacle row same as queen
		if o[0] == r_q {
			// obstacle is to the right from queen
			if o[1] > c_q {

			} else {
				// obstacle is to the left from queen
			}
		}
		// obstacle col same as queen
		if o[1] == c_q {
			// obstacle upper then queen
			if o[0] > r_q {

			} else {
			// obstacle lower then queen
			}
		}
	}

	// row left could be limited if some obstacle (row_r, col_c)  has row_r on [1...r_q)
	rowLeftToAttack := c_q - 1

	// row right could be limited if some obstacle (row_r, col_c) has row_r on (r_q...n]
	rowRightToAttack := n-c_q

	// col up could be limited if some obstacle (row_r, col_c) has col on [n...c_q)
	colUpToAttack := n-r_q
	// col down could be limited if some obstacle (row_r, col_c) has col on (c_q...1]
	colDownToAttack := r_q-1
	fmt.Println(rowLeftToAttack, rowRightToAttack, colUpToAttack, colDownToAttack)

	// row limits upper diagonal bound to [1...(n-row)] and down diagonal limit to [1...(row-1)]
	// columns limits left diagonal bounds to [1...(col-1)] and right diagonal bound [1...(n-col)]
	var leftUpDiagonal, leftDownDiagonal, rightUpDiagonal, rightDownDiagonal int32

	//leftUpDiagonal
	if (n-r_q) > (c_q - 1) {
		leftUpDiagonal = c_q - 1
	} else {
		leftUpDiagonal = n-r_q
	}
	//leftDownDiagonal
	if (r_q-1) > (c_q-1) {
		leftDownDiagonal = c_q-1
	} else {
		leftDownDiagonal = r_q-1
	}
	//rightUpDiagonal
	if (n-r_q) > (n-c_q) {
		rightUpDiagonal = n-c_q
	} else {
		rightUpDiagonal = n-r_q
	}
	//rightDownDiagonal
	if (r_q-1) > (n-c_q) {
		rightDownDiagonal = n-c_q
	} else {
		rightDownDiagonal = r_q-1
	}
	fmt.Println(rowLeftToAttack+rowRightToAttack+colUpToAttack+colDownToAttack+leftUpDiagonal+leftDownDiagonal+rightUpDiagonal+rightDownDiagonal)
	fmt.Println(rowLeftToAttack, rowRightToAttack, colUpToAttack, colDownToAttack, leftUpDiagonal, leftDownDiagonal, rightUpDiagonal, rightDownDiagonal)
	// obstracles limits row if its is on the row
	// obstracles limits col if its is on col
	return 0
}