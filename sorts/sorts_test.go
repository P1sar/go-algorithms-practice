package sorts

import (
	"bufio"
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

func TestRSelect(t *testing.T) {
	f := readFromFile("../test-data/integerArray.txt")
	res := RSelect(f, 5000)
	if res != 5001 {
		t.Fail()
	}
}

func BenchmarkRSelection(b *testing.B) {
	f := readFromFile("../test-data/input_dgrcode_20_1000000.txt")
	for n := 0; n < b.N; n++ {
		RSelect(f, int64(len(f)/2))
	}
	//fmt.Println(res)
}

func TestDSelect(t *testing.T) {
	f := readFromFile("../test-data/integerArray.txt")
	res := DSelect(f, 5000)
	if res != 5001 {
		t.Fail()
	}
}

func BenchmarkDSelection(b *testing.B) {
	f := readFromFile("../test-data/input_dgrcode_20_1000000.txt")
	for n := 0; n < b.N; n++ {
		DSelect(f, int64(len(f)/2))
	}
	//fmt.Println(res)
}