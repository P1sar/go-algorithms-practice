package hackerrank

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func readFromFile(fp string) []int64 {
	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		lines = append(lines, i)
	}
	return lines
}

func TestNonDivisibleSubSet(t *testing.T) {
	res := nonDivisibleSubset(7, []int32{278,576,496,727,410,124,338,149,209,702,282,718,771,575,436})
	if res != 11 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestNonDS(t *testing.T) {
	res := nonDivisibleSubset(3, []int32{1,7,2,4})
	if res != 3 {
		fmt.Println(res)
		t.Fail()
	}
}


func TestNonDS3(t *testing.T) {
	res := nonDivisibleSubset(5, []int32{770528134,663501748,384261537,800309024,103668401,538539662,385488901,101262949,557792122,46058493})
	if res != 6 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestNonDS4(t *testing.T) {
	res := nonDivisibleSubset(4, []int32{1,2,3,4,5,6,7,8,9,10})
	if res != 5 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestRepeatedStringA(t *testing.T) {
	if res := repeatedString("aba", 10); res != 7 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestJumpingOnClouds(t *testing.T) {
	if re := jumpingOnClouds([]int32{0,0,0,1,0,0}); re != 3 {
		fmt.Println(re)
		t.Fail()
	}
}

func TestEqualizeArray(t *testing.T) {
	if re := equalizeArray([]int32{3,3,2,1,3}); re != 2 {
		fmt.Println(re)
		t.Fail()
	}
}

func TestQueensAttack(t *testing.T) {
	if re := queensAttack(5, 3, 4, 3, [][]int32{{5,5}, {4,2}, {2,3}}); re != 10 {
		fmt.Println(re)
		t.Fail()
	}
}

