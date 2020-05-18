package backtracking

import (
	"fmt"
	"testing"
)

func TestNQueens(t *testing.T) {
	if re := nQueensBacktracking([]int32{}, 8); len(re) != 8 {
		fmt.Println(re)
		t.Fail()
	}
}
