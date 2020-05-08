package main

import (
	"algs/sorts"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//a := []int64{13,-3,-25,20,-16,-23,18,-7,12,-5,-22,15,-4,7}
	r:= sorts.RSelect(readFromFile(), 5000)
	fmt.Println(r)
}


func readFromFile() []int64 {
	file, err := os.Open("integerArray.txt")
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